package pkg

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
	"os"
	"time"
)

// TODO: Zerolog
// TODO: Store DB/AMQP connections at global static level.

func HandleCheckInterval() {
	log.Debug().Msg("check cycle started")
	// Get feeds from DB
	sources, db, err := getFeedSources()
	if err != nil {
		log.Error().Err(err).Msg("unable to get feed sources")
		return
	}
	_, ch, err := getAMQPChannel()
	if err != nil {
		log.Error().Err(err).Msg("unable to connect to RabbitMQ")
	}
	defer ch.Close()
	// Pull items from each feed
	for _, source := range sources {
		feed, err := getFeed(source.Url)
		if err != nil {
			fmt.Printf("unable to read feed: %w", err)
			log.Error().Err(err).Str("source", source.Url).Msg("unable to read feed")
			return
		}
		checkTime := time.Now()
		// Send items to MQ
		for _, item := range feed.Items {
			if item.PublishedParsed.After(source.LastChecked) {
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
		source.LastChecked = checkTime
		db.Save(&source)
	}
}
