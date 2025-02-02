package fincode

import (
	"context"
	"testing"
	"time"

	"github.com/k1LoW/runn"
	"github.com/pepabo/fincode-go/api"
)

func TestPayments(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	ts := NewPaymentServer(t)
	opts := []runn.Option{
		runn.T(t),
		runn.Book("testdata/scenarios/register_card_tds_with_challenge.yml"),
		runn.SkipIncluded(true),
		runn.Var("endpoint", ts.URL),
	}
	o, err := runn.New(opts...)
	if err != nil {
		t.Fatal(err)
	}
	if err := o.Run(ctx); err != nil {
		t.Fatal(err)
	}

	store := o.Store()
	customerID, ok := store["customer_id"].(string)
	if !ok {
		t.Fatal("customer_id not found")
	}
	cardID, ok := store["card_id"].(string)
	if !ok {
		t.Fatal("card_id not found")
	}

	holderName := "FINCODE MEMBER" // same as testdata/scenarios/register_card_tds_with_challenge.yml

	t.Logf("customer_id: %s", customerID)
	t.Logf("card_id: %s", cardID)

	c, err := New(Endpoint(testEndpoint), WithHTTPClient(retryableHTTPClient(t)))
	if err != nil {
		t.Fatal(err)
	}

	orderID := newOrderID(t)

	var accessID string

	t.Run("Capture Card Payment (Non 3D Secure)", func(t *testing.T) {
		{
			res, err := c.PaymentsPost(ctx, api.PaymentsPostReq{
				Type: api.PaymentCardPaymentsPostReq,
				PaymentCard: api.PaymentCard{
					PayType: api.PaymentCardPayTypeCard,
					ID:      api.NewOptString(orderID),
					JobCode: api.PaymentCardJobCodeCAPTURE,
					Amount:  api.NewOptString("1000"),
					Tax:     api.NewOptString("80"),
					TdsType: api.NewOptPaymentCardTdsType(api.PaymentCardTdsType0),
				},
			})
			if err != nil {
				t.Fatal(err)
			}
			v, ok := res.(*api.PaymentsPostOK)
			if !ok {
				t.Fatalf("unexpected response: %T, %#v", res, res)
			}
			if want := orderID; v.PaymentCardResponse.ID.Value != want {
				t.Errorf("want %s, got %s", want, v.PaymentCardResponse.ID.Value)
			}
			accessID = v.PaymentCardResponse.AccessID.Value
		}
		{
			res, err := c.PaymentsIDPut(ctx, api.PaymentsIDPutReq{
				Type: api.PaymentDoCardWithout3DSecurePaymentsIDPutReq,
				PaymentDoCardWithout3DSecure: api.PaymentDoCardWithout3DSecure{
					PayType:    api.PaymentDoCardWithout3DSecurePayTypeCard,
					AccessID:   accessID,
					CustomerID: api.NewOptString(customerID),
					CardID:     api.NewOptString(cardID),
					Method:     api.NewOptPaymentDoCardWithout3DSecureMethod(api.PaymentDoCardWithout3DSecureMethod1),
				},
			}, api.PaymentsIDPutParams{
				ID: orderID,
			})
			if err != nil {
				t.Fatal(err)
			}
			v, ok := res.(*api.PaymentsIDPutOK)
			if !ok {
				t.Fatalf("unexpected response: %T %#v", res, res)
			}
			if want := orderID; v.PaymentDoCardResponse.ID.Value != want {
				t.Errorf("want %s, got %s", want, v.PaymentDoCardResponse.ID.Value)
			}
		}
	})

	t.Run("Get a Payment", func(t *testing.T) {
		res, err := c.PaymentsIDGet(ctx, api.PaymentsIDGetParams{
			ID:      orderID,
			PayType: "Card",
		})
		if err != nil {
			t.Fatal(err)
		}
		v, ok := res.(*api.PaymentsIDGetOK)
		if !ok {
			t.Fatalf("unexpected response: %T, %#v", res, res)
		}
		if want := orderID; v.PaymentCardResponse.ID.Value != want {
			t.Errorf("want %s, got %s", want, v.PaymentCardResponse.ID.Value)
		}

		if want := accessID; v.PaymentCardResponse.AccessID.Value != want {
			t.Errorf("want %s, got %s", want, v.PaymentCardResponse.AccessID.Value)
		}
	})

	t.Run("List Payments", func(t *testing.T) {
		today := time.Now().Format("2006/01/02")
		res, err := c.PaymentsGet(ctx, api.PaymentsGetParams{
			PayType:         "Card",
			ProcessDataFrom: api.NewOptString(today),
			Limit:           api.NewOptInt(100),
			CustomerID:      api.NewOptString(customerID),
		})
		if err != nil {
			t.Fatal(err)
		}
		v, ok := res.(*api.PaymentsGetOK)
		if !ok {
			t.Fatalf("unexpected response: %T, %#v", res, res)
		}
		for _, p := range v.List {
			if p.AccessID.Value == accessID {
				return
			}
		}
		t.Errorf("payment not found: %s", accessID)
	})

	t.Run("Cancel Payment", func(t *testing.T) {
		res, err := c.PaymentsIDCancelPut(ctx, api.PaymentsIDCancelPutReq{
			Type: api.PaymentCancelCardPaymentsIDCancelPutReq,
			PaymentCancelCard: api.PaymentCancelCard{
				PayType:  "Card",
				AccessID: accessID,
			},
		},
			api.PaymentsIDCancelPutParams{
				ID: orderID,
			})
		if err != nil {
			t.Fatal(err)
		}
		v, ok := res.(*api.PaymentsIDCancelPutOK)
		if !ok {
			t.Fatalf("unexpected response: %T, %#v", res, res)
		}
		if want := orderID; v.PaymentCancelCardResponse.ID.Value != want {
			t.Errorf("want %s, got %s", want, v.PaymentCancelCardResponse.ID.Value)
		}
		if want := api.PaymentCancelCardResponseJobCodeCANCEL; v.PaymentCancelCardResponse.JobCode.Value != want {
			t.Errorf("want %s, got %s", want, v.PaymentCancelCardResponse.JobCode.Value)
		}
		if want := api.PaymentCancelCardResponseStatusCANCELED; v.PaymentCancelCardResponse.Status.Value != want {
			t.Errorf("want %s, got %s", want, v.PaymentCancelCardResponse.Status.Value)
		}
		paymentsIDGetRes, err := c.PaymentsIDGet(ctx, api.PaymentsIDGetParams{
			ID:      orderID,
			PayType: "Card",
		})
		if err != nil {
			t.Fatal(err)
		}
		val, ok := paymentsIDGetRes.(*api.PaymentsIDGetOK)
		if !ok {
			t.Fatalf("unexpected response: %T, %#v", res, res)
		}
		if want := orderID; val.PaymentCardResponse.ID.Value != want {
			t.Errorf("want %s, got %s", want, val.PaymentCardResponse.ID.Value)
		}
		if want := api.PaymentCardResponseStatusCANCELED; val.PaymentCardResponse.Status.Value != want {
			t.Errorf("want %s, got %s", want, val.PaymentCardResponse.Status.Value)
		}
	})

	t.Run("Auth Payment", func(t *testing.T) {
		res, err := c.PaymentsIDAuthPut(ctx, api.PaymentsIDAuthPutReq{
			Type: api.PaymentAuthCardPaymentsIDAuthPutReq,
			PaymentAuthCard: api.PaymentAuthCard{
				PayType:  "Card",
				AccessID: accessID,
				Method:   api.PaymentAuthCardMethod1,
			},
		},
			api.PaymentsIDAuthPutParams{
				ID: orderID,
			})
		if err != nil {
			t.Fatal(err)
		}
		v, ok := res.(*api.PaymentsIDAuthPutOK)
		if !ok {
			t.Fatalf("unexpected response: %T, %#v", res, res)
		}
		if want := orderID; v.PaymentAuthCardResponse.ID.Value != want {
			t.Errorf("want %s, got %s", want, v.PaymentAuthCardResponse.ID.Value)
		}
		if want := api.PaymentAuthCardResponseJobCodeAUTH; v.PaymentAuthCardResponse.JobCode.Value != want {
			t.Errorf("want %s, got %s", want, v.PaymentAuthCardResponse.JobCode.Value)
		}
		if want := api.PaymentAuthCardResponseStatusAUTHORIZED; v.PaymentAuthCardResponse.Status.Value != want {
			t.Errorf("want %s, got %s", want, v.PaymentAuthCardResponse.Status.Value)
		}
		paymentsIDGetRes, err := c.PaymentsIDGet(ctx, api.PaymentsIDGetParams{
			ID:      orderID,
			PayType: "Card",
		})
		if err != nil {
			t.Fatal(err)
		}
		val, ok := paymentsIDGetRes.(*api.PaymentsIDGetOK)
		if !ok {
			t.Fatalf("unexpected response: %T, %#v", res, res)
		}
		if want := orderID; val.PaymentCardResponse.ID.Value != want {
			t.Errorf("want %s, got %s", want, val.PaymentCardResponse.ID.Value)
		}
		if want := api.PaymentCardResponseStatusAUTHORIZED; val.PaymentCardResponse.Status.Value != want {
			t.Errorf("want %s, got %s", want, val.PaymentCardResponse.Status.Value)
		}
	})

	t.Run("Capture Payment", func(t *testing.T) {
		res, err := c.PaymentsIDCapturePut(ctx, api.PaymentsIDCapturePutReq{
			Type: api.PaymentCaptureCardPaymentsIDCapturePutReq,
			PaymentCaptureCard: api.PaymentCaptureCard{
				PayType:  "Card",
				AccessID: accessID,
			},
		},
			api.PaymentsIDCapturePutParams{
				ID: orderID,
			})
		if err != nil {
			t.Fatal(err)
		}
		v, ok := res.(*api.PaymentsIDCapturePutOK)
		if !ok {
			t.Fatalf("unexpected response: %T, %#v", res, res)
		}
		if want := orderID; v.PaymentCaptureCardResponse.ID.Value != want {
			t.Errorf("want %s, got %s", want, v.PaymentCaptureCardResponse.ID.Value)
		}
		if want := api.PaymentCaptureCardResponseJobCodeSALES; v.PaymentCaptureCardResponse.JobCode.Value != want {
			t.Errorf("want %s, got %s", want, v.PaymentCaptureCardResponse.JobCode.Value)
		}
		if want := api.PaymentCaptureCardResponseStatusCAPTURED; v.PaymentCaptureCardResponse.Status.Value != want {
			t.Errorf("want %s, got %s", want, v.PaymentCaptureCardResponse.Status.Value)
		}
		paymentsIDGetRes, err := c.PaymentsIDGet(ctx, api.PaymentsIDGetParams{
			ID:      orderID,
			PayType: "Card",
		})
		if err != nil {
			t.Fatal(err)
		}
		val, ok := paymentsIDGetRes.(*api.PaymentsIDGetOK)
		if !ok {
			t.Fatalf("unexpected response: %T, %#v", res, res)
		}
		if want := orderID; val.PaymentCardResponse.ID.Value != want {
			t.Errorf("want %s, got %s", want, val.PaymentCardResponse.ID.Value)
		}
		if want := api.PaymentCardResponseStatusCAPTURED; val.PaymentCardResponse.Status.Value != want {
			t.Errorf("want %s, got %s", want, val.PaymentCardResponse.Status.Value)
		}
	})

	t.Run("Get Card", func(t *testing.T) {
		res, err := c.CustomersCustomerIDCardsIDGet(ctx, api.CustomersCustomerIDCardsIDGetParams{
			CustomerID: customerID,
			ID:         cardID,
		})
		if err != nil {
			t.Fatal(err)
		}
		v, ok := res.(*api.CustomersCustomerIDCardsIDGetOK)
		if !ok {
			t.Fatalf("unexpected response: %T, %#v", res, res)
		}
		if want := holderName; v.HolderName.Value != want {
			t.Errorf("want %s, got %s", want, v.HolderName.Value)
		}
	})

	t.Run("Delete Card", func(t *testing.T) {
		res, err := c.CustomersCustomerIDCardsIDDelete(ctx, api.CustomersCustomerIDCardsIDDeleteParams{
			CustomerID: customerID,
			ID:         cardID,
		})
		if err != nil {
			t.Fatal(err)
		}
		v, ok := res.(*api.CustomersCustomerIDCardsIDDeleteOK)
		if !ok {
			t.Fatalf("unexpected response: %T, %#v", res, res)
		}
		if want := cardID; v.ID != want {
			t.Errorf("want %s, got %s", want, v.ID)
		}
		if want := "1"; v.DeleteFlag != want {
			t.Errorf("want %s, got %s", want, v.DeleteFlag)
		}
	})

	t.Cleanup(func() {
		o.Close(false)
	})
}
