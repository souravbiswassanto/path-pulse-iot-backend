package handler

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db/influx"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/service"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/tracker"
	user "github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
)

type TrackerHandler struct {
	svc *service.TrackerService
	tracker.UnimplementedTrackerServer
}

func NewTrackerHandler(options *influx.InfluxDBOptions) *TrackerHandler {
	return &TrackerHandler{
		svc: service.NewTrackerService(options),
	}
}

func (th *TrackerHandler) GetLocation(ctx context.Context, userID *user.UserID) (*tracker.Position, error) {
	if userID == nil || userID.Id == 0 {
		return nil, fmt.Errorf("userId not given")
	}
	position, err := th.svc.GetLocation(ctx, (*models.UserID)(&userID.Id))
	if err != nil {
		return nil, err
	}
	return trackerPositionModelToProto(position), nil
}

func (th *TrackerHandler) UpdateLocation(stream tracker.Tracker_UpdateLocationServer) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wp := NewWorkerPool(th)
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.startUpdateLocationWorker(ctx, func(job interface{}) error {
			pos := job.(*models.Position)
			return wp.svc.UpdateLocation(ctx, pos)
		})
	}
	go wp.dropPositionFromQueueOnLoad(ctx)
	go wp.handleClientPositionUpdate(func() (interface{}, error) {
		position, err := stream.Recv()
		if err != nil {
			return nil, err
		}
		if position.UserId == 0 {
			return nil, fmt.Errorf("userId 0 not allowed")
		}
		return trackerPositionProtoToModel(position), nil
	})
	wp.HandleStop()
	return wp.checkForErrors()
}

func (th *TrackerHandler) Checkpoint(ctx context.Context, position *tracker.Position) (*tracker.CheckpointID, error) {
	ckId, err := th.svc.Checkpoint(ctx, trackerPositionProtoToModel(position))
	if err != nil {
		return nil, err
	}
	return &tracker.CheckpointID{CkId: ckId}, nil
}
