package health

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestNewHealthCheckGrpcClient(t *testing.T) {
	t.Run("should success create a new health check service", func(t *testing.T) {
		service := NewHealthCheckGrpcClient(&grpc.ClientConn{})

		assert.NotNil(t, service)
	})
}

func TestIsAvailable(t *testing.T) {
	t.Run("should return true when healthy", func(t *testing.T) {
		connection, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
		assert.NoError(t, err)

		service := NewHealthCheckGrpcClient(connection)
		isAvailable, status := service.IsAvailable()

		assert.True(t, isAvailable)
		assert.NotEmpty(t, status)
	})

	t.Run("should return false when unhealthy", func(t *testing.T) {
		connection, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
		assert.NoError(t, err)

		err = connection.Close()
		assert.NoError(t, err)

		service := NewHealthCheckGrpcClient(connection)
		isAvailable, status := service.IsAvailable()

		assert.False(t, isAvailable)
		assert.NotEmpty(t, status)
	})
}
