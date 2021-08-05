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

package swagger

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestNewSwagger(t *testing.T) {
	t.Run("should success create a new swagger service", func(t *testing.T) {
		assert.NotNil(t, NewSwagger(&chi.Mux{}, "9999"))
	})
}

func TestSetupSwagger(t *testing.T) {
	t.Run("should panic when invalid router", func(t *testing.T) {
		swaggerService := NewSwagger(&chi.Mux{}, "9999")
		assert.NotNil(t, swaggerService)

		assert.Panics(t, func() {
			swaggerService.SetupSwagger()
		})
	})
}

func TestGetSwaggerHost(t *testing.T) {
	t.Run("should success get swagger host", func(t *testing.T) {
		swaggerService := NewSwagger(&chi.Mux{}, "9999")
		assert.NotNil(t, swaggerService)

		swaggerHost := swaggerService.GetSwaggerHost()
		assert.NotEmpty(t, swaggerHost)
		assert.Equal(t, "localhost:9999", swaggerHost)
	})
}
