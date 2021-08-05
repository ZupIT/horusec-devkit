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

package router

import (
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/tracer"

	"github.com/go-chi/cors"
)

var router = NewHTTPRouter(&cors.Options{}, "8000", tracer.Jaeger{Name: "test"})

func TestNewHTTPRouter(t *testing.T) {
	t.Run("should return a new router service with default config", func(t *testing.T) {
		assert.NotNil(t, router)
		err := router.CloseJaeger()
		assert.Nil(t, err)
	})
}

func TestGetMux(t *testing.T) {
	t.Run("should return a chi mux instance", func(t *testing.T) {
		assert.NotNil(t, router)
		assert.NotNil(t, router.GetMux())
	})
}

func TestSetTimeout(t *testing.T) {
	t.Run("should return a chi router interface", func(t *testing.T) {
		assert.NotNil(t, router)
		assert.NotNil(t, router.Route("/test", func(router2 chi.Router) {}))
	})
}

func TestGetPort(t *testing.T) {
	t.Run("should success get router server port", func(t *testing.T) {
		assert.Equal(t, "8000", router.GetPort())
	})
}

func TestListenAndServe(t *testing.T) {
	t.Run("should panic when failed to serve", func(t *testing.T) {
		router := NewHTTPRouter(&cors.Options{}, "test", tracer.Jaeger{})
		assert.Panics(t, func() {
			router.ListenAndServe()
		})
	})
}
