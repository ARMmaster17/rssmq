package pkg

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

func getFeed(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("unable to read RSS feed: %w", err)
	}
	return feed, nil
}
