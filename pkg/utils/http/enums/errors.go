package enums

import "errors"

var (
	ErrorBrokerIsNotHealth    = errors.New("{ERROR_HTTP} broker is not health")
	ErrorDatabaseIsNotHealth  = errors.New("{ERROR_HTTP} database is not health")
	ErrorGrpcIsNotHealth      = errors.New("{ERROR_HTTP} grpc is not health")
	ErrorGenericInternalError = errors.New("{ERROR_HTTP} something went wrong, sorry for the inconvenience")
)
