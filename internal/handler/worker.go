package handler

import (
	"context"
	"errors"
	"io"
	"runtime"
	"sync"
	"time"
)

type WorkerWithResult func(interface{}) (interface{}, error)
type WorkerWithoutResult func(interface{}) error

type WorkerPool struct {
	*TrackerHandler
	workers  int
	wg       *sync.WaitGroup
	errChan  chan error
	jobs     chan interface{}
	userDone chan struct{}
}

func NewWorkerPool(th *TrackerHandler) *WorkerPool {
	cpu := runtime.NumCPU()
	workers := cpu * 2
	errChan := make(chan error, workers+1)
	jobs := make(chan interface{}, workers+1)
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

// It will receive a function and will call it
func (wp *WorkerPool) startUpdateLocationWorker(ctx context.Context, fn interface{}) {
	defer wp.wg.Done()
	// fn func(job interface{}) error
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
			var err error
			switch fn.(type) {
			case WorkerWithoutResult:
				err = fn.(WorkerWithoutResult)(pos)
			case WorkerWithResult:
				_, err = fn.(WorkerWithResult)(pos)
			}
			if err != nil {
				wp.errChan <- err
			}
		}
	}

}

func (wp *WorkerPool) handleClientPositionUpdate(fn func() (interface{}, error)) {
	defer close(wp.userDone)
	for {
		position, err := fn()
		if err == io.EOF {
			break
		}
		if err != nil {
			wp.errChan <- err
			break
		}

		if len(wp.errChan) > 0 {
			break
		}
		wp.jobs <- position
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
