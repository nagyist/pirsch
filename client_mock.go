package pirsch

import (
	"sync"
	"time"
)

// ClientMock is a mock Store implementation.
type ClientMock struct {
	Hits          []Hit
	Events        []Event
	UserAgents    []UserAgent
	ReturnSession *Hit
	m             sync.Mutex
}

// NewMockClient returns a new mock client.
func NewMockClient() *ClientMock {
	return &ClientMock{
		Hits:       make([]Hit, 0),
		Events:     make([]Event, 0),
		UserAgents: make([]UserAgent, 0),
	}
}

// SaveHits implements the Store interface.
func (client *ClientMock) SaveHits(hits []Hit) error {
	client.m.Lock()
	defer client.m.Unlock()
	client.Hits = append(client.Hits, hits...)
	return nil
}

// SaveEvents implements the Store interface.
func (client *ClientMock) SaveEvents(events []Event) error {
	client.m.Lock()
	defer client.m.Unlock()
	client.Events = append(client.Events, events...)
	return nil
}

// SaveUserAgents implements the Store interface.
func (client *ClientMock) SaveUserAgents(userAgents []UserAgent) error {
	client.m.Lock()
	defer client.m.Unlock()
	client.UserAgents = append(client.UserAgents, userAgents...)
	return nil
}

// Session implements the Store interface.
func (client *ClientMock) Session(clientID, fingerprint uint64, maxAge time.Time) (*Hit, error) {
	if client.ReturnSession != nil {
		return client.ReturnSession, nil
	}

	return nil, nil
}

// Count implements the Store interface.
func (client *ClientMock) Count(query string, args ...interface{}) (int, error) {
	return 0, nil
}

// Get implements the Store interface.
func (client *ClientMock) Get(result interface{}, query string, args ...interface{}) error {
	return nil
}

// Select implements the Store interface.
func (client *ClientMock) Select(results interface{}, query string, args ...interface{}) error {
	return nil
}