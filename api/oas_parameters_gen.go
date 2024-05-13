// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// CustomersCustomerIDPaymentMethodsPaymentMethodIDDeleteParams is parameters of DELETE /customers/{customer_id}/payment_methods/{payment_method_id} operation.
type CustomersCustomerIDPaymentMethodsPaymentMethodIDDeleteParams struct {
	CustomerID      string
	PaymentMethodID string
}

func unpackCustomersCustomerIDPaymentMethodsPaymentMethodIDDeleteParams(packed middleware.Parameters) (params CustomersCustomerIDPaymentMethodsPaymentMethodIDDeleteParams) {
	{
		key := middleware.ParameterKey{
			Name: "customer_id",
			In:   "path",
		}
		params.CustomerID = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "payment_method_id",
			In:   "path",
		}
		params.PaymentMethodID = packed[key].(string)
	}
	return params
}

func decodeCustomersCustomerIDPaymentMethodsPaymentMethodIDDeleteParams(args [2]string, argsEscaped bool, r *http.Request) (params CustomersCustomerIDPaymentMethodsPaymentMethodIDDeleteParams, _ error) {
	// Decode path: customer_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "customer_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.CustomerID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "customer_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: payment_method_id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "payment_method_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.PaymentMethodID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "payment_method_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// CustomersCustomerIDPaymentMethodsPaymentMethodIDGetParams is parameters of GET /customers/{customer_id}/payment_methods/{payment_method_id} operation.
type CustomersCustomerIDPaymentMethodsPaymentMethodIDGetParams struct {
	CustomerID      string
	PaymentMethodID string
}

func unpackCustomersCustomerIDPaymentMethodsPaymentMethodIDGetParams(packed middleware.Parameters) (params CustomersCustomerIDPaymentMethodsPaymentMethodIDGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "customer_id",
			In:   "path",
		}
		params.CustomerID = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "payment_method_id",
			In:   "path",
		}
		params.PaymentMethodID = packed[key].(string)
	}
	return params
}

func decodeCustomersCustomerIDPaymentMethodsPaymentMethodIDGetParams(args [2]string, argsEscaped bool, r *http.Request) (params CustomersCustomerIDPaymentMethodsPaymentMethodIDGetParams, _ error) {
	// Decode path: customer_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "customer_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.CustomerID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "customer_id",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: payment_method_id.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "payment_method_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.PaymentMethodID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "payment_method_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// CustomersCustomerIDPaymentMethodsPostParams is parameters of POST /customers/{customer_id}/payment_methods operation.
type CustomersCustomerIDPaymentMethodsPostParams struct {
	CustomerID string
}

func unpackCustomersCustomerIDPaymentMethodsPostParams(packed middleware.Parameters) (params CustomersCustomerIDPaymentMethodsPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "customer_id",
			In:   "path",
		}
		params.CustomerID = packed[key].(string)
	}
	return params
}

func decodeCustomersCustomerIDPaymentMethodsPostParams(args [1]string, argsEscaped bool, r *http.Request) (params CustomersCustomerIDPaymentMethodsPostParams, _ error) {
	// Decode path: customer_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "customer_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.CustomerID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "customer_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// CustomersIDDeleteParams is parameters of DELETE /customers/{id} operation.
type CustomersIDDeleteParams struct {
	ID string
}

func unpackCustomersIDDeleteParams(packed middleware.Parameters) (params CustomersIDDeleteParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(string)
	}
	return params
}

func decodeCustomersIDDeleteParams(args [1]string, argsEscaped bool, r *http.Request) (params CustomersIDDeleteParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// CustomersIDGetParams is parameters of GET /customers/{id} operation.
type CustomersIDGetParams struct {
	ID string
}

func unpackCustomersIDGetParams(packed middleware.Parameters) (params CustomersIDGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(string)
	}
	return params
}

func decodeCustomersIDGetParams(args [1]string, argsEscaped bool, r *http.Request) (params CustomersIDGetParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PaymentsGetParams is parameters of GET /payments operation.
type PaymentsGetParams struct {
	// 1回で取得するデータ件数.
	Limit OptInt
	// ページ数.
	Page OptInt
	// 総件数のみを取得するフラグ.
	CountOnly OptBool
	// ソート順の定義
	// ※ソート可能な項目
	// status - 決済ステータス
	// process_date - 処理日時
	// total_amount - 利用金額と税送料の合計金額
	// auth_max_date - 仮売上有効期限
	// created - 作成日時
	// updated - 更新日時.
	Sort OptNilString
	// 決済種別
	// Card- クレジットカード決済
	// Applepay - Apple Pay
	// Konbini- コンビニ決済
	// Paypay- PayPay.
	PayType string
	// 加盟店自由項目1-3に対する部分一致.
	Keyword OptString
	// 利用金額+税送料の合計(min).
	TotalAmountMin OptInt
	// 利用金額+税送料の合計(max).
	TotalAmountMax OptString
	// 顧客ID.
	CustomerID OptString
	// 作成日の範囲指定の始点。YYYY/MM/dd 形式.
	ProcessDataFrom OptString
	// 作成日の範囲指定の終点。YYYY/MM/dd 形式.
	ProcessDataTo OptString
	// 仮売上有効期限の始点。YYYY/MM/dd 形式.
	AuthMaxDateFrom OptString
	// 仮売上有効期限の終点。YYYY/MM/dd 形式.
	AuthMaxDateTo OptString
	// 更新日の範囲指定の始値。YYYY/MM/dd 形式.
	UpdateDateFrom OptString
	// 更新日の範囲指定の終点。YYYY/MM/dd 形式.
	UpdateDateTo OptString
	// 決済ステータス
	// (カンマ区切りで複数指定可能)
	// UNPROCESSED - 未決済
	// CHECKED - 有効性チェック
	// AUTHORIZED - 仮売上
	// CAPTURED - 売上確定
	// CANCELED - キャンセル
	// AUTHENTICATED - 未決済（3Dセキュア）.
	Status OptString
	// 課金種別
	// (カンマ区切りで複数指定可能)
	// onetime - サブスクリプション以外
	// subscription - サブスクリプション.
	PayPattern OptString
	// サブスクリプションID.
	SubscriptionID OptString
}

func unpackPaymentsGetParams(packed middleware.Parameters) (params PaymentsGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "limit",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Limit = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "count_only",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.CountOnly = v.(OptBool)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "sort",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Sort = v.(OptNilString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "pay_type",
			In:   "query",
		}
		params.PayType = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "keyword",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Keyword = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "total_amount_min",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.TotalAmountMin = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "total_amount_max",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.TotalAmountMax = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "customer_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.CustomerID = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "process_data_from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ProcessDataFrom = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "process_data_to",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.ProcessDataTo = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "auth_max_date_from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.AuthMaxDateFrom = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "auth_max_date_to",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.AuthMaxDateTo = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "update_date_from",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UpdateDateFrom = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "update_date_to",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.UpdateDateTo = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "status",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Status = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "pay_pattern",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.PayPattern = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "subscription_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.SubscriptionID = v.(OptString)
		}
	}
	return params
}

func decodePaymentsGetParams(args [0]string, argsEscaped bool, r *http.Request) (params PaymentsGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: limit.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLimitVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotLimitVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Limit.SetTo(paramsDotLimitVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "limit",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: count_only.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "count_only",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCountOnlyVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotCountOnlyVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CountOnly.SetTo(paramsDotCountOnlyVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "count_only",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: sort.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "sort",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSortVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotSortVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Sort.SetTo(paramsDotSortVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "sort",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: pay_type.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "pay_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.PayType = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "pay_type",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: keyword.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "keyword",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotKeywordVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotKeywordVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Keyword.SetTo(paramsDotKeywordVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "keyword",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: total_amount_min.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "total_amount_min",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotTotalAmountMinVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotTotalAmountMinVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.TotalAmountMin.SetTo(paramsDotTotalAmountMinVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "total_amount_min",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: total_amount_max.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "total_amount_max",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotTotalAmountMaxVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotTotalAmountMaxVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.TotalAmountMax.SetTo(paramsDotTotalAmountMaxVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "total_amount_max",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: customer_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "customer_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCustomerIDVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotCustomerIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CustomerID.SetTo(paramsDotCustomerIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "customer_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: process_data_from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "process_data_from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotProcessDataFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotProcessDataFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ProcessDataFrom.SetTo(paramsDotProcessDataFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "process_data_from",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: process_data_to.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "process_data_to",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotProcessDataToVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotProcessDataToVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.ProcessDataTo.SetTo(paramsDotProcessDataToVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "process_data_to",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: auth_max_date_from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "auth_max_date_from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotAuthMaxDateFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotAuthMaxDateFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.AuthMaxDateFrom.SetTo(paramsDotAuthMaxDateFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "auth_max_date_from",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: auth_max_date_to.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "auth_max_date_to",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotAuthMaxDateToVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotAuthMaxDateToVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.AuthMaxDateTo.SetTo(paramsDotAuthMaxDateToVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "auth_max_date_to",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: update_date_from.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "update_date_from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotUpdateDateFromVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotUpdateDateFromVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UpdateDateFrom.SetTo(paramsDotUpdateDateFromVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "update_date_from",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: update_date_to.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "update_date_to",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotUpdateDateToVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotUpdateDateToVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.UpdateDateTo.SetTo(paramsDotUpdateDateToVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "update_date_to",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: status.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "status",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStatusVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotStatusVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Status.SetTo(paramsDotStatusVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "status",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: pay_pattern.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "pay_pattern",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPayPatternVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotPayPatternVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.PayPattern.SetTo(paramsDotPayPatternVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "pay_pattern",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: subscription_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "subscription_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSubscriptionIDVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotSubscriptionIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.SubscriptionID.SetTo(paramsDotSubscriptionIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "subscription_id",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// PaymentsIDPutParams is parameters of PUT /payments/{id} operation.
type PaymentsIDPutParams struct {
	ID string
}

func unpackPaymentsIDPutParams(packed middleware.Parameters) (params PaymentsIDPutParams) {
	{
		key := middleware.ParameterKey{
			Name: "id",
			In:   "path",
		}
		params.ID = packed[key].(string)
	}
	return params
}

func decodePaymentsIDPutParams(args [1]string, argsEscaped bool, r *http.Request) (params PaymentsIDPutParams, _ error) {
	// Decode path: id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.ID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
