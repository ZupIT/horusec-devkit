package database

import "github.com/ZupIT/horusec-devkit/pkg/services/database/response"

type IDatabaseWrite interface {
	StartTransaction() IDatabaseWrite
	RollbackTransaction() response.IResponse
	CommitTransaction() response.IResponse
	IsAvailable() bool
	Create(entityPointer interface{}, table string) response.IResponse
	CreateOrUpdate(entityPointer interface{}, where map[string]interface{}, table string) response.IResponse
	Update(entityPointer interface{}, where map[string]interface{}, table string) response.IResponse
	Delete(where map[string]interface{}, table string) response.IResponse
}
