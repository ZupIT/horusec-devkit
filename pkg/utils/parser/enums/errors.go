package enums

import "errors"

var ErrorBodyEmpty = errors.New("{ERROR_PARSER} body can not be empty")
var ErrorBodyInvalid = errors.New("{ERROR_PARSER} body is invalid format type")
