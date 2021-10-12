package pkg

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

type App struct {
	DB *gorm.DB
	Router *mux.Router
	CORS []handlers.CORSOption
}

func (a *App) Init() error {
	var err error
	a.DB, err = GetDB()
	if err != nil {
		return err
	}
	err = a.DB.AutoMigrate(&FeedSource{})
	if err != nil {
		return err
	}

	err = a.registerHTTPRoutes()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to initialize HTTP server")
	}
	return nil
}

func (a *App) StartAPIBlocking() error {
	return http.ListenAndServe(":8080", handlers.CORS(a.CORS...)(a.Router))
}

func (a *App) HandleCheckInterval() {
	log.Debug().Msg("check cycle started")
	// Get feeds from DB
	sources, err := getFeedSources(a.DB)
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
		TotalChecks.Inc()
		feed, err := getFeed(source.Url)
		if err != nil {
			fmt.Printf("unable to read feed: %s", err.Error())
			log.Error().Err(err).Str("source", source.Url).Msg("unable to read feed")
			return
		}
		checkTime := time.Now()
		// Send items to MQ
		for _, item := range feed.Items {
			if feedItemIsNew(source.LastChecked, item) {
				NewItems.WithLabelValues(source.Url).Inc()
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
					fmt.Printf("unable to send queue message for %s: %s", item.Link, err.Error())
				}
			}
		}
		source.LastChecked = checkTime
		a.DB.Save(&source)
	}
}