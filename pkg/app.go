package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
	"time"
)

// TODO: Store DB/AMQP connections at global static level.

type App struct {
	DB *gorm.DB
	Router *mux.Router
}

var TotalChecks = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "rss_checks_total",
		Help: "Number of RSS feed checks",
	},
)

var NewItems = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "rss_new_items_total",
		Help: "Number of new RSS items found",
	},
	[]string{"url"},
)

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
			fmt.Printf("unable to read feed: %w", err)
			log.Error().Err(err).Str("source", source.Url).Msg("unable to read feed")
			return
		}
		checkTime := time.Now()
		// Send items to MQ
		for _, item := range feed.Items {
			if item.PublishedParsed.After(source.LastChecked) {
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
					fmt.Printf("unable to send queue message for %s: %w", item.Link, err)
				}
			}
		}
		source.LastChecked = checkTime
		a.DB.Save(&source)
	}
}

func (a *App) respondWithError(w http.ResponseWriter) {
	a.respondWithErrorMessage(w, "Internal server error")
}

func (a *App) respondWithErrorMessage(w http.ResponseWriter, message string) {
	a.respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": message})
}

func (a *App) respondOKWithJSON(w http.ResponseWriter, payload interface{}) {
	a.respondWithJSON(w, http.StatusOK, payload)
}

func (a *App) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func (a *App) HandleGetFeeds(w http.ResponseWriter, r *http.Request) {
	var feeds []FeedSource
	result := a.DB.Find(&feeds)
	if result.Error != nil {
		log.Error().Stack().Err(result.Error).Msgf("unable to process request %s", r.RequestURI)
		a.respondWithErrorMessage(w, result.Error.Error())
		return
	}
	a.respondOKWithJSON(w, feeds)
}

func (a *App) HandleCreateFeed(w http.ResponseWriter, r *http.Request) {
	var feed FeedSource
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&feed); err != nil {
		a.respondWithErrorMessage(w, err.Error())
		return
	}
	result := a.DB.Create(&feed)
	if result.Error != nil {
		log.Error().Stack().Err(result.Error).Msgf("unable to process request %s", r.RequestURI)
		a.respondWithErrorMessage(w, result.Error.Error())
		return
	}
	a.respondOKWithJSON(w, nil)
}

func (a *App) HandleDeleteFeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fsId, err := strconv.ParseFloat(vars["id"], 64)
	if err != nil {
		a.respondWithErrorMessage(w, err.Error())
		return
	}
	result := a.DB.Delete(&FeedSource{}, fsId)
	if result.Error != nil {
		log.Error().Stack().Err(result.Error).Msgf("unable to process request %s", r.RequestURI)
		a.respondWithErrorMessage(w, result.Error.Error())
		return
	}
}