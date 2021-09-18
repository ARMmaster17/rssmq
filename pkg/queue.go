package pkg

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

func getAMQPChannel() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(os.Getenv("RSSMQ_MQ_URL"))
	if err != nil {
		return nil, nil, fmt.Errorf("unable to dial RabbitMQ: %w", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to establish MQ channel: %w", err)
	}
	_, err = ch.QueueDeclare(
		os.Getenv("RSSMQ_MQ_QUEUE"),
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to establish MQ queue: %w", err)
	}
	return conn, ch, nil
}
