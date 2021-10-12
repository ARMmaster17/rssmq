package pkg

import (
	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCheckItemIsNew(t *testing.T) {
	currentTime := time.Now()
	laterTime := currentTime.Add(2 * time.Hour)
	feedItem := gofeed.Item{
		PublishedParsed: &laterTime,
	}
	assert.True(t, feedItemIsNew(currentTime, &feedItem))
}
