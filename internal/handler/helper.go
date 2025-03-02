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
	return positionProtoToModel(position), nil
}

func (uls *UpdateLocationServerHandler) Perform(job interface{}) (interface{}, error) {
	pos := job.(*models.Position)
	return nil, uls.svc.UpdateLocation(uls.ctx, pos)
}

func (uls *UpdateLocationServerHandler) Send(interface{}) error {
	return nil
}

// UpdatePulseRateServerHandler should be in the same package as tracker_handler.go
type UpdatePulseRateServerHandler struct {
	*TrackerHandler
	ctx    context.Context
	stream tracker.Tracker_UpdatePulseRateServer
}

func NewUpdatePulseRateServerHandler(ctx context.Context, th *TrackerHandler, stream tracker.Tracker_UpdatePulseRateServer) *UpdatePulseRateServerHandler {
	return &UpdatePulseRateServerHandler{
		ctx:            ctx,
		TrackerHandler: th,
		stream:         stream,
	}
}

func (uls *UpdatePulseRateServerHandler) Receive() (interface{}, error) {
	pr, err := uls.stream.Recv()
	if err != nil {
		return nil, err
	}
	if pr.UserId == 0 {
		return nil, fmt.Errorf("userId 0 not allowed")
	}
	return pulseRateWithUserIDProtoToModel(pr), nil
}

func (uls *UpdatePulseRateServerHandler) Perform(job interface{}) (interface{}, error) {
	pr := job.(*models.PulseRateWithUserID)
	return uls.svc.UpdatePulseRate(uls.ctx, pr)
}

func (uls *UpdatePulseRateServerHandler) Send(val interface{}) error {
	id := val.(*models.UserID)
	alert, err := uls.svc.PulseRateAlert(uls.ctx, id)
	if err != nil {
		return err
	}

	return uls.stream.Send(alertModelToProto(alert))
}

// UpdatePulseRateServerHandler should be in the same package as tracker_handler.go
type RealTimeDistanceServerHandler struct {
	*TrackerHandler
	ctx          context.Context
	stream       tracker.Tracker_GetRealTimeDistanceCoveredServer
	input        chan *models.Position
	outputStream <-chan float64
}

func NewRealTimeDistanceServerHandler(ctx context.Context, th *TrackerHandler, stream tracker.Tracker_GetRealTimeDistanceCoveredServer, input chan *models.Position, outputStream <-chan float64) *RealTimeDistanceServerHandler {
	return &RealTimeDistanceServerHandler{
		ctx:            ctx,
		TrackerHandler: th,
		stream:         stream,
		input:          input,
		outputStream:   outputStream,
	}
}

func (uls *RealTimeDistanceServerHandler) Receive() (interface{}, error) {
	pos, err := uls.stream.Recv()
	if err != nil {
		return nil, err
	}
	if pos.UserId == 0 {
		return nil, fmt.Errorf("userId 0 not allowed")
	}
	return positionProtoToModel(pos), nil
}

func (uls *RealTimeDistanceServerHandler) Perform(job interface{}) (interface{}, error) {
	pos := job.(*models.Position)
	uls.input <- pos
	return pos, nil
}

func (uls *RealTimeDistanceServerHandler) Send(_ interface{}) error {
	distance := <-uls.outputStream
	return uls.stream.Send(&tracker.Distance{
		Meter: distance,
	})
}
