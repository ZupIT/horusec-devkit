package database

import "github.com/ZupIT/horusec-devkit/pkg/services/database/response"

type IDatabaseWrite interface {
	StartTransaction() IDatabaseWrite
	RollbackTransaction() response.IResponse
	CommitTransaction() response.IResponse
	IsAvailable() bool
	Create(data interface{}, table string) response.IResponse
	CreateOrUpdate(entity interface{}, where map[string]interface{}, table string) response.IResponse
	Update(entity interface{}, where map[string]interface{}, table string) response.IResponse
	Delete(where map[string]interface{}, table string) response.IResponse
}
