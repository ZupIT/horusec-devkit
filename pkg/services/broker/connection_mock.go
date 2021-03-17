package broker

import (
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/mock"

	mockUtils "github.com/ZupIT/horusec-devkit/pkg/utils/mock"
)

type connectionMock struct {
	mock.Mock
}

func (c *connectionMock) IsClosed() bool {
	args := c.MethodCalled("IsClosed")
	return args.Get(0).(bool)
}

func (c *connectionMock) Channel() (*amqp.Channel, error) {
	args := c.MethodCalled("Channel")
	return args.Get(0).(*amqp.Channel), mockUtils.ReturnNilOrError(args, 1)
}

func (c *connectionMock) Close() error {
	args := c.MethodCalled("Close")
	return mockUtils.ReturnNilOrError(args, 0)
}
