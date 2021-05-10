package enums

import "errors"

var ErrorRequestServerError = errors.New("{ERROR_HTTP} request returned a error status code from the server")
var ErrorRequestClientError = errors.New("{ERROR_HTTP} request returned a error status code from the client")
