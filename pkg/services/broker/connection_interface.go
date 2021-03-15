package broker

import "github.com/streadway/amqp"

type iConnection interface {
	IsClosed() bool
	Channel() (*amqp.Channel, error)
	Close() error
}
