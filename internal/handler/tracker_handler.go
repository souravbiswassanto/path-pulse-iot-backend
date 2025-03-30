package handler

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/service"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/tracker"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	"google.golang.org/grpc"
	"time"
)

type TrackerHandlerServer struct {
	svc *service.TrackerService
	tracker.UnimplementedTrackerServer
}

func NewTrackerServerHandler(svc *service.TrackerService) *TrackerHandlerServer {
	return &TrackerHandlerServer{
		svc: svc,
	}
}

func (th *TrackerHandlerServer) GetLocation(ctx context.Context, userID *user.UserID) (*tracker.Position, error) {
	if userID == nil || userID.Id == 0 {
		return nil, fmt.Errorf("userId not given")
	}
	position, err := th.svc.GetLocation(ctx, (*models.UserID)(&userID.Id))
	if err != nil {
		return nil, err
	}
	return positionModelToProto(position), nil
}

func (th *TrackerHandlerServer) UpdateLocation(stream tracker.Tracker_UpdateLocationServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	b := func(wp *WorkerPool) {
		wp.drop = true
	}
	return HandleClientStream(ctx, NewUpdateLocationStreamHandler(ctx, th, stream), b)
}

func (th *TrackerHandlerServer) Checkpoint(ctx context.Context, position *tracker.Position) (*tracker.CheckpointID, error) {
	ckId, err := th.svc.Checkpoint(ctx, positionProtoToModel(position))
	if err != nil {
		return nil, err
	}
	return &tracker.CheckpointID{CkId: ckId}, nil
}

func (th *TrackerHandlerServer) UpdatePulseRate(stream tracker.Tracker_UpdatePulseRateServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	b := func(wp *WorkerPool) {
		wp.drop = true
	}
	return HandleClientStream(ctx, NewUpdatePulseRateServerHandler(ctx, th, stream), b)
}

func (th *TrackerHandlerServer) GetRealTimeDistanceCovered(stream tracker.Tracker_GetRealTimeDistanceCoveredServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	inputStream := make(chan *models.Position)
	outputStream := th.svc.GetRealTimeDistanceCovered(ctx, inputStream)
	return HandleClientStream(ctx, NewRealTimeDistanceServerHandler(ctx, th, stream, inputStream, outputStream))
}

func (th *TrackerHandlerServer) GetTotalDistanceBetweenCheckpoint(ctx context.Context, ctf *tracker.CheckpointToAndFrom) (*tracker.Distance, error) {
	distance, err := th.svc.GetTotalDistanceBetweenCheckpoint(ctx, checkpointToAndFromProtoToModel(ctf))
	if err != nil {
		return nil, err
	}
	return &tracker.Distance{Meter: distance}, nil
}

// ----------- Client ----------------

type TrackerHandlerClient struct {
	cc tracker.TrackerClient
}

func NewTrackerHandlerClient(cc grpc.ClientConnInterface) *TrackerHandlerClient {
	return &TrackerHandlerClient{
		cc: tracker.NewTrackerClient(cc),
	}
}

func (tc *TrackerHandlerClient) HandleLocationUpdate(ctx context.Context, data <-chan interface{}, updateInterval time.Duration) error {
	stream, err := tc.cc.UpdateLocation(ctx)
	if err != nil {
		return err
	}
	return UpdateLocation(ctx, stream, data, updateInterval)
}

func (tc *TrackerHandlerClient) HandlePulseRate(ctx context.Context, data <-chan interface{}, updateInterval time.Duration) error {
	stream, err := tc.cc.UpdatePulseRate(ctx)
	if err != nil {
		return err
	}
	return HandlePulseRateStream(ctx, stream, data, updateInterval)
}
