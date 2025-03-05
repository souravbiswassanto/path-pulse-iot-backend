package handler

import (
	"context"
	errors2 "errors"
	"fmt"
	"github.com/pkg/errors"
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

func UpdateLocation(ctx context.Context, stream tracker.Tracker_UpdateLocationClient, data <-chan interface{}, updateInterval time.Duration) error {
	defer func() {
		endErr := stream.CloseSend()
		if endErr != nil {
			log.Fatalln(endErr)
		}
	}()
	return HandleClientSend(ctx, NewClientLocationStreamHandler(stream), data, updateInterval)
}

func (tc *TrackerHandlerClient) HandlePulseRate(ctx context.Context, data <-chan interface{}, updateInterval time.Duration) error {
	stream, err := tc.cc.UpdatePulseRate(ctx)
	if err != nil {
		return err
	}
	return HandlePulseRateStream(ctx, stream, data, updateInterval)
}

// TODO: fix this function
func HandlePulseRateStream(ctx context.Context, stream tracker.Tracker_UpdatePulseRateClient, data <-chan interface{}, updateInterval time.Duration) error {
	var wg sync.WaitGroup

	wg.Add(2)
	var updateErr, alertErr chan error
	go HandlePulseRateUpdate(ctx, stream, data, &wg, updateErr, updateInterval)
	go HandlePulseRateAlert(ctx, stream, &wg, alertErr)
	err := HandleSendAndRecvError(updateErr, alertErr, models.PulseRateWithUserID{}, models.Alert{})
	wg.Wait()
	return err
}

func HandleSendAndRecvError(sendChan, recvChan chan error, sendItem, recvItem interface{}) error {
	var mu sync.Mutex
	var wg sync.WaitGroup
	var sendErr, recvErr error
	wg.Add(2)
	go func() {
		defer wg.Done()
		e := <-sendChan
		mu.Lock()
		defer mu.Unlock()
		if e != nil {
			sendErr = e
		}
	}()
	go func() {
		defer wg.Done()
		e := <-recvChan
		mu.Lock()
		defer mu.Unlock()
		if e != nil {
			recvErr = e
		}
	}()
	wg.Wait()
	if sendErr != nil && recvErr != nil {
		return errors.Wrap(errors2.Join(sendErr, recvErr), "failed on both send and receive")
	} else if sendErr != nil {
		return errors.Wrap(sendErr, fmt.Sprintf("failed on sending %v", sendItem))
	} else if recvErr != nil {
		return errors.Wrap(recvErr, fmt.Sprintf("failed on reciving %v", recvItem))
	}
	return nil
}

func HandlePulseRateUpdate(ctx context.Context, stream tracker.Tracker_UpdatePulseRateClient, data <-chan interface{}, wg *sync.WaitGroup, errChan chan error, updateInterval time.Duration) {
	defer func() {
		wg.Done()
		endErr := stream.CloseSend()
		if endErr != nil {
			log.Fatalln(endErr)
		}
		close(errChan)
	}()
	err := HandleClientSend(ctx, NewClientPulseRateStreamHandler(stream), data, updateInterval)
	if err != nil {
		errChan <- errors.Wrap(fmt.Errorf("failed to handle client send for pulse rate update"), err.Error())
	}
}

func HandlePulseRateAlert(ctx context.Context, stream tracker.Tracker_UpdatePulseRateClient, wg *sync.WaitGroup, errChan chan error) {
	defer func() {
		close(errChan)
		wg.Done()
	}()
	err := HandleClientReceive(ctx, NewClientPulseRateStreamHandler(stream))
	if err != nil {
		errChan <- errors.Wrap(fmt.Errorf("failed to handle client receive for alert"), err.Error())
	}
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
