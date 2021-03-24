// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package http

import (
	"testing"

	"github.com/go-chi/cors"
	"github.com/stretchr/testify/assert"
)

func TestNewHTTPRouter(t *testing.T) {
	t.Run("should return a new router service with default config", func(t *testing.T) {
		assert.NotNil(t, NewHTTPRouter(&cors.Options{}, "8000"))
	})
}

func TestGetMux(t *testing.T) {
	t.Run("should return a chi mux instance", func(t *testing.T) {
		router := NewHTTPRouter(&cors.Options{}, "8000")

		assert.NotNil(t, router)
		assert.NotNil(t, router.GetMux())
	})
}

func TestSetTimeout(t *testing.T) {
	t.Run("should return a chi router interface", func(t *testing.T) {
		router := NewHTTPRouter(&cors.Options{}, "8000")

		assert.NotNil(t, router)
		assert.NotNil(t, router.Route("/test", nil))
	})
}
