package in_memory

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
)

type EventInMemoryStore struct {
	store *Store[uint64, *models.Event]
}

func NewEventInMemoryStore(s *Store[uint64, *models.Event]) *EventInMemoryStore {
	return &EventInMemoryStore{s}
}

func (es *EventInMemoryStore) Create(_ context.Context, v interface{}) error {
	e, ok := v.(*models.Event)
	if !ok {
		return fmt.Errorf("expected %T, got %T", &models.Event{}, v)
	}
	if e == nil || e.EventID == 0 {
		return fmt.Errorf("event can't be nil")
	}

	es.store.Create(e.EventID, e)
	return nil
}

func (es *EventInMemoryStore) Update(_ context.Context, v interface{}) error {
	e, ok := v.(*models.Event)
	if !ok {
		return fmt.Errorf("expected %T, got %T", &models.Event{}, v)
	}
	if e == nil || e.EventID == 0 {
		return fmt.Errorf("event can't be nil")
	}
	es.store.Update(e.EventID, e)

	return nil
}

func (es *EventInMemoryStore) Delete(_ context.Context, v interface{}) error {
	eventId, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("got %v, expected uint64 eventID", v)
	}
	if eventId == 0 {
		return fmt.Errorf("eventId can't be 0")
	}
	es.store.Delete(eventId)
	return nil
}

func (es *EventInMemoryStore) Get(_ context.Context, v interface{}) (interface{}, error) {
	eventId, ok := v.(uint64)
	if !ok {
		return nil, fmt.Errorf("got %v, expected uint64 eventID", v)
	}
	if eventId == 0 {
		return nil, fmt.Errorf("eventId can't be 0")
	}
	e, err := es.store.Get(eventId)
	if err != nil {
		return nil, err
	}
	return e, nil
}
