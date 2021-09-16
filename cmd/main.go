package main

import (
	"github.com/go-co-op/gocron"
	"time"
)

func main() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(6).Hours().Do(nil)
	s.StartBlocking()
}
