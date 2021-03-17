package database

import (
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
)

type IDatabaseRead interface {
	IsAvailable() bool
	Find(entity interface{}, where map[string]interface{}, table string) response.IResponse
	First(entity interface{}, where map[string]interface{}, table string) response.IResponse
	Raw(rawSQL string, entity interface{}) response.IResponse
}
