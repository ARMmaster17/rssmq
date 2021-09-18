package main

import (
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	s.StartBlocking()
}
