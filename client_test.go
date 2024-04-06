package fincode

import (
	"context"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

const serverURL = "https://api.test.fincode.jp/v1"

var _ SecuritySource = (*testSecuritySource)(nil)

type testSecuritySource struct {
	token string
}

func (t *testSecuritySource) BearerAuth(ctx context.Context, operationName string) (BearerAuth, error) {
	return BearerAuth{Token: t.token}, nil
}

func TestCustomers(t *testing.T) {
	ctx := context.Background()
	c, err := NewClient(serverURL, &testSecuritySource{token: os.Getenv("FINCODE_API_SECRET_KEY")})
	if err != nil {
		t.Fatal(err)
	}
	faker := gofakeit.NewCrypto()
	id := faker.UUID()
	name := faker.Name()
	email := faker.Email()

	t.Run("Create customer", func(t *testing.T) {
		res, err := c.CustomersPost(ctx, &CustomersPostReq{
			ID:    id,
			Name:  name,
			Email: email,
		})
		if err != nil {
			t.Fatal(err)
		}
		v, ok := res.(*CustomersPostOK)
		if !ok {
			t.Errorf("unexpected response type: %T", res)
		}
		if want := id; v.ID != want {
			t.Errorf("unexpected ID: got %q, want %q", v.ID, want)
		}
	})

	t.Run("Delete customer", func(t *testing.T) {
		res, err := c.CustomersIDDelete(ctx, CustomersIDDeleteParams{
			ID: id,
		})
		if err != nil {
			t.Fatal(err)
		}
		v, ok := res.(*CustomersIDDeleteOK)
		if !ok {
			t.Errorf("unexpected response type: %T", res)
		}
		if want := id; v.ID != want {
			t.Errorf("unexpected ID: got %q, want %q", v.ID, want)
		}
	})
}
