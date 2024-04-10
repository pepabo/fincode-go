package fincode

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/k1LoW/httpstub"
	"github.com/pepabo/fincode-go/api"
)

//go:embed testdata/templates/*
var tmplFS embed.FS

/*
		{
		  "card_no": "************1111",
		  "expire": "2021/12/19 21:34:58.464",
		  "list": [
		    { "token": '3561656330616461626163303663363330306' }
		  ],
		  "security_code_set": "1"
		}
	  with customer_id
*/
type tokenResponse struct {
	CardNo string `json:"card_no"`
	Expire string `json:"expire"`
	List   []struct {
		Token string `json:"token"`
	} `json:"list"`
	SecurityCodeSet string `json:"security_code_set"`
	// with customer_id
	CustomerID string `json:"customer_id"`
}

func NewPaymentServer(t *testing.T) *httptest.Server {
	t.Helper()
	ctx := context.Background()
	faker := gofakeit.NewCrypto()
	c, err := New(Endpoint(testEndpoint))
	if err != nil {
		t.Fatal(err)
	}
	ts := httpstub.NewServer(t)
	t.Cleanup(func() {
		ts.Close()
	})

	data := map[string]any{
		"Endpoint":     ts.URL,
		"APIPublicKey": os.Getenv("FINCODE_API_PUBLIC_KEY"),
	}
	{
		ts.Method(http.MethodGet).Path("/card").Handler(func(w http.ResponseWriter, r *http.Request) {
			b, err := tmplFS.ReadFile("testdata/templates/card.html.tmpl")
			if err != nil {
				t.Fatal(err)
			}
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)

			// Create customer
			customerID := r.URL.Query().Get("customer_id")
			res, err := c.CustomersPost(ctx, &api.CustomersPostReq{
				ID:    customerID,
				Name:  faker.Name(),
				Email: faker.Email(),
			})
			if err != nil {
				errorWithResponse(t, err, w)
				return
			}
			_, ok := res.(*api.CustomersPostOK)
			if !ok {
				errorWithResponse(t, fmt.Errorf("unexpected response type: %T", res), w)
				return
			}
			t.Cleanup(func() {
				if _, err := c.CustomersIDDelete(ctx, api.CustomersIDDeleteParams{
					ID: customerID,
				}); err != nil {
					t.Error(err)
				}
			})

			data["CustomerID"] = customerID
			tmpl := template.Must(template.New("/card").Parse(string(b)))
			if err := tmpl.Execute(w, data); err != nil {
				t.Fatal(err)
			}
		})
	}
	{
		ts.Method(http.MethodPost).Path("/card/register").Handler(func(w http.ResponseWriter, r *http.Request) {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				errorWithResponse(t, err, w)
				return
			}
			r.Body.Close()
			var res tokenResponse
			if err := json.Unmarshal(b, &res); err != nil {
				errorWithResponse(t, err, w)
				return
			}
			customerID := res.CustomerID
			var redirectURL string
			// Create payment method
			{
				req := api.CustomersCustomerIDPaymentMethodsPostReq{
					Type: api.PaymentMethodCardWith3DSecureCustomersCustomerIDPaymentMethodsPostReq,
					PaymentMethodCardWith3DSecure: api.PaymentMethodCardWith3DSecure{
						PayType:            api.PaymentMethodCardWith3DSecurePayTypeCard,
						ReturnURL:          api.NewOptString(ts.URL + "/card/success"),
						ReturnURLOnFailure: api.NewOptString(ts.URL + "/card/failure"),
						DefaultFlag:        api.PaymentMethodCardWith3DSecureDefaultFlag1,
						Card: api.PaymentMethodCardWith3DSecureCard{
							Token:   res.List[0].Token,
							TdsType: api.PaymentMethodCardWith3DSecureCardTdsType2,
						},
					},
				}
				params := api.CustomersCustomerIDPaymentMethodsPostParams{
					CustomerID: customerID,
				}
				res, err := c.CustomersCustomerIDPaymentMethodsPost(ctx, req, params)
				if err != nil {
					errorWithResponse(t, err, w)
					return
				}
				resOK, ok := res.(*api.CustomersCustomerIDPaymentMethodsPostOK)
				if !ok {
					errorWithResponse(t, fmt.Errorf("unexpected response type: %T", res), w)
					return
				}
				redirectURL = resOK.PaymentMethodCardResponse.RedirectURL.Value

				t.Cleanup(func() {
					if _, err := c.CustomersCustomerIDPaymentMethodsPaymentMethodIDDelete(ctx, api.CustomersCustomerIDPaymentMethodsPaymentMethodIDDeleteParams{
						CustomerID:      resOK.PaymentMethodCardResponse.CustomerID.Value,
						PaymentMethodID: resOK.PaymentMethodCardResponse.ID.Value,
					}); err != nil {
						t.Error(err)
					}
				})
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(fmt.Sprintf(`{"redirect_url":"%s"}`, redirectURL)))
		})
	}
	ts.Method(http.MethodPost).Path("/card/success").ResponseString(http.StatusOK, "Success")
	ts.Method(http.MethodPost).Path("/card/failure").ResponseString(http.StatusBadRequest, "Failure")
	{
		b, err := tmplFS.ReadFile("testdata/templates/util.js")
		if err != nil {
			t.Fatal(err)
		}
		ts.Method(http.MethodGet).Path("/util.js").ResponseString(http.StatusOK, string(b))
	}

	ts.Method(http.MethodGet).ResponseString(http.StatusNotFound, "Not Found")

	return ts.Server()
}

func newID(t *testing.T) string {
	t.Helper()
	faker := gofakeit.NewCrypto()
	return fmt.Sprintf("test-fincode-go-%s", faker.UUID())
}

func errorWithResponse(t *testing.T, err error, w http.ResponseWriter) {
	t.Helper()
	t.Error(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"error":"internal server error: %v"}`, err)))
}
