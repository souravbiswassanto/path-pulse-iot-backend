package service

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
)

type EventService struct {
	eventDb db.DB
	cache   db.DB
}

func NewEventService(db db.DB, cache db.DB) *EventService {
	return &EventService{
		eventDb: db,
		cache:   cache,
	}
}

func (es *EventService) AddEvent(ctx context.Context, event *models.Event) error {
	if event == nil || event.EventID == 0 {
		return fmt.Errorf("event can't be nil")
	}
	return es.eventDb.Create(ctx, event)
}

func (es *EventService) UpdateEvent(ctx context.Context, event *models.Event) error {
	if event == nil || event.EventID == 0 {
		return fmt.Errorf("event can't be nil")
	}
	err := es.eventDb.Update(ctx, event)
	if err != nil {
		return err
	}
	_ = es.cache.Update(ctx, event)
	return nil
}

func (es *EventService) DeleteEvent(ctx context.Context, eventId uint64) error {
	if eventId == 0 {
		return fmt.Errorf("eventId can't be 0")
	}
	err := es.eventDb.Delete(ctx, eventId)
	if err != nil {
		return err
	}
	_ = es.cache.Delete(ctx, eventId)
	return nil
}

func (es *EventService) GetSingleEventDetails(ctx context.Context, eventId uint64) (*models.Event, error) {
	v, err := es.cache.Get(ctx, eventId)
	if err == nil {
		event, ok := v.(*models.Event)
		if ok {
			return event, nil
		}
		// need a log here
	}
	k, err := es.eventDb.Get(ctx, eventId)
	if err != nil {
		return nil, err
	}
	event, ok := k.(*models.Event)
	if !ok {
		return nil, fmt.Errorf("can't decode event, expected %T, got %T", &models.Event{}, k)
	}
	_ = es.cache.Create(ctx, event)
	return event, nil
}

//TODO: implement later

func (es *EventService) ListEventsOfSingleUser(_ context.Context, userId uint64) ([]*models.Event, error) {
	//events := make([]*models.Event, 0)
	//eventLists := es.eventDb.List()
	//for _, event := range eventLists {
	//	if event.PublisherID != nil && uint64(*event.PublisherID) == userId {
	//		events = append(events, event)
	//	}
	//}
	return nil, fmt.Errorf("not implemented")
}

// TODO: Implement later

func (es *EventService) ListEventsOfSingleGroup(_ context.Context, groupId uint64) ([]*models.Event, error) {
	//events := make([]*models.Event, 0)
	//eventLists := es.eventDb.List()
	//for _, event := range eventLists {
	//	if event.GroupID == groupId {
	//		events = append(events, event)
	//	}
	//}
	return nil, fmt.Errorf("not implemented")
}
