package handler

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/service"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/event"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/group"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	"google.golang.org/grpc"
	"log"
)

type EventServerHandler struct {
	svc *service.EventService
	event.UnimplementedEventManagerServer
}

func NewEventServerHandler() *EventServerHandler {
	return &EventServerHandler{
		// TODO: need to fix
		svc: service.NewEventService(nil, nil),
	}
}

func (esh *EventServerHandler) AddEvent(ctx context.Context, event *event.Event) (*user.Empty, error) {
	return &user.Empty{}, esh.svc.AddEvent(ctx, eventProtoToModel(event))
}

func (esh *EventServerHandler) UpdateEvent(ctx context.Context, event *event.Event) (*user.Empty, error) {
	return &user.Empty{}, esh.svc.UpdateEvent(ctx, eventProtoToModel(event))
}

func (esh *EventServerHandler) DeleteEvent(ctx context.Context, eventId *event.EventId) (*user.Empty, error) {
	return &user.Empty{}, esh.svc.DeleteEvent(ctx, eventId.EId)
}

func (esh *EventServerHandler) GetSingleEventDetails(ctx context.Context, eventId *event.EventId) (*event.Event, error) {
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

type EventClientHandler struct {
	cc event.EventManagerClient
}

func NewEventManagerClientHandler(cc grpc.ClientConnInterface) *EventClientHandler {
	return &EventClientHandler{
		cc: event.NewEventManagerClient(cc),
	}
}

func (ech *EventClientHandler) CreateEvent(event *event.Event) error {
	if event == nil || event.EId == 0 {
		return fmt.Errorf("cant't create event, event or eventID is nil")
	}
	_, err := ech.cc.AddEvent(context.TODO(), event)
	if err != nil {
		return err
	}
	log.Println("successfully created event with EventID: ", event.EId)
	return err
}

func (ech *EventClientHandler) GetEvent(eventID uint64) (*event.Event, error) {

	event, err := ech.cc.GetSingleEventDetails(context.TODO(), &event.EventId{EId: eventID})
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (ech *EventClientHandler) UpdateEvent(event *event.Event) error {
	if event == nil {
		return fmt.Errorf("updating event can't be nil")
	}
	_, err := ech.cc.UpdateEvent(context.TODO(), event)
	if err != nil {
		return err
	}
	log.Println("successfully updated event")
	return nil
}

func (ech *EventClientHandler) DeleteEvent(eventEId uint64) error {
	_, err := ech.cc.DeleteEvent(context.TODO(), &event.EventId{
		EId: eventEId,
	})
	return err
}

func (ech *EventClientHandler) ListEventsOfSingleUser(ctx context.Context, uid *models.UserID) error {

	return nil
}
func (ech *EventClientHandler) ListEventsOfSingleGroup(ctx context.Context, gid uint64) error {

	return nil
}
