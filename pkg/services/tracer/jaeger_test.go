package tracer

import (
	"os"
	"testing"

	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/tracer/enums"
)

func TestNewJaeger(t *testing.T) {
	t.Run("should return a error when HORUSEC_JAGER_NAME is not setted", func(t *testing.T) {
		_, err := NewJaeger()
		assert.Error(t, err)
	})
	t.Run("should return a new Jaeger with default config", func(t *testing.T) {
		err := os.Setenv(enums.HorusecJaegerName, "test")
		assert.NoError(t, err)
		j, err := NewJaeger()
		assert.NoError(t, err)
		assert.Equal(t, "test", j.Name)
	})

}

func TestConfig(t *testing.T) {
	t.Run("should register a global tracer when config is called", func(t *testing.T) {
		err := os.Setenv(enums.HorusecJaegerName, "test2")
		assert.NoError(t, err)

		j, err := NewJaeger()
		assert.NoError(t, err)

		closer, err := j.Config(false)
		defer func() {
			err := closer.Close()
			assert.NoError(t, err)
		}()
		assert.NoError(t, err)
		assert.True(t, opentracing.IsGlobalTracerRegistered())
	})
	t.Run("should get a error when fromEnv method fails", func(t *testing.T) {
		err := os.Setenv(enums.HorusecJaegerName, "test2")
		assert.NoError(t, err)
		err = os.Setenv("JAEGER_RPC_METRICS", "test2")
		assert.NoError(t, err)
		j, err := NewJaeger()
		assert.NoError(t, err)

		_, err = j.Config(false)
		assert.Error(t, err)
		err = os.Setenv("JAEGER_RPC_METRICS", "false")
	})
}
