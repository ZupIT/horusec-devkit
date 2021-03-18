package enums

import "errors"

const (
	FailedToConnectToDatabase = "{ERROR_DATABASE} failed to connect with postgres database"
	FailedToVerifyIsAvailable = "{ERROR_DATABASE} failed to get database while checking if is available"
)

var ErrNotFoundRecords = errors.New("{ERROR_DATABASE} database not found records")
