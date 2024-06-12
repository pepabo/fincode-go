// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CustomersCustomerIDCardsIDDelete implements DELETE /customers/{customer_id}/cards/{id} operation.
	//
	// DELETE /customers/{customer_id}/cards/{id}
	CustomersCustomerIDCardsIDDelete(ctx context.Context, params CustomersCustomerIDCardsIDDeleteParams) (CustomersCustomerIDCardsIDDeleteRes, error)
	// CustomersCustomerIDCardsIDGet implements GET /customers/{customer_id}/cards/{id} operation.
	//
	// GET /customers/{customer_id}/cards/{id}
	CustomersCustomerIDCardsIDGet(ctx context.Context, params CustomersCustomerIDCardsIDGetParams) (CustomersCustomerIDCardsIDGetRes, error)
	// CustomersCustomerIDPaymentMethodsPaymentMethodIDDelete implements DELETE /customers/{customer_id}/payment_methods/{payment_method_id} operation.
	//
	// DELETE /customers/{customer_id}/payment_methods/{payment_method_id}
	CustomersCustomerIDPaymentMethodsPaymentMethodIDDelete(ctx context.Context, params CustomersCustomerIDPaymentMethodsPaymentMethodIDDeleteParams) (CustomersCustomerIDPaymentMethodsPaymentMethodIDDeleteRes, error)
	// CustomersCustomerIDPaymentMethodsPaymentMethodIDGet implements GET /customers/{customer_id}/payment_methods/{payment_method_id} operation.
	//
	// GET /customers/{customer_id}/payment_methods/{payment_method_id}
	CustomersCustomerIDPaymentMethodsPaymentMethodIDGet(ctx context.Context, params CustomersCustomerIDPaymentMethodsPaymentMethodIDGetParams) (CustomersCustomerIDPaymentMethodsPaymentMethodIDGetRes, error)
	// CustomersCustomerIDPaymentMethodsPost implements POST /customers/{customer_id}/payment_methods operation.
	//
	// POST /customers/{customer_id}/payment_methods
	CustomersCustomerIDPaymentMethodsPost(ctx context.Context, req CustomersCustomerIDPaymentMethodsPostReq, params CustomersCustomerIDPaymentMethodsPostParams) (CustomersCustomerIDPaymentMethodsPostRes, error)
	// CustomersIDDelete implements DELETE /customers/{id} operation.
	//
	// DELETE /customers/{id}
	CustomersIDDelete(ctx context.Context, params CustomersIDDeleteParams) (CustomersIDDeleteRes, error)
	// CustomersIDGet implements GET /customers/{id} operation.
	//
	// GET /customers/{id}
	CustomersIDGet(ctx context.Context, params CustomersIDGetParams) (CustomersIDGetRes, error)
	// CustomersPost implements POST /customers operation.
	//
	// POST /customers
	CustomersPost(ctx context.Context, req *CustomersPostReq) (CustomersPostRes, error)
	// PaymentsGet implements GET /payments operation.
	//
	// GET /payments
	PaymentsGet(ctx context.Context, params PaymentsGetParams) (PaymentsGetRes, error)
	// PaymentsIDAuthPut implements PUT /payments/{id}/auth operation.
	//
	// PUT /payments/{id}/auth
	PaymentsIDAuthPut(ctx context.Context, req PaymentsIDAuthPutReq, params PaymentsIDAuthPutParams) (PaymentsIDAuthPutRes, error)
	// PaymentsIDCancelPut implements PUT /payments/{id}/cancel operation.
	//
	// PUT /payments/{id}/cancel
	PaymentsIDCancelPut(ctx context.Context, req PaymentsIDCancelPutReq, params PaymentsIDCancelPutParams) (PaymentsIDCancelPutRes, error)
	// PaymentsIDCapturePut implements PUT /payments/{id}/capture operation.
	//
	// PUT /payments/{id}/capture
	PaymentsIDCapturePut(ctx context.Context, req PaymentsIDCapturePutReq, params PaymentsIDCapturePutParams) (PaymentsIDCapturePutRes, error)
	// PaymentsIDGet implements GET /payments/{id} operation.
	//
	// GET /payments/{id}
	PaymentsIDGet(ctx context.Context, params PaymentsIDGetParams) (PaymentsIDGetRes, error)
	// PaymentsIDPut implements PUT /payments/{id} operation.
	//
	// PUT /payments/{id}
	PaymentsIDPut(ctx context.Context, req PaymentsIDPutReq, params PaymentsIDPutParams) (PaymentsIDPutRes, error)
	// PaymentsPost implements POST /payments operation.
	//
	// POST /payments
	PaymentsPost(ctx context.Context, req PaymentsPostReq) (PaymentsPostRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
