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
	"errors"
	"testing"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"
)

func getTestConfig() *config.Config {
	brokerConfig := &config.Config{}

	brokerConfig.SetHost("test")
	brokerConfig.SetPort("test")
	brokerConfig.SetUsername("test")
	brokerConfig.SetPassword("test")

	return brokerConfig
}

func testConsumer(_ packet.IPacket) {}

func TestNewBroker(t *testing.T) {
	t.Run("should return error when failed to connect", func(t *testing.T) {
		broker, err := NewBroker(getTestConfig())

		assert.Nil(t, broker)
		assert.Error(t, err)
	})

	t.Run("should return error when invalid config", func(t *testing.T) {
		broker, err := NewBroker(&config.Config{})

		assert.Nil(t, broker)
		assert.Error(t, err)
	})
}

func TestSetupChannel(t *testing.T) {
	t.Run("should success setup channel without errors", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		connectionMock.On("IsClosed").Return(false)
		channelMock.On("Flow").Return(nil)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.NoError(t, broker.setupChannel())
	})

	t.Run("should return error when failed to setup connection", func(t *testing.T) {
		broker := &Broker{
			connection: nil,
			channel:    nil,
			config:     getTestConfig(),
		}

		assert.Error(t, broker.setupChannel())
	})
}

func TestVerifyEmptyChannelAndSetFlow(t *testing.T) {
	t.Run("should success verify no empty or nil channel and set flow", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		channelMock.On("Flow").Return(nil)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.NoError(t, broker.verifyEmptyChannelAndSetFlow())
	})

	t.Run("should return error when failed to set flow", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		channelMock.On("Flow").Return(errors.New("test"))
		connectionMock.On("Channel").Return(&amqp.Channel{}, errors.New("test"))

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.Error(t, broker.verifyEmptyChannelAndSetFlow())
	})

	t.Run("should panic when trying to set channel with nil connection", func(t *testing.T) {
		broker := &Broker{
			connection: nil,
			channel:    nil,
			config:     getTestConfig(),
		}

		assert.Panics(t, func() {
			_ = broker.verifyEmptyChannelAndSetFlow()
		})
	})
}

func TestIsAvailable(t *testing.T) {
	t.Run("should return true when everything it is ok", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		connectionMock.On("Channel").Return(&amqp.Channel{}, nil)
		connectionMock.On("IsClosed").Return(false)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.True(t, broker.IsAvailable())
	})

	t.Run("should return false when failed to setup connection", func(t *testing.T) {
		broker := &Broker{
			connection: nil,
			channel:    nil,
			config:     getTestConfig(),
		}

		assert.False(t, broker.IsAvailable())
	})
}

func TestIsNotClosedOrNil(t *testing.T) {
	t.Run("should return true when everything it is ok", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		connectionMock.On("IsClosed").Return(false)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.True(t, broker.isNotClosedOrNil())
	})

	t.Run("should return false when closed connection", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		connectionMock.On("IsClosed").Return(true)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.False(t, broker.isNotClosedOrNil())
	})

	t.Run("should return false when nil connection", func(t *testing.T) {
		broker := &Broker{
			connection: nil,
			channel:    nil,
			config:     getTestConfig(),
		}

		assert.False(t, broker.isNotClosedOrNil())
	})
}

func TestClose(t *testing.T) {
	t.Run("should success close connection with no errors", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		connectionMock.On("Close").Return(nil)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.NoError(t, broker.Close())
	})
}

func TestPublish(t *testing.T) {
	t.Run("should success publish packet without exchange and no errors", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		channelMock.On("Publish").Return(nil)
		channelMock.On("Flow").Return(nil)
		connectionMock.On("IsClosed").Return(false)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.NoError(t, broker.Publish("", "", "", []byte("")))
	})

	t.Run("should success publish packet with exchange and no errors", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		channelMock.On("Publish").Return(nil)
		channelMock.On("Flow").Return(nil)
		channelMock.On("ExchangeDeclare").Return(nil)
		connectionMock.On("IsClosed").Return(false)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.NoError(t, broker.Publish("", "test", "test", []byte("")))
	})

	t.Run("should return error when failed to declare exchange", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		channelMock.On("Flow").Return(nil)
		channelMock.On("ExchangeDeclare").Return(errors.New("test"))
		connectionMock.On("IsClosed").Return(false)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.Error(t, broker.Publish("", "test", "test", []byte("")))
	})

	t.Run("should return error when failed setup channel", func(t *testing.T) {
		broker := &Broker{
			connection: nil,
			channel:    nil,
			config:     getTestConfig(),
		}

		assert.Error(t, broker.Publish("", "", "", []byte("")))
	})
}

func TestConsume(t *testing.T) {
	t.Run("should success start a consumer without errors", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		channelMock.On("Flow").Return(nil)
		channelMock.On("Qos").Return(nil)
		channelMock.On("QueueDeclare").Return(amqp.Queue{}, nil)
		channelMock.On("Consume").Return(make(<-chan amqp.Delivery), nil)
		connectionMock.On("IsClosed").Return(false)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.NotPanics(t, func() {
			go broker.Consume("", "", "", testConsumer)
		})
	})

	t.Run("should panic when failed to consume", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		channelMock.On("Flow").Return(nil)
		channelMock.On("Qos").Return(nil)
		channelMock.On("QueueDeclare").Return(amqp.Queue{}, nil)
		channelMock.On("Consume").Return(make(<-chan amqp.Delivery), errors.New("test"))
		connectionMock.On("IsClosed").Return(false)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.Panics(t, func() {
			broker.Consume("", "", "", testConsumer)
		})
	})

	t.Run("should panic when failed to queue bind", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		channelMock.On("Flow").Return(nil)
		channelMock.On("Qos").Return(nil)
		channelMock.On("QueueDeclare").Return(amqp.Queue{}, nil)
		channelMock.On("ExchangeDeclare").Return(nil)
		channelMock.On("QueueBind").Return(errors.New("test"))
		connectionMock.On("IsClosed").Return(false)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.Panics(t, func() {
			broker.Consume("", "test", "test", testConsumer)
		})
	})

	t.Run("should panic when failed to exchange declare", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		channelMock.On("Flow").Return(nil)
		channelMock.On("Qos").Return(nil)
		channelMock.On("QueueDeclare").Return(amqp.Queue{}, nil)
		channelMock.On("ExchangeDeclare").Return(errors.New("test"))
		connectionMock.On("IsClosed").Return(false)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.Panics(t, func() {
			broker.Consume("", "test", "test", testConsumer)
		})
	})

	t.Run("should panic when failed to queue declare", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		channelMock.On("Flow").Return(nil)
		channelMock.On("Qos").Return(nil)
		channelMock.On("QueueDeclare").Return(amqp.Queue{}, errors.New("test"))
		connectionMock.On("IsClosed").Return(false)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.Panics(t, func() {
			broker.Consume("", "", "", testConsumer)
		})
	})

	t.Run("should panic when failed to set consumer prefetch", func(t *testing.T) {
		connectionMock := &connectionMock{}
		channelMock := &channelMock{}

		channelMock.On("Flow").Return(nil)
		channelMock.On("Qos").Return(errors.New("test"))
		connectionMock.On("IsClosed").Return(false)

		broker := &Broker{
			connection: connectionMock,
			channel:    channelMock,
			config:     getTestConfig(),
		}

		assert.Panics(t, func() {
			broker.Consume("", "", "", testConsumer)
		})
	})

	t.Run("should panic when failed to setup channel", func(t *testing.T) {
		broker := &Broker{
			connection: nil,
			channel:    nil,
			config:     getTestConfig(),
		}

		assert.Panics(t, func() {
			broker.Consume("", "", "", testConsumer)
		})
	})
}
