// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CustomersIDDelete implements DELETE /customers/{id} operation.
//
// DELETE /customers/{id}
func (UnimplementedHandler) CustomersIDDelete(ctx context.Context, params CustomersIDDeleteParams) (r CustomersIDDeleteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CustomersIDGet implements GET /customers/{id} operation.
//
// GET /customers/{id}
func (UnimplementedHandler) CustomersIDGet(ctx context.Context, params CustomersIDGetParams) (r CustomersIDGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CustomersPost implements POST /customers operation.
//
// POST /customers
func (UnimplementedHandler) CustomersPost(ctx context.Context, req *CustomersPostReq) (r CustomersPostRes, _ error) {
	return r, ht.ErrNotImplemented
}
