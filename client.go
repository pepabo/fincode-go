package fincode

import (
	"context"
	"os"

	"github.com/pepabo/fincode-go/api"
)

var _ api.SecuritySource = (*clientConfig)(nil)

type clientConfig struct {
	endpoint     string
	apiSecretKey string
}

// Option is a functional option for the fincode client.
type Option func(*clientConfig) error

// Endpoint sets the endpoint of the fincode API.
func Endpoint(e string) Option {
	return func(c *clientConfig) error {
		c.endpoint = e
		return nil
	}
}

// APISecretKey sets the API secret key for the fincode API.
func APISecretKey(k string) Option {
	return func(c *clientConfig) error {
		c.apiSecretKey = k
		return nil
	}
}

func (c *clientConfig) BearerAuth(ctx context.Context, operationName string) (api.BearerAuth, error) {
	return api.BearerAuth{Token: c.apiSecretKey}, nil
}

// New creates a new fincode client.
func New(opts ...Option) (*api.Client, error) {
	c := &clientConfig{
		endpoint:     "https://api.fincode.jp/v1",
		apiSecretKey: os.Getenv("FINCODE_API_SECRET_KEY"),
	}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return api.NewClient(c.endpoint, c)
}
