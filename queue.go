package amqp

import (
	"github.com/streadway/amqp"
)

type (
	Queue struct {
		Name        string
		Durable     bool
		AutoDelete  bool
		AutoAck     bool
		Exclusive   bool
		NoWait      bool
		Args        amqp.Table
		Passive     bool
		ConsumerTag string

		Bindings []QueueBinding

		q *amqp.Queue
	}

	QueueBinding struct {
		Name       string
		RoutingKey string
		Exchange   string

		NoWait bool
		Args   amqp.Table

		SvcHandler ServiceHandler
	}
)
