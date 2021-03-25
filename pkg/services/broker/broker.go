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
	"github.com/pkg/errors"
	"github.com/streadway/amqp"

	"github.com/ZupIT/horusec-devkit/pkg/services/app"
	brokerConfig "github.com/ZupIT/horusec-devkit/pkg/services/broker/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/broker/enums"
	brokerPacket "github.com/ZupIT/horusec-devkit/pkg/services/broker/packet"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

type IBroker interface {
	IsAvailable() bool
	Consume(queue, exchange, exchangeKind string, handler func(packet brokerPacket.IPacket))
	Publish(queue, exchange, exchangeKind string, body []byte) error
	Close() error
}

type Broker struct {
	connection iConnection
	channel    iChannel
	config     brokerConfig.IConfig
}

func NewBroker(config brokerConfig.IConfig, appConfig app.IConfig) (IBroker, error) {
	if appConfig.IsBrokerDisabled() {
		return nil, nil
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	broker := &Broker{config: config}
	if err := broker.setupConnection(); err != nil {
		return nil, errors.Wrap(err, enums.MessageFailedConnectBroker)
	}

	return broker, broker.setupChannel()
}

func (b *Broker) setupConnection() (err error) {
	if b.isEmptyOrNilConnection() {
		b.connection, err = b.makeConnection()
	}

	if err != nil || b.connection.IsClosed() {
		b.connection, err = b.makeConnection()
	}

	return err
}

func (b *Broker) isEmptyOrNilConnection() bool {
	return b.connection == nil || b.connection == (&amqp.Connection{})
}

func (b *Broker) makeConnection() (*amqp.Connection, error) {
	return amqp.Dial(b.config.GetConnectionString())
}

func (b *Broker) setupChannel() (channelErr error) {
	if err := b.setupConnection(); err != nil {
		return err
	}

	return b.verifyEmptyChannelAndSetFlow()
}

func (b *Broker) verifyEmptyChannelAndSetFlow() (channelErr error) {
	if b.isEmptyOrNilChannel() {
		b.channel, channelErr = b.connection.Channel()
	}

	if err := b.channel.Flow(true); err != nil {
		b.channel, channelErr = b.connection.Channel()
	}

	return channelErr
}

func (b *Broker) isEmptyOrNilChannel() bool {
	return b.channel == nil || b.channel == (&amqp.Channel{})
}

func (b *Broker) IsAvailable() bool {
	if err := b.setupConnection(); err != nil {
		return false
	}

	return b.isNotClosedOrNil()
}

func (b *Broker) isNotClosedOrNil() bool {
	if b.isEmptyOrNilConnection() {
		return false
	}

	return !b.connection.IsClosed()
}

func (b *Broker) Close() error {
	return b.connection.Close()
}

func (b *Broker) publish(queue string, data []byte, exchange string) error {
	packet := amqp.Publishing{
		ContentType: "text/plain",
		Body:        data,
	}

	return b.channel.Publish(exchange, queue, false, false, packet)
}

func (b *Broker) exchangeDeclare(exchange, exchangeKind string) error {
	if exchange == "" || exchangeKind == "" {
		return nil
	}

	return b.channel.ExchangeDeclare(exchange, exchangeKind, true, false, false,
		false, nil)
}

func (b *Broker) Publish(queue, exchange, exchangeKind string, body []byte) error {
	if err := b.setupChannel(); err != nil {
		logger.LogError(enums.MessageFailedCreateChannelPublish, err)
		return err
	}

	if err := b.exchangeDeclare(exchange, exchangeKind); err != nil {
		logger.LogError(enums.MessageFailedDeclareExchangePublish, err)
		return err
	}

	return b.publish(queue, body, exchange)
}

func (b *Broker) Consume(queue, exchange, exchangeKing string, handler func(packet brokerPacket.IPacket)) {
	for {
		if err := b.setupChannel(); err != nil {
			logger.LogPanic(enums.MessageFailedCreateChannelConsume, err)
		}

		b.setConsumerPrefetch()
		b.declareQueueAndBind(queue, exchange, exchangeKing)
		b.handleDeliveries(queue, handler)
	}
}

func (b *Broker) declareQueueAndBind(queue, exchange, exchangeKing string) {
	if _, err := b.channel.QueueDeclare(queue, true, false, false,
		false, nil); err != nil {
		logger.LogPanic(enums.MessageFailedCreateQueueConsume, err)
	}

	if exchange != "" && exchangeKing != "" {
		b.declareExchangeAndBind(queue, exchange, exchangeKing)
	}
}

func (b *Broker) handleDeliveries(queue string, handler func(packet brokerPacket.IPacket)) {
	deliveries, err := b.channel.Consume(queue, "", false, false, false,
		false, nil)
	if err != nil {
		logger.LogPanic(enums.MessageFailedConsumeHandlingDelivery, err)
	}

	for delivery := range deliveries {
		message := delivery
		handler(brokerPacket.NewPacket(&message))
	}
}

func (b *Broker) setConsumerPrefetch() {
	if err := b.channel.Qos(1, 0, false); err != nil {
		logger.LogPanic(enums.MessageFailedSetConsumerPrefetch, err)
	}
}

func (b *Broker) declareExchangeAndBind(queue, exchange, exchangeKing string) {
	if err := b.exchangeDeclare(exchange, exchangeKing); err != nil {
		logger.LogPanic(enums.MessageFailedToDeclareExchangeQueue, err)
	}

	if err := b.channel.QueueBind(queue, "", exchange, false, nil); err != nil {
		logger.LogPanic(enums.MessageFailedBindQueueConsume, err)
	}
}
