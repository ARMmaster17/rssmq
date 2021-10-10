package pkg

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"gorm.io/gorm"
)

func getFeed(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("unable to read RSS feed: %w", err)
	}
	return feed, nil
}

func getFeedSources(db *gorm.DB) ([]FeedSource, error) {

	var feedSources []FeedSource
	result := db.Find(&feedSources)
	if result.Error != nil {
		return nil, fmt.Errorf("unable to get feed sources: %w", result.Error)
	}
	return feedSources, nil
}
