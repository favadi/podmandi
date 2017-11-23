package podmandi

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

// feedParser is the podcast URL parser.
type feedParser interface {
	ParseURL(feedURL string) (feed *gofeed.Feed, err error)
}

// Manager is the podcast manager.
type Manager struct {
	// Podcasts are all podcasts that being managed.
	Podcasts []Podcast

	// dataFile location of persistent data file.
	dataFile string

	// parser is the podcast URL parser.
	parser feedParser
}

// Option is the configuration function of Manager instance.
type Option func(m *Manager)

// WithDataFile makes the Manager instance use a persistent data file.
func WithDataFile(df string) Option {
	return func(m *Manager) {
		m.dataFile = df
	}
}

// NewManager creates new Manager instance.
func NewManager(parser feedParser, options ...Option) (*Manager, error) {
	m := &Manager{
		parser: parser,
	}
	for _, option := range options {
		option(m)
	}

	if len(m.dataFile) == 0 {
		return m, nil
	}

	fp, err := os.Open(m.dataFile)
	if os.IsNotExist(err) {
		return m, nil
	}
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	var podcasts []Podcast
	decoder := json.NewDecoder(fp)
	if err = decoder.Decode(&podcasts); err != nil {
		return nil, err
	}
	m.Podcasts = podcasts
	return m, nil
}

// Add adds a new podcast.
func (m *Manager) Add(url string) error {
	for _, pod := range m.Podcasts {
		if pod.URL == url {
			return fmt.Errorf("podcast: %q is already exists", url)
		}
	}

	feed, err := m.parser.ParseURL(url)
	if err != nil {
		return err
	}

	m.Podcasts = append(m.Podcasts, Podcast{URL: url, Feed: feed})
	return m.save()
}

// save saves the application data to persistent file.
func (m *Manager) save() error {
	if len(m.dataFile) == 0 {
		return nil
	}

	fp, err := os.Create(m.dataFile)
	if err != nil {
		return err
	}
	defer fp.Close()

	encoder := json.NewEncoder(fp)
	return encoder.Encode(m.Podcasts)
}
