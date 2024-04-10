package fincode

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pepabo/fincode-go/api"
)

const testEndpoint = "https://api.test.fincode.jp/v1"

func TestCustomers(t *testing.T) {
	ctx := context.Background()
	c, err := New(Endpoint(testEndpoint))
	if err != nil {
		t.Fatal(err)
	}
	faker := gofakeit.NewCrypto()
	id := newID(t)
	name := faker.Name()
	email := faker.Email()

	t.Run("Create customer", func(t *testing.T) {
		res, err := c.CustomersPost(ctx, &api.CustomersPostReq{
			ID:    id,
			Name:  name,
			Email: email,
		})
		if err != nil {
			t.Fatal(err)
		}
		v, ok := res.(*api.CustomersPostOK)
		if !ok {
			t.Errorf("unexpected response type: %T", res)
		}
		if want := id; v.ID != want {
			t.Errorf("unexpected ID: got %q, want %q", v.ID, want)
		}
	})

	t.Run("Get customer", func(t *testing.T) {
		res, err := c.CustomersIDGet(ctx, api.CustomersIDGetParams{
			ID: id,
		})
		if err != nil {
			t.Fatal(err)
		}
		v, ok := res.(*api.CustomersIDGetOK)
		if !ok {
			t.Errorf("unexpected response type: %T", res)
		}
		if want := id; v.ID != want {
			t.Errorf("unexpected ID: got %q, want %q", v.ID, want)
		}
	})

	t.Run("Delete customer", func(t *testing.T) {
		res, err := c.CustomersIDDelete(ctx, api.CustomersIDDeleteParams{
			ID: id,
		})
		if err != nil {
			t.Fatal(err)
		}
		v, ok := res.(*api.CustomersIDDeleteOK)
		if !ok {
			t.Errorf("unexpected response type: %T", res)
		}
		if want := id; v.ID != want {
			t.Errorf("unexpected ID: got %q, want %q", v.ID, want)
		}
	})
}
