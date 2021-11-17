// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
