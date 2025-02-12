package handler

import (
	"context"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/service"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/event"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/group"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	"google.golang.org/genproto/googleapis/type/datetime"
	"time"
)

type EventServerHandler struct {
	svc *service.EventService
	event.UnimplementedEventManagerServer
}

func NewEventServerHandler() *EventServerHandler {
	return &EventServerHandler{
		svc: service.NewEventService(),
	}
}

func (esh *EventServerHandler) AddEvent(ctx context.Context, event *event.Event) (*user.Empty, error) {
	return &user.Empty{}, nil
}

func (esh *EventServerHandler) UpdateEvent(ctx context.Context, event *event.Event) (*user.Empty, error) {
	return &user.Empty{}, nil
}

func (esh *EventServerHandler) DeleteEvent(ctx context.Context, eventId *event.EvenetId) (*user.Empty, error) {
	return &user.Empty{}, nil
}

func (esh *EventServerHandler) GetSingleEventDetails(ctx context.Context, eventId *event.EvenetId) (*event.Event, error) {
	return &event.Event{}, nil
}

func (esh *EventServerHandler) ListEventsOfSingleUser(ctx context.Context, userId *user.UserID) (*event.EventList, error) {
	return &event.EventList{}, nil
}

func (esh *EventServerHandler) ListEventsOfSingleGroup(ctx context.Context, groupId *group.GroupId) (*event.EventList, error) {
	return &event.EventList{}, nil
}

func eventProtoToModel(event *event.Event) *models.Event {
	return &models.Event{
		EventID:     event.EId,
		GroupID:     event.GId,
		PublisherID: (*models.UserID)(&event.Publisher),
		State:       eventStateProtoToModel(event.State),
		Interested: func() []*models.UserID {
			if event.Interested == nil {
				return nil
			}
			var ids []*models.UserID
			for _, id := range event.Interested {
				ids = append(ids, (*models.UserID)(&id.Id))
			}
			return ids
		}(),
		Going: func() []*models.UserID {
			if event.Going == nil {
				return nil
			}
			var ids []*models.UserID
			for _, id := range event.Going {
				ids = append(ids, (*models.UserID)(&id.Id))
			}
			return ids
		}(),
		NotInterested: func() []*models.UserID {
			if event.NotInterested == nil {
				return nil
			}
			var ids []*models.UserID
			for _, id := range event.NotInterested {
				ids = append(ids, (*models.UserID)(&id.Id))
			}
			return ids
		}(),
		EventDesc: models.EventDescription{
			Description: event.EventDesc.GetDesc(),
			Name:        event.EventDesc.GetName(),
		},
		EventDateTime: func() *time.Time {
			t := ProtoDateTimeToTime(event.EventDateTime)
			return &t
		}(),
	}
}

func eventModelToProto(e *models.Event) *event.Event {
	return &event.Event{
		EId:       e.EventID,
		GId:       e.GroupID,
		Publisher: (uint64)(*e.PublisherID),
		State:     eventStateModelToProto(e.State),
		Interested: func() []*user.UserID {
			if e.Interested == nil {
				return nil
			}
			var ids []*user.UserID
			for _, id := range e.Interested {
				ids = append(ids, &user.UserID{Id: uint64(*id)})
			}
			return ids
		}(),
		Going: func() []*user.UserID {
			if e.Going == nil {
				return nil
			}
			var ids []*user.UserID
			for _, id := range e.Going {
				ids = append(ids, &user.UserID{Id: uint64(*id)})
			}
			return ids
		}(),
		NotInterested: func() []*user.UserID {
			if e.NotInterested == nil {
				return nil
			}
			var ids []*user.UserID
			for _, id := range e.NotInterested {
				ids = append(ids, &user.UserID{Id: uint64(*id)})
			}
			return ids
		}(),
		EventDesc: &event.EventDescription{
			Desc: e.EventDesc.Description,
			Name: e.EventDesc.Name,
		},
		EventDateTime: TimeToProtoDateTime(*e.EventDateTime),
	}
}

func ProtoDateTimeToTime(dt *datetime.DateTime) time.Time {
	if dt == nil {
		return time.Time{}
	}
	nanos := time.Duration(dt.Nanos) * time.Nanosecond

	return time.Date(
		int(dt.Year),
		time.Month(dt.Month),
		int(dt.Day),
		int(dt.Hours),
		int(dt.Minutes),
		int(dt.Seconds),
		int(nanos),
		time.UTC,
	)
}

func TimeToProtoDateTime(t time.Time) *datetime.DateTime {
	return &datetime.DateTime{
		Year:    int32(t.Year()),
		Month:   int32(t.Month()),
		Day:     int32(t.Day()),
		Hours:   int32(t.Hour()),
		Minutes: int32(t.Minute()),
		Seconds: int32(t.Second()),
		Nanos:   int32(t.Nanosecond()),
	}
}

func eventStateProtoToModel(state event.EventState) models.EventState {
	switch state {
	case 0:
		return models.EventOngoing
	case 1:
		return models.EventClosed
	case 2:
		return models.EventUpcoming
	default:
		return models.EventUnknown
	}
}

func eventStateModelToProto(e models.EventState) event.EventState {
	switch e {
	case models.EventOngoing:
		return event.EventState_ongoing
	case models.EventClosed:
		return event.EventState_closed
	case models.EventUpcoming:
		return event.EventState_upcoming
	}
	return event.EventState_unknown
}
