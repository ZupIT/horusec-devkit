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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type test struct{ test string }

func TestNewCache(t *testing.T) {
	t.Run("should success create a new cache", func(t *testing.T) {
		assert.NotNil(t, NewCache())
	})
}

func TestGet(t *testing.T) {
	t.Run("should success set and get data", func(t *testing.T) {
		cache := NewCache()

		cache.Set("test", "test", time.Minute*1)
		result := cache.Get("test")

		assert.NotNil(t, result)
		assert.Equal(t, "test", result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should success delete data", func(t *testing.T) {
		cache := NewCache()

		cache.Set("test", "test", time.Minute*1)
		cache.Delete("test")
		result := cache.Get("test")

		assert.Nil(t, result)
	})
}

func TestGetAndParse(t *testing.T) {
	toSet := &test{test: "test"}
	toParse := &test{test: "test"}

	t.Run("should success get and parse data", func(t *testing.T) {
		cache := NewCache()

		cache.Set("test", toSet, time.Minute*1)

		assert.NoError(t, cache.GetAndParse("test", toParse))
		assert.NotEmpty(t, toParse)
		assert.Equal(t, "test", toParse.test)
	})

	t.Run("should return error while marshall invalid data", func(t *testing.T) {
		cache := NewCache()

		cache.Set("test", make(chan string), time.Minute*1)

		assert.Error(t, cache.GetAndParse("test", toParse))
	})

	t.Run("should return error while unmarshal invalid data", func(t *testing.T) {
		cache := NewCache()

		cache.Set("test", toSet, time.Minute*1)

		assert.Error(t, cache.GetAndParse("test", ""))
	})
}

func TestGetString(t *testing.T) {
	t.Run("should success get and parse data", func(t *testing.T) {
		cache := NewCache()

		cache.Set("test", "test", time.Minute*1)

		result, err := cache.GetString("test")
		assert.NotEmpty(t, result)
		assert.NoError(t, err)
		assert.Equal(t, "test", result)
	})

	t.Run("should return error while marshall invalid data", func(t *testing.T) {
		cache := NewCache()

		cache.Set("test", make(chan string), time.Minute*1)

		result, err := cache.GetString("test")
		assert.Empty(t, result)
		assert.Error(t, err)
	})

	t.Run("should return error while unmarshal invalid data", func(t *testing.T) {
		cache := NewCache()

		cache.Set("test", test{}, time.Minute*1)

		result, err := cache.GetString("test")
		assert.Empty(t, result)
		assert.Error(t, err)
	})
}
