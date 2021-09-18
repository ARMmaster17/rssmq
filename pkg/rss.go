package pkg

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func getFeed(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("unable to read RSS feed: %w", err)
	}
	return feed, nil
}

func getFeedSources() ([]FeedSource, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("RSSMQ_DB")))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to DB: %w", err)
	}
	var feedSources []FeedSource
	result := db.Find(&feedSources)
	if result.Error != nil {
		return nil, fmt.Errorf("unable to get feed sources: %w", err)
	}
	return feedSources, nil
}
