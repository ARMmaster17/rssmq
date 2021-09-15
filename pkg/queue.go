package pkg

import (
	"fmt"
	"github.com/streadway/amqp"
)

func Connect(uri string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to RabbitMQ: %w", err)
	}
	return conn, nil
}
