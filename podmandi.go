package podmandi

import (
	"encoding/json"
	"os"
)

// Manager is the podcast manager.
type Manager struct {
	// Podcasts are all podcasts that being managed.
	Podcasts []Podcast

	// dataFile location of persistent data file.
	dataFile string
}

// Option is the configuration function of Manager instance.
type Option func(m *Manager)

// WithDataFile makes the Manager instance to use a persistent data file.
func WithDataFile(df string) Option {
	return func(m *Manager) {
		m.dataFile = df
	}
}

// NewManager creates new Manager instance.
func NewManager(options ...Option) (*Manager, error) {
	m := &Manager{}
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
	m.Podcasts = append(m.Podcasts, Podcast{URL: url})
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
