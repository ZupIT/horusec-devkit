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

package broker

import (
	"os"
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"
	"github.com/stretchr/testify/assert"
)

func setDefaultEnvVars() {
	_ = os.Setenv("HORUSEC_BROKER_HOST", "")
	_ = os.Setenv("HORUSEC_BROKER_PORT", "")
	_ = os.Setenv("HORUSEC_BROKER_USERNAME", "")
	_ = os.Setenv("HORUSEC_BROKER_PASSWORD", "")
}

func TestNewBroker(t *testing.T) {
	t.Run("should success return a new broker without errors", func(t *testing.T) {
		setDefaultEnvVars()

		broker, err := NewBroker(config.NewBrokerConfig())
		assert.NotNil(t, broker)
		assert.NoError(t, err)
	})

	t.Run("should return error when invalid config", func(t *testing.T) {
		_, err := NewBroker(&config.Config{})
		assert.Error(t, err)
	})

	t.Run("should return error when making connection", func(t *testing.T) {
		_ = os.Setenv("HORUSEC_BROKER_HOST", "test")
		_ = os.Setenv("HORUSEC_BROKER_PORT", "test")
		_ = os.Setenv("HORUSEC_BROKER_USERNAME", "test")
		_ = os.Setenv("HORUSEC_BROKER_PASSWORD", "test")

		_, err := NewBroker(config.NewBrokerConfig())
		assert.Error(t, err)
	})
}

func TestIsAvailable(t *testing.T) {
	setDefaultEnvVars()

	t.Run("should return true when its available", func(t *testing.T) {
		broker, _ := NewBroker(config.NewBrokerConfig())
		assert.True(t, broker.IsAvailable())
	})
}

func TestClose(t *testing.T) {
	t.Run("should close connection without errors", func(t *testing.T) {
		broker, _ := NewBroker(config.NewBrokerConfig())
		assert.NoError(t, broker.Close())
	})
}

func TestPublish(t *testing.T) {
	t.Run("should publish packet without errors and no exchange", func(t *testing.T) {
		broker, _ := NewBroker(config.NewBrokerConfig())

		err := broker.Publish("test", "", "", []byte("test"))
		assert.NoError(t, err)
	})

	t.Run("should publish packet without errors with exchange", func(t *testing.T) {
		broker, _ := NewBroker(config.NewBrokerConfig())

		err := broker.Publish("test", "test", "topic", []byte("test"))
		assert.NoError(t, err)
	})
}

func TestIsAvailableMock(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		mock := &Mock{}

		mock.On("IsAvailable").Return(true)

		result := mock.IsAvailable()
		assert.True(t, result)
	})
}

func TestPublishMock(t *testing.T) {
	t.Run("should return no error", func(t *testing.T) {
		mock := &Mock{}

		mock.On("Publish").Return(nil)

		result := mock.Publish("", "", "", nil)
		assert.NoError(t, result)
	})
}

func TestConsumeMock(t *testing.T) {
	t.Run("should not panic", func(t *testing.T) {
		mock := &Mock{}

		mock.On("Consume")

		assert.NotPanics(t, func() {
			mock.Consume("", "", "", testConsumer)
		})
	})
}

func testConsumer(_ packet.IPacket) {}

func TestCloseMock(t *testing.T) {
	t.Run("should return no error", func(t *testing.T) {
		mock := &Mock{}

		mock.On("Close").Return(nil)

		result := mock.Close()
		assert.NoError(t, result)
	})
}
