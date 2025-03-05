package handler

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/tracker"
	"log"
)

// UpdateLocationServerHandler should be in the same package as tracker_handler.go
type UpdateLocationServerHandler struct {
	*TrackerHandlerServer
	ctx    context.Context
	stream tracker.Tracker_UpdateLocationServer
}

func NewUpdateLocationStreamHandler(ctx context.Context, th *TrackerHandlerServer, stream tracker.Tracker_UpdateLocationServer) *UpdateLocationServerHandler {
	return &UpdateLocationServerHandler{
		ctx:                  ctx,
		TrackerHandlerServer: th,
		stream:               stream,
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
	*TrackerHandlerServer
	ctx    context.Context
	stream tracker.Tracker_UpdatePulseRateServer
}

func NewUpdatePulseRateServerHandler(ctx context.Context, th *TrackerHandlerServer, stream tracker.Tracker_UpdatePulseRateServer) *UpdatePulseRateServerHandler {
	return &UpdatePulseRateServerHandler{
		ctx:                  ctx,
		TrackerHandlerServer: th,
		stream:               stream,
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
	*TrackerHandlerServer
	ctx          context.Context
	stream       tracker.Tracker_GetRealTimeDistanceCoveredServer
	input        chan *models.Position
	outputStream <-chan float64
}

func NewRealTimeDistanceServerHandler(ctx context.Context, th *TrackerHandlerServer, stream tracker.Tracker_GetRealTimeDistanceCoveredServer, input chan *models.Position, outputStream <-chan float64) *RealTimeDistanceServerHandler {
	return &RealTimeDistanceServerHandler{
		ctx:                  ctx,
		TrackerHandlerServer: th,
		stream:               stream,
		input:                input,
		outputStream:         outputStream,
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

type currentLocation struct {
	latitude, longitude float64
}

func (cl *currentLocation) GetCurrentLocation() (interface{}, error) {
	// will implement location getting function here

	return &currentLocation{
		latitude:  0.0,
		longitude: 0.0,
	}, nil
}

func (cl *currentLocation) Latitude() float64 {
	return cl.latitude
}
func (cl *currentLocation) Longitude() float64 {
	return cl.longitude
}

// ------ clientLocationStreamHandler

type clientLocationStreamHandler struct {
	stream tracker.Tracker_UpdateLocationClient
}

func NewClientLocationStreamHandler(stream tracker.Tracker_UpdateLocationClient) *clientLocationStreamHandler {
	return &clientLocationStreamHandler{stream: stream}
}

func (cs *clientLocationStreamHandler) Send(val interface{}) error {
	pos := val.(*models.Position)
	if pos.Longitude == 0.0 || pos.Latitude == 0.0 {
		return fmt.Errorf("longitude latitude can't be 0.0")
	}
	return cs.stream.Send(positionModelToProto(pos))
}

func (cs *clientLocationStreamHandler) Receive() (interface{}, error) {
	return nil, nil
}

func (cs *clientLocationStreamHandler) Perform(interface{}) (interface{}, error) {
	return nil, nil
}

// ------ clientLocationStreamHandler

type clientPulseRateStreamHandler struct {
	stream tracker.Tracker_UpdatePulseRateClient
}

func NewClientPulseRateStreamHandler(stream tracker.Tracker_UpdatePulseRateClient) *clientPulseRateStreamHandler {
	return &clientPulseRateStreamHandler{stream: stream}
}

func (cs *clientPulseRateStreamHandler) Send(val interface{}) error {
	pr := val.(*models.PulseRateWithUserID)
	return cs.stream.Send(pulseRateWithUserIDModelToProto(pr))
}

func (cs *clientPulseRateStreamHandler) Receive() (interface{}, error) {
	alert, err := cs.stream.Recv()
	if err != nil {
		return nil, err
	}
	return alertProtoToModel(alert), nil
}

func (cs *clientPulseRateStreamHandler) Perform(obj interface{}) (interface{}, error) {
	alert := obj.(*models.Alert)
	log.Println(alert)
	return nil, nil
}

// ------ clientDistanceStreamHandler

type clientDistanceStreamHandler struct {
	tc     *TrackerHandlerClient
	stream tracker.Tracker_GetRealTimeDistanceCoveredClient
}

func (cs *clientDistanceStreamHandler) Send(val interface{}) error {
	pos := val.(*models.Position)
	return cs.stream.Send(positionModelToProto(pos))
}

func (cs *clientDistanceStreamHandler) Receive() (interface{}, error) {
	distance, err := cs.stream.Recv()
	if err != nil {
		return nil, err
	}
	return distance.Meter, nil
}

func (cs *clientDistanceStreamHandler) Perform(obj interface{}) (interface{}, error) {
	distance := obj.(float64)
	log.Println("distance covered so far:", distance)
	return nil, nil
}
