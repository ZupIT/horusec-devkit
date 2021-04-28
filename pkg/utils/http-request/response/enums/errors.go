package enums

import "errors"

var ErrDoHTTPServiceSide = errors.New("{ERROR_HTTP_CLIENT} Response HTTP return server side error")
var ErrDoHTTPClientSide = errors.New("{ERROR_HTTP_CLIENT} Response HTTP return client side error")
var ErrHTTPResponse = errors.New("{ERROR_HTTP_REQUEST}")
