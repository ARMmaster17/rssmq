package main

import (
	"github.com/go-co-op/gocron"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"rssmq/pkg"
	"time"
)

func main() {
	log.Info().Msg("rssmq starting up")
	log.Info().Msg("connecting to database using RSSMQ_DB")
	db, err := gorm.Open(postgres.Open(os.Getenv("RSSMQ_DB")))
	if err != nil {
		log.Fatal().Err(err).Str("ENV:RSSMQ_DB", os.Getenv("RSSMQ_DB")).Msg("unable to connect to database")
	}
	err = db.AutoMigrate(&pkg.FeedSource{})
	if err != nil {
		log.Fatal().Err(err).Str("ENV:RSSMQ_DB", os.Getenv("RSSMQ_DB")).Msg("unable to migrate database")
	}
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minutes().Do(pkg.HandleCheckInterval)
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

	r := mux.NewRouter()
	r.Path("/").Handler(http.FileServer(http.Dir("./static/")))
	r.Path("/metrics").Handler(promhttp.Handler())
	log.Info().Msg("API is available on port 8080")
	log.Fatal().Err(http.ListenAndServe(":8080", r)).Msg("server failed")
}
