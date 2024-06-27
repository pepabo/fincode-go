package fincode

import (
	"context"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/pepabo/fincode-go/api"
	"golang.org/x/sync/errgroup"
)

func TestWithHTTPClient(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	c, err := New(Endpoint(testEndpoint), WithHTTPClient(retryableHTTPClient(t)))
	if err != nil {
		t.Fatal(err)
	}
	faker := gofakeit.NewCrypto()
	eg := new(errgroup.Group)
	for range 10 {
		eg.Go(func() error {
			id := newID(t)
			name := faker.Name()
			email := faker.Email()
			res, err := c.CustomersPost(ctx, &api.CustomersPostReq{
				ID:    id,
				Name:  name,
				Email: email,
			})
			if err != nil {
				return err
			}
			_, ok := res.(*api.CustomersPostOK)
			if !ok {
				t.Errorf("unexpected response type: %T", res)
			}
			{
				res, err := c.CustomersIDDelete(ctx, api.CustomersIDDeleteParams{
					ID: id,
				})
				if err != nil {
					t.Fatal(err)
				}
				_, ok := res.(*api.CustomersIDDeleteOK)
				if !ok {
					t.Errorf("unexpected response type: %T", res)
				}
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		t.Error(err)
	}
}

func retryableHTTPClient(t *testing.T) *http.Client {
	t.Helper()
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	retryClient.Logger = nil
	retryClient.CheckRetry = func(ctx context.Context, res *http.Response, err error) (bool, error) {
		if res.StatusCode == 403 {
			return true, nil
		}
		return false, nil
	}
	return retryClient.StandardClient()
}
