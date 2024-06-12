// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [2]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "customers"
				origElem := elem
				if l := len("customers"); len(elem) >= l && elem[0:l] == "customers" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "POST":
						s.handleCustomersPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "customer_id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleCustomersIDDeleteRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleCustomersIDGetRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						origElem := elem
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'c': // Prefix: "cards/"
							origElem := elem
							if l := len("cards/"); len(elem) >= l && elem[0:l] == "cards/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "id"
							// Leaf parameter
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "DELETE":
									s.handleCustomersCustomerIDCardsIDDeleteRequest([2]string{
										args[0],
										args[1],
									}, elemIsEscaped, w, r)
								case "GET":
									s.handleCustomersCustomerIDCardsIDGetRequest([2]string{
										args[0],
										args[1],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "DELETE,GET")
								}

								return
							}

							elem = origElem
						case 'p': // Prefix: "payment_methods"
							origElem := elem
							if l := len("payment_methods"); len(elem) >= l && elem[0:l] == "payment_methods" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch r.Method {
								case "POST":
									s.handleCustomersCustomerIDPaymentMethodsPostRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}
							switch elem[0] {
							case '/': // Prefix: "/"
								origElem := elem
								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "payment_method_id"
								// Leaf parameter
								args[1] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf node.
									switch r.Method {
									case "DELETE":
										s.handleCustomersCustomerIDPaymentMethodsPaymentMethodIDDeleteRequest([2]string{
											args[0],
											args[1],
										}, elemIsEscaped, w, r)
									case "GET":
										s.handleCustomersCustomerIDPaymentMethodsPaymentMethodIDGetRequest([2]string{
											args[0],
											args[1],
										}, elemIsEscaped, w, r)
									default:
										s.notAllowed(w, r, "DELETE,GET")
									}

									return
								}

								elem = origElem
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'p': // Prefix: "payments"
				origElem := elem
				if l := len("payments"); len(elem) >= l && elem[0:l] == "payments" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handlePaymentsGetRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handlePaymentsPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handlePaymentsIDGetRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "PUT":
							s.handlePaymentsIDPutRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET,PUT")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/cancel"
						origElem := elem
						if l := len("/cancel"); len(elem) >= l && elem[0:l] == "/cancel" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "PUT":
								s.handlePaymentsIDCancelPutRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "PUT")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [2]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "customers"
				origElem := elem
				if l := len("customers"); len(elem) >= l && elem[0:l] == "customers" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "POST":
						r.name = "CustomersPost"
						r.summary = ""
						r.operationID = ""
						r.pathPattern = "/customers"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "customer_id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "CustomersIDDelete"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/customers/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "CustomersIDGet"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/customers/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						origElem := elem
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'c': // Prefix: "cards/"
							origElem := elem
							if l := len("cards/"); len(elem) >= l && elem[0:l] == "cards/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "id"
							// Leaf parameter
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "DELETE":
									r.name = "CustomersCustomerIDCardsIDDelete"
									r.summary = ""
									r.operationID = ""
									r.pathPattern = "/customers/{customer_id}/cards/{id}"
									r.args = args
									r.count = 2
									return r, true
								case "GET":
									r.name = "CustomersCustomerIDCardsIDGet"
									r.summary = ""
									r.operationID = ""
									r.pathPattern = "/customers/{customer_id}/cards/{id}"
									r.args = args
									r.count = 2
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'p': // Prefix: "payment_methods"
							origElem := elem
							if l := len("payment_methods"); len(elem) >= l && elem[0:l] == "payment_methods" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "POST":
									r.name = "CustomersCustomerIDPaymentMethodsPost"
									r.summary = ""
									r.operationID = ""
									r.pathPattern = "/customers/{customer_id}/payment_methods"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
							switch elem[0] {
							case '/': // Prefix: "/"
								origElem := elem
								if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
									elem = elem[l:]
								} else {
									break
								}

								// Param: "payment_method_id"
								// Leaf parameter
								args[1] = elem
								elem = ""

								if len(elem) == 0 {
									// Leaf node.
									switch method {
									case "DELETE":
										r.name = "CustomersCustomerIDPaymentMethodsPaymentMethodIDDelete"
										r.summary = ""
										r.operationID = ""
										r.pathPattern = "/customers/{customer_id}/payment_methods/{payment_method_id}"
										r.args = args
										r.count = 2
										return r, true
									case "GET":
										r.name = "CustomersCustomerIDPaymentMethodsPaymentMethodIDGet"
										r.summary = ""
										r.operationID = ""
										r.pathPattern = "/customers/{customer_id}/payment_methods/{payment_method_id}"
										r.args = args
										r.count = 2
										return r, true
									default:
										return
									}
								}

								elem = origElem
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'p': // Prefix: "payments"
				origElem := elem
				if l := len("payments"); len(elem) >= l && elem[0:l] == "payments" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "PaymentsGet"
						r.summary = ""
						r.operationID = ""
						r.pathPattern = "/payments"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "PaymentsPost"
						r.summary = ""
						r.operationID = ""
						r.pathPattern = "/payments"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "PaymentsIDGet"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/payments/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PUT":
							r.name = "PaymentsIDPut"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/payments/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/cancel"
						origElem := elem
						if l := len("/cancel"); len(elem) >= l && elem[0:l] == "/cancel" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "PUT":
								r.name = "PaymentsIDCancelPut"
								r.summary = ""
								r.operationID = ""
								r.pathPattern = "/payments/{id}/cancel"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	return r, false
}
