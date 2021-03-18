package health

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHealthCheckGrpcServer(t *testing.T) {
	t.Run("should success create a new service", func(t *testing.T) {
		assert.NotNil(t, NewHealthCheckGrpcServer())
	})
}

func TestCheck(t *testing.T) {
	t.Run("should success get check result", func(t *testing.T) {
		service := NewHealthCheckGrpcServer()

		response, err := service.Check(nil, nil)

		assert.NotNil(t, response)
		assert.NoError(t, err)
	})
}

func TestWatch(t *testing.T) {
	t.Run("should panic when no active server", func(t *testing.T) {
		service := NewHealthCheckGrpcServer()

		assert.Panics(t, func() {
			_ = service.Watch(nil, nil)
		})
	})
}
