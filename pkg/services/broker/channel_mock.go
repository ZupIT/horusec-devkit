package broker

import (
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/mock"

	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
)

type channelMock struct {
	mock.Mock
}

func (c *channelMock) ExchangeDeclare(_, _ string, _, _, _, _ bool, _ amqp.Table) error {
	args := c.MethodCalled("ExchangeDeclare")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (c *channelMock) Publish(_, _ string, _, _ bool, _ amqp.Publishing) error {
	args := c.MethodCalled("Publish")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (c *channelMock) QueueDeclare(_ string, _, _, _, _ bool, _ amqp.Table) (amqp.Queue, error) {
	args := c.MethodCalled("QueueDeclare")
	return args.Get(0).(amqp.Queue), mockUtils.ReturnNilOrError(args, 1)
}

func (c *channelMock) Flow(_ bool) error {
	args := c.MethodCalled("Flow")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (c *channelMock) Consume(_, _ string, _, _, _, _ bool, _ amqp.Table) (<-chan amqp.Delivery, error) {
	args := c.MethodCalled("Consume")
	return args.Get(0).(<-chan amqp.Delivery), mockUtils.ReturnNilOrError(args, 1)
}

func (c *channelMock) Qos(_, _ int, _ bool) error {
	args := c.MethodCalled("Qos")
	return mockUtils.ReturnNilOrError(args, 0)
}

func (c *channelMock) QueueBind(_, _, _ string, _ bool, _ amqp.Table) error {
	args := c.MethodCalled("QueueBind")
	return mockUtils.ReturnNilOrError(args, 0)
}
