package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db/influx"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/service"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/tracker"
	user "github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	"io"
	"runtime"
	"sync"
	"time"
)

type TrackerHandler struct {
	svc *service.TrackerService
	tracker.UnimplementedTrackerServer
}

type WorkerPool struct {
	*TrackerHandler
	workers  int
	wg       *sync.WaitGroup
	errChan  chan error
	jobs     chan *models.Position
	userDone chan struct{}
}

func NewWorkerPool(th *TrackerHandler) *WorkerPool {
	cpu := runtime.NumCPU()
	workers := cpu * 2
	errChan := make(chan error, workers+1)
	jobs := make(chan *models.Position, workers+1)
	userDone := make(chan struct{})
	return &WorkerPool{
		TrackerHandler: th,
		workers:        workers,
		wg:             &sync.WaitGroup{},
		jobs:           jobs,
		userDone:       userDone,
		errChan:        errChan,
	}
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
		go wp.startUpdateLocationWorker(ctx)
	}
	go wp.dropPositionFromQueueOnLoad(ctx)
	go wp.handleClientPositionUpdate(stream)
	wp.HandleStop()
	return wp.checkForErrors()
}

func (wp *WorkerPool) startUpdateLocationWorker(ctx context.Context) {
	defer wp.wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case pos, ok := <-wp.jobs:
			if len(wp.errChan) > 1 {
				return
			}
			if !ok {
				return
			}
			err := wp.svc.UpdateLocation(ctx, pos)
			if err != nil {
				wp.errChan <- err
			}
		}
	}

}

func (wp *WorkerPool) handleClientPositionUpdate(stream tracker.Tracker_UpdateLocationServer) {
	defer close(wp.userDone)
	for {
		position, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			wp.errChan <- err
			break
		}
		if position.UserId == 0 {
			wp.errChan <- fmt.Errorf("userId 0 not allowed")
			break
		}
		if len(wp.errChan) > 0 {
			break
		}
		wp.jobs <- trackerPositionProtoToModel(position)
	}
	return
}

func (wp *WorkerPool) dropPositionFromQueueOnLoad(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Millisecond * 50):
			if len(wp.jobs) >= wp.workers {
				// drop a job
				select {
				case <-wp.jobs:
				default:
				}
			}
		}
	}
}

func (wp *WorkerPool) HandleStop() {
	<-wp.userDone
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.errChan)
}

func (wp *WorkerPool) checkForErrors() error {
	if len(wp.errChan) > 0 {
		var err []error
		for e := range wp.errChan {
			err = append(err, e)
		}
		return errors.Join(err...)
	}
	return nil
}

func trackerPositionModelToProto(pos *models.Position) *tracker.Position {
	return &tracker.Position{
		Longitude: pos.Longitude,
		Latitude:  pos.Latitude,
		UserId:    uint64(pos.UID),
		Time:      TimeToProtoDateTime(pos.Time),
	}
}
func trackerPositionProtoToModel(pos *tracker.Position) *models.Position {
	return &models.Position{
		Longitude:    pos.Longitude,
		Latitude:     pos.Latitude,
		CheckPointID: pos.CkId,
		Time:         ProtoDateTimeToTime(pos.Time),
	}
}
