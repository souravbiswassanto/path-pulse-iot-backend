package handler

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db/influx"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/service"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/tracker"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	"google.golang.org/grpc"
	"io"
	"log"
	"sync"
	"time"
)

type TrackerHandlerServer struct {
	svc *service.TrackerService
	tracker.UnimplementedTrackerServer
}

func NewTrackerHandlerServer(options *influx.InfluxDBOptions) *TrackerHandlerServer {
	return &TrackerHandlerServer{
		svc: service.NewTrackerService(options),
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
	lp LocationProvider
	pr PulseRateProvider
}

func NewTrackerHandlerClient(cc grpc.ClientConnInterface, lp LocationProvider, pr PulseRateProvider) *TrackerHandlerClient {
	return &TrackerHandlerClient{
		cc: tracker.NewTrackerClient(cc),
		lp: lp,
		pr: pr,
	}
}

func (tc *TrackerHandlerClient) HandleLocationUpdate(ctx context.Context, data <-chan interface{}, updateInterval time.Duration) error {
	stream, err := tc.cc.UpdateLocation(ctx)
	if err != nil {
		return err
	}
	return tc.UpdateLocation(ctx, stream, data, updateInterval)
}

func (tc *TrackerHandlerClient) UpdateLocation(ctx context.Context, stream tracker.Tracker_UpdateLocationClient, data <-chan interface{}, updateInterval time.Duration) error {
	defer func() {
		endErr := stream.CloseSend()
		if endErr != nil {
			log.Fatalln(endErr)
		}
	}()
	return HandleClientSend(ctx, NewClientLocationStreamHandler(stream), data, updateInterval)
}

func (tc *TrackerHandlerClient) CurrentLocation() (*models.Position, error) {
	l, err := tc.lp.GetCurrentLocation()
	if err != nil {
		return nil, err
	}
	return &models.Position{
		Latitude:  l.Latitude(),
		Longitude: l.Longitude(),
	}, nil
}

func (tc *TrackerHandlerClient) HandlePulseRate(ctx context.Context, data <-chan interface{}, updateInterval time.Duration) error {
	stream, err := tc.cc.UpdatePulseRate(ctx)
	if err != nil {
		return err
	}

	return tc.HandlePulseRateStream(ctx, stream, data, updateInterval)
}

// TODO: fix this function
func (tc *TrackerHandlerClient) HandlePulseRateStream(ctx context.Context, stream tracker.Tracker_UpdatePulseRateClient, data <-chan interface{}, updateInterval time.Duration) error {
	var wg sync.WaitGroup
	wg.Add(2)
	go tc.HandlePulseRateUpdate(ctx, stream, data, &wg, updateInterval)
	go tc.HandlePulseRateAlert(ctx, stream, &wg)
	wg.Wait()
	return nil
}

func (tc *TrackerHandlerClient) HandlePulseRateUpdate(ctx context.Context, stream tracker.Tracker_UpdatePulseRateClient, data <-chan interface{}, wg *sync.WaitGroup, updateInterval time.Duration) error {
	defer func() {
		wg.Done()
		endErr := stream.CloseSend()
		if endErr != nil {
			log.Fatalln(endErr)
		}
	}()
	return HandleClientSend(ctx, NewClientPulseRateStreamHandler(stream), data, updateInterval)
}

func (tc *TrackerHandlerClient) HandlePulseRateAlert(ctx context.Context, stream tracker.Tracker_UpdatePulseRateClient, wg *sync.WaitGroup) error {
	defer wg.Done()
	return HandleClientReceive(ctx, NewClientPulseRateStreamHandler(stream))
}

func (tc *TrackerHandlerClient) UpdatePulseRate() {

}

func (tc *TrackerHandlerClient) CurrentPulseRate() (float32, error) {
	pr, err := tc.pr.GetCurrentPulseRate()
	if err != nil {
		return 0.0, err
	}
	return pr.Pulse(), nil
}

func HandleClientSend(ctx context.Context, st StreamHandler, data <-chan interface{}, updateInterval time.Duration) error {

	ticker := time.NewTicker(updateInterval)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			pos, ok := <-data
			if !ok {
				return nil
			}
			err := st.Send(pos)
			if err != nil {
				return err
			}
		}
	}
}

func HandleClientReceive(ctx context.Context, st StreamHandler) error {
	ticker := time.NewTicker(time.Microsecond * 5)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			val, err := st.Receive()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			_, err = st.Perform(val)
			if err != nil {
				return err
			}
		}
	}
}
