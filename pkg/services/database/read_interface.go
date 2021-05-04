package database

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
)

type IDatabaseRead interface {
	IsAvailable() bool
	FindPreload(entityPointer interface{}, where map[string]interface{}, preloads []string,
		table string) response.IResponse
	Find(entityPointer interface{}, where map[string]interface{}, table string) response.IResponse
	First(entityPointer interface{}, where map[string]interface{}, table string) response.IResponse
	Raw(rawSQL string, entityPointer interface{}, values ...interface{}) response.IResponse
}
