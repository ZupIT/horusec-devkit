package swagger

import (
	"testing"

	"github.com/go-chi/chi"
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
