package database

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
)

type IDatabaseRead interface {
	IsAvailable() bool
	Find(entityPointer interface{}, where map[string]interface{}, table string) response.IResponse
	First(entityPointer interface{}, where map[string]interface{}, table string) response.IResponse
	Raw(rawSQL string, entityPointer interface{}) response.IResponse
}
