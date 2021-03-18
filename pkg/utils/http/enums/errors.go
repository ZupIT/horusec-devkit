package enums

import "errors"

const (
	FailedToMakeHTTPRequest = "{ERROR_HTTP} failed to make request"
)

var ErrorGenericInternalError = errors.New("{ERROR_HTTP} something went wrong, sorry for the inconvenience")
