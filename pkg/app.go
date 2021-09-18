package pkg

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

// TODO: Zerolog
// TODO: Update last-checked timestamp.
// TODO: Check publish timestamp against last check timestamp.
// TODO: Only update last-checked timestamp if successful.
// TODO: Store DB/AMQP connections at global static level.

func HandleCheckInterval() {
	// Get feeds from DB
	sources, err := getFeedSources()
	if err != nil {
		fmt.Printf("unable to get feed sources: %w", err)
		return
	}
	conn, ch, err := getAMQPChannel()
	if err != nil {
		fmt.Printf("unable to connect to RabbitMQ: %w", err)
	}
	defer conn.Close()
	defer ch.Close()
	// Pull items from each feed
	for _, source := range sources {
		feed, err := getFeed(source.Url)
		if err != nil {
			fmt.Printf("unable to read feed: %w", err)
			return
		}
		// Send items to MQ
		for _, item := range feed.Items {
			err = ch.Publish(
				"",
				os.Getenv("RSSMQ_MQ_QUEUE"),
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(item.Link),
				},
			)
			if err != nil {
				fmt.Printf("unable to send queue message for %s: %w", item.Link, err)
			}
		}
	}
}
