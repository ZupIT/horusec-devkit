package enums

import "errors"

var ErrorUnauthorized = errors.New("{HORUSEC_MIDDLEWARE} you do not have enough privileges for this action")
var ErrorFailedToVerifyRequest = errors.New("{HORUSEC_MIDDLEWARE} something went wrong while verifying " +
	"if request is authorized")
var ErrorWhenGettingAuthConfig = errors.New("{HORUSEC_MIDDLEWARE} failed to get auth config")
