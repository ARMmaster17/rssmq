package pkg

import (
	"gorm.io/gorm"
	"time"
)

type FeedSource struct {
	gorm.Model
	Url         string
	LastChecked time.Time
}
