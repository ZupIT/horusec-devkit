package cache

import (
	"encoding/json"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/ZupIT/horusec-devkit/pkg/services/cache/enums"
)

type ICache interface {
	Get(key string) interface{}
	GetAndParse(key string, entityPointer interface{}) error
	GetString(key string) (result string, err error)
	Delete(key string)
	Set(key string, value interface{}, duration time.Duration)
}

type Cache struct {
	cache *cache.Cache
}

func NewCache() ICache {
	return &Cache{
		cache: cache.New(enums.DefaultExpirationTime, enums.DefaultCheckExpiredTime),
	}
}

func (c *Cache) Get(key string) interface{} {
	data, _ := c.cache.Get(key)
	return data
}

func (c *Cache) Delete(key string) {
	c.cache.Delete(key)
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.cache.Set(key, value, duration)
}

func (c *Cache) GetAndParse(key string, entityPointer interface{}) error {
	data, _ := c.cache.Get(key)

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, entityPointer)
}

func (c *Cache) GetString(key string) (result string, err error) {
	data, _ := c.cache.Get(key)

	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return result, json.Unmarshal(bytes, &result)
}
