package handler

import (
	"context"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/service"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/event"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/group"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
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
	return &user.Empty{}, esh.svc.AddEvent(ctx, eventProtoToModel(event))
}

func (esh *EventServerHandler) UpdateEvent(ctx context.Context, event *event.Event) (*user.Empty, error) {
	return &user.Empty{}, esh.svc.UpdateEvent(ctx, eventProtoToModel(event))
}

func (esh *EventServerHandler) DeleteEvent(ctx context.Context, eventId *event.EvenetId) (*user.Empty, error) {
	return &user.Empty{}, esh.svc.DeleteEvent(ctx, eventId.EId)
}

func (esh *EventServerHandler) GetSingleEventDetails(ctx context.Context, eventId *event.EvenetId) (*event.Event, error) {
	se, err := esh.svc.GetSingleEventDetails(ctx, eventId.EId)
	if err != nil {
		return nil, err
	}
	return eventModelToProto(se), nil
}

func (esh *EventServerHandler) ListEventsOfSingleUser(ctx context.Context, userId *user.UserID) (*event.EventList, error) {
	el, err := esh.svc.ListEventsOfSingleUser(ctx, userId.Id)
	if err != nil {
		return nil, err
	}
	pel := &event.EventList{}
	for _, e := range el {
		pel.EventList = append(pel.EventList, eventModelToProto(e))
	}
	return pel, nil
}

func (esh *EventServerHandler) ListEventsOfSingleGroup(ctx context.Context, groupId *group.GroupId) (*event.EventList, error) {
	el, err := esh.svc.ListEventsOfSingleGroup(ctx, groupId.GId)
	if err != nil {
		return nil, err
	}
	pel := &event.EventList{}
	for _, e := range el {
		pel.EventList = append(pel.EventList, eventModelToProto(e))
	}
	return pel, nil
}
