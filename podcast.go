package podmandi

import "github.com/mmcdole/gofeed"

// Podcast represents a podcast.
type Podcast struct {
	// URL of the podcast
	URL string `json:"url"`
	// Feed is the content of podcast feed
	Feed *gofeed.Feed
}
