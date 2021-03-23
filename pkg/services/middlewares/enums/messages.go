package enums

const (
	MessageIsAuthorizedGRPCRequestError = "{HORUSEC_MIDDLEWARE} is authorized grpc method returned a error"
	MessageUnauthorizedHTTPRequest      = "{HORUSEC_MIDDLEWARE} http request made by account id \"%s\" in url \"%s\" " +
		"with method \"%s\" returned unauthorized to \"%s\""
	MessageFailedToGetAccountID  = "{HORUSEC_MIDDLEWARE} failed to get account id for unauthorized request warning"
	MessageFailedToGetAuthConfig = "{HORUSEC_MIDDLEWARE} grpc method to get auth config failed"
)
