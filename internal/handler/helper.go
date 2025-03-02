package handler

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/tracker"
)

// UpdateLocationServerHandler should be in the same package as tracker_handler.go
type UpdateLocationServerHandler struct {
	*TrackerHandler
	ctx    context.Context
	stream tracker.Tracker_UpdateLocationServer
}

func NewUpdateLocationStreamHandler(ctx context.Context, th *TrackerHandler, stream tracker.Tracker_UpdateLocationServer) *UpdateLocationServerHandler {
	return &UpdateLocationServerHandler{
		ctx:            ctx,
		TrackerHandler: th,
		stream:         stream,
	}
}

func (uls *UpdateLocationServerHandler) Receive() (interface{}, error) {
	position, err := uls.stream.Recv()
	if err != nil {
		return nil, err
	}
	if position.UserId == 0 {
		return nil, fmt.Errorf("userId 0 not allowed")
	}
	return trackerPositionProtoToModel(position), nil
}

func (uls *UpdateLocationServerHandler) Perform(job interface{}) (interface{}, error) {
	pos := job.(*models.Position)
	return nil, uls.svc.UpdateLocation(uls.ctx, pos)
}

func (uls *UpdateLocationServerHandler) Send(interface{}) error {
	return nil
}
