package fincode

import (
	"context"
	"testing"
	"time"

	"github.com/k1LoW/runn"
	"github.com/pepabo/fincode-go/api"
)

func TestPayments(t *testing.T) {
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

	c, err := New(Endpoint(testEndpoint))
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
			ID: orderID,
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
			PayType: api.PaymentsGetPayTypeCard,
			ProcessDataFrom: api.NewOptString(today),
			Limit: api.NewOptInt(100),
			CustomerID: api.NewOptString(customerID),
		})
		if err != nil {
			t.Fatal(err)
		}
		v, ok := res.(*api.PaymentsGetOK)
		if !ok {
			t.Fatalf("unexpected response: %T, %#v", res, res)
		}
		for _, p := range v.List {
			if p.AccessID.Value ==  accessID {
				return
			}
		}
		t.Errorf("payment not found: %s", accessID)
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
