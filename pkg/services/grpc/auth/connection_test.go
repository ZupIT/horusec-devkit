package auth

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthGRPCConnection(t *testing.T) {
	t.Run("should success make connection without certs", func(t *testing.T) {
		_ = os.Setenv("HORUSEC_GRPC_USE_CERTS", "false")

		assert.NotNil(t, NewAuthGRPCConnection())
	})

	t.Run("should panic when failed to make connection with certs", func(t *testing.T) {
		_ = os.Setenv("HORUSEC_GRPC_USE_CERTS", "true")

		assert.Panics(t, func() {
			NewAuthGRPCConnection()
		})
	})
}
