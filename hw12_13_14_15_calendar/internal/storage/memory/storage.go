package memorystorage

import (
	"github.com/gen-maksim/otus-hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"log"
	"sync"
	"time"
)

type Storage struct {
	collection map[string]storage.Event
	mu         sync.RWMutex //nolint:unused
}

func New() *Storage {
	return &Storage{collection: make(map[string]storage.Event)}
}

func (s *Storage) Store(event storage.Event) storage.Event {
	s.mu.Lock()
	defer s.mu.Unlock()
	if event.StartTime.After(event.EndTime) {
		log.Println("Start time is after end time")
		return storage.Event{}
	}

	id := uuid.New()
	event.ID = id.String()
	s.collection[event.ID] = event

	return event
}

func (s *Storage) Update(id string, event storage.Event) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if event.StartTime.After(event.EndTime) {
		log.Println("Start time is after end time")
		return
	}

	_, ok := s.collection[id]
	if !ok {
		log.Println("id is not exist")
		return
	}

	s.collection[id] = event
}

func (s *Storage) Delete(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.collection, id)
}

func (s *Storage) ListDay(date time.Time) []storage.Event {
	startSearch := date.Truncate(24 * time.Hour)
	endSearch := startSearch.AddDate(0, 0, 1).Add(-time.Second)

	return s.listPeriod(startSearch, endSearch)
}

func (s *Storage) ListWeek(date time.Time) []storage.Event {
	startSearch := date.Truncate(24 * time.Hour * 7)
	endSearch := startSearch.AddDate(0, 0, 7).Add(-time.Second)

	return s.listPeriod(startSearch, endSearch)
}

func (s *Storage) ListMonth(date time.Time) []storage.Event {
	startSearch := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
	endSearch := startSearch.AddDate(0, 1, 0).Add(-time.Second)

	return s.listPeriod(startSearch, endSearch)
}

func (s *Storage) listPeriod(startSearch, endSearch time.Time) []storage.Event {
	s.mu.RLock()
	defer s.mu.RUnlock()

	events := make([]storage.Event, 0)
	for _, e := range s.collection {
		if startSearch.After(e.StartTime) || endSearch.Before(e.EndTime) {
			continue
		}

		events = append(events, e)
	}

	return events
}
