package pkg

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/mmcdole/gofeed"
	"github.com/spf13/viper"
	"time"
)

func Run() {
	s := gocron.NewScheduler(time.UTC)
	lastCheckedTime := time.Now()
	fp := gofeed.NewParser()
	s.Every(viper.GetInt("checkIntervalHours")).Hours().Do(func() {
		for _, feed := range viper.GetStringSlice("feeds") {
			go func(feed string, lct time.Time) {
				fmt.Println("Fetching feed: " + feed)
				feedsData, err := fp.ParseURL(feed)
				if err != nil {
					fmt.Println(err)
					return
				}
				for _, item := range feedsData.Items {
					if item.PublishedParsed.After(lct) {
						fmt.Println("New item found: " + item.Title)
					}
				}
			}(feed, lastCheckedTime)
		}
		lastCheckedTime = time.Now()
	})
	s.StartBlocking()
}
