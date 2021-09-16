package pkg

import (
	"fmt"
	"github.com/streadway/amqp"
)

func HandleCheckInterval(ch *amqp.Channel, queueName string) {
	feed, err := getFeed("https://xkcd.com/atom.xml")
	if err != nil {
		fmt.Printf("Unable to read feed: %w", err)
	}
	// TODO: Compare publish time with last checked time.
	for _, i := range feed.Items {
		err = ch.Publish(
			"",
			queueName,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("Hello World"),
			},
		)
		if err != nil {
			fmt.Printf("Unable to send queue message for %s: %w", i.Title, err)
		}
	}
}
