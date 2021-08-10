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
