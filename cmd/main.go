package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"rssmq/pkg"
	"time"
)

func main() {
	db, err := gorm.Open(postgres.Open(os.Getenv("RSSMQ_DB")))
	if err != nil {
		fmt.Printf("unable to connect to DB: %w", err)
		os.Exit(1)
	}
	err = db.AutoMigrate(&pkg.FeedSource{})
	if err != nil {
		fmt.Printf("unable to migrate DB: %w", err)
		os.Exit(1)
	}
	s := gocron.NewScheduler(time.UTC)
	s.Every(6).Hours().Do(pkg.HandleCheckInterval)
	s.StartBlocking()
}
