package fincode

import (
	"context"
	"os"

	ht "github.com/ogen-go/ogen/http"
	"github.com/pepabo/fincode-go/api"
)

var _ api.SecuritySource = (*clientConfig)(nil)

type clientConfig struct {
	endpoint     string
	apiSecretKey string
	httpClient   ht.Client
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

// WithHTTPClient specifies http client to use.
func WithHTTPClient(client ht.Client) Option {
	return func(c *clientConfig) error {
		c.httpClient = client
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
	var copts []api.ClientOption
	if c.httpClient != nil {
		copts = append(copts, api.WithClient(c.httpClient))
	}

	return api.NewClient(c.endpoint, c, copts...)
}
