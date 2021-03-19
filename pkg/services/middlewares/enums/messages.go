package enums

const (
	GRPCRequestError    = "{HORUSEC_MIDDLEWARE} is authorized grpc method returned a error"
	UnauthorizedRequest = "{HORUSEC_MIDDLEWARE} http request made by account id \"%s\" in url \"%s\" with method" +
		" \"%s\" returned unauthorized to \"%s\""
	FailedToGetAccountID  = "{HORUSEC_MIDDLEWARE} failed to get account id for unauthorized request warning"
	FailedToGetAuthConfig = "{HORUSEC_MIDDLEWARE} grpc method to get auth config failed"
)
