package main

import (
	"github.com/go-co-op/gocron"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"rssmq/pkg"
	"time"
)

func main() {
	log.Info().Msg("rssmq starting up")
	log.Info().Msg("connecting to database using RSSMQ_DB_*")
	app := pkg.App{}
	var err error
	app.DB, err = pkg.GetDB()
	if err != nil {
		log.Fatal().Err(err).Str("ENV:RSSMQ_DB", os.Getenv("RSSMQ_DB")).Msg("unable to connect to database")
	}
	err = app.DB.AutoMigrate(&pkg.FeedSource{})
	if err != nil {
		log.Fatal().Err(err).Str("ENV:RSSMQ_DB", os.Getenv("RSSMQ_DB")).Msg("unable to migrate database")
	}
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minutes().Do(app.HandleCheckInterval)
	log.Info().Msg("starting check scheduler")
	s.StartAsync()
	log.Info().Msg("Setting up prometheus")
	err = prometheus.Register(pkg.TotalChecks)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to initialize metrics")
	}
	err = prometheus.Register(pkg.NewItems)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to initialize metrics")
	}

	app.Router = mux.NewRouter()
	app.Router.Path("/metrics").Handler(promhttp.Handler())
	app.Router.HandleFunc("/api/feeds", app.HandleGetFeeds).Methods("GET")
	app.Router.HandleFunc("/api/feed/new", app.HandleCreateFeed).Methods("POST")
	app.Router.HandleFunc("/api/feed/{id:[0-9]+}/delete", app.HandleDeleteFeed).Methods("POST")
	app.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	log.Info().Msg("API is available on port 8080")
	log.Fatal().Err(http.ListenAndServe(":8080", app.Router)).Msg("server failed")
}
