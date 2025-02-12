package service

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db/in_memory"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
)

type EventService struct {
	eventDb in_memory.EventStore[uint64, *models.Event]
}

func NewEventService() *EventService {
	return &EventService{
		eventDb: in_memory.NewEventStore[uint64, *models.Event](),
	}
}

func (es *EventService) AddEvent(_ context.Context, event *models.Event) error {
	if event == nil || event.EventID == 0 {
		return fmt.Errorf("event can't be nil")
	}
	es.eventDb.Create(event.EventID, event)

	return nil
}

func (es *EventService) UpdateEvent(_ context.Context, event *models.Event) error {
	if event == nil || event.EventID == 0 {
		return fmt.Errorf("event can't be nil")
	}
	es.eventDb.Update(event.EventID, event)

	return nil
}

func (es *EventService) DeleteEvent(_ context.Context, eventId uint64) error {
	if eventId == 0 {
		return fmt.Errorf("eventId can't be 0")
	}
	es.eventDb.Delete(eventId)
	return nil
}

func (es *EventService) GetSingleEventDetails(ctx context.Context, eventId uint64) (*models.Event, error) {
	user, err := es.eventDb.Get(eventId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (es *EventService) ListEventsOfSingleUser(_ context.Context, userId uint64) ([]*models.Event, error) {
	events := make([]*models.Event, 0)
	eventLists := es.eventDb.List()
	for _, event := range eventLists {
		if event.PublisherID != nil && uint64(*event.PublisherID) == userId {
			events = append(events, event)
		}
	}
	return events, nil
}

func (es *EventService) ListEventsOfSingleGroup(_ context.Context, groupId uint64) ([]*models.Event, error) {
	events := make([]*models.Event, 0)
	eventLists := es.eventDb.List()
	for _, event := range eventLists {
		if event.GroupID == groupId {
			events = append(events, event)
		}
	}
	return events, nil
}
