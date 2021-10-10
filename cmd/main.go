package main

import (
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
	"rssmq/pkg"
	"time"
)

func main() {
	log.Info().Msg("rssmq starting up")
	log.Info().Msg("connecting to database using RSSMQ_DB_*")
	app := pkg.App{}
	err := app.Init()
	if err != nil {
		log.Fatal().Err(err).Msg("server failed to initialize")
	}
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minutes().Do(app.HandleCheckInterval)
	log.Info().Msg("starting check scheduler")
	s.StartAsync()
	log.Info().Msg("API is available on port 8080")
	log.Fatal().Err(app.StartAPIBlocking()).Msg("server failed")
}
