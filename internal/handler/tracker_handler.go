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

	cpu := runtime.NumCPU()
	workers := cpu * 2

	var wg sync.WaitGroup
	errChan := make(chan error, workers+1)
	jobs := make(chan *models.Position, workers+1)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go th.startUpdateLocationWorker(ctx, jobs, errChan, &wg)
	}
	go dropPositionFromQueueOnLoad(ctx, jobs, workers)

	for {
		position, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			errChan <- err
			break
		}
		if position.UserId == 0 {
			errChan <- fmt.Errorf("userId 0 not allowed")
			break
		}
		if len(errChan) > 0 {
			break
		}
		jobs <- trackerPositionProtoToModel(position)
	}

	close(jobs)
	wg.Wait()
	close(errChan)
	if len(errChan) > 0 {
		var err []error
		for e := range errChan {
			err = append(err, e)
		}
		return errors.Join(err...)
	}
	return nil
}

func (th *TrackerHandler) startUpdateLocationWorker(ctx context.Context, job <-chan *models.Position, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case pos, ok := <-job:
			if len(errChan) > 1 {
				return
			}
			if !ok {
				return
			}
			err := th.svc.UpdateLocation(ctx, pos)
			if err != nil {
				errChan <- err
			}
		}
	}

}
func dropPositionFromQueueOnLoad(ctx context.Context, job <-chan *models.Position, worker int) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Millisecond * 50):
			if len(job) > worker {
				// drop a job
				select {
				case <-job:
				default:
				}
			}
		}
	}
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
