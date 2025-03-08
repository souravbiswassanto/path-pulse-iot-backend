package handler

import (
	"context"
	e2 "errors"
	"fmt"
	"github.com/pkg/errors"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/tracker"
	"io"
	"log"
	"sync"
	"time"
)

func UpdateLocation(ctx context.Context, stream tracker.Tracker_UpdateLocationClient, data <-chan interface{}, updateInterval time.Duration) error {
	defer func() {
		endErr := stream.CloseSend()
		if endErr != nil {
			log.Fatalln(endErr)
		}
	}()
	return HandleClientSend(ctx, NewClientLocationStreamHandler(stream), data, updateInterval)
}

func HandlePulseRateStream(ctx context.Context, stream tracker.Tracker_UpdatePulseRateClient, data <-chan interface{}, updateInterval time.Duration) error {
	var streamSendErr, streamRecvErr chan error
	go HandlePulseRateUpdate(ctx, stream, data, streamSendErr, updateInterval)
	go HandlePulseRateAlert(ctx, stream, streamRecvErr)
	return WaitAndHandleSendRecvError(ctx, streamSendErr, streamRecvErr, models.PulseRateWithUserID{}, models.Alert{})
}

func WaitAndHandleSendRecvError(ctx context.Context, sendChan, recvChan chan error, sendItem, recvItem interface{}) error {
	var mu sync.Mutex
	var wg sync.WaitGroup
	var sendErr, recvErr error
	wg.Add(2)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			return
		case e := <-sendChan:
			mu.Lock()
			defer mu.Unlock()
			if e != nil {
				sendErr = e
			}
		}
	}()
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			return
		case e := <-recvChan:
			mu.Lock()
			defer mu.Unlock()
			if e != nil {
				recvErr = e
			}
		}
	}()
	wg.Wait()
	if sendErr != nil && recvErr != nil {
		return errors.Wrap(e2.Join(sendErr, recvErr), "failed on both send and receive")
	} else if sendErr != nil {
		return errors.Wrap(sendErr, fmt.Sprintf("failed on sending %v", sendItem))
	} else if recvErr != nil {
		return errors.Wrap(recvErr, fmt.Sprintf("failed on reciving %v", recvItem))
	}
	return nil
}

func HandlePulseRateUpdate(ctx context.Context, stream tracker.Tracker_UpdatePulseRateClient, data <-chan interface{}, errChan chan error, updateInterval time.Duration) {
	defer func() {
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

func HandlePulseRateAlert(ctx context.Context, stream tracker.Tracker_UpdatePulseRateClient, errChan chan error) {
	defer func() {
		close(errChan)
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
