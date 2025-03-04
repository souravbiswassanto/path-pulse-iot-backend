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
	ctx      context.Context
	sh       StreamHandler
	workers  int
	wg       *sync.WaitGroup
	errChan  chan error
	jobs     chan interface{}
	userDone chan struct{}
	drop     bool
}

func NewWorkerPool(ctx context.Context, sh StreamHandler) *WorkerPool {
	cpu := runtime.NumCPU()
	workers := cpu * 2
	errChan := make(chan error, workers+1)
	jobs := make(chan interface{}, workers+1)
	userDone := make(chan struct{})
	return &WorkerPool{
		ctx:      ctx,
		sh:       sh,
		workers:  workers,
		wg:       &sync.WaitGroup{},
		jobs:     jobs,
		userDone: userDone,
		errChan:  errChan,
		drop:     true,
	}
}

func HandleClientStream(ctx context.Context, sh StreamHandler, applyFn ...func(wp *WorkerPool)) error {
	myCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	wp := NewWorkerPool(myCtx, sh)
	for _, fn := range applyFn {
		fn(wp)
	}
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.startWorker()
	}
	go wp.dropOnLoad()
	go wp.receiveClientRequest()
	wp.handleStop()
	return wp.checkForErrors()
}

// It will receive a function and will call it
func (wp *WorkerPool) startWorker() {
	defer wp.wg.Done()
	// fn func(job interface{}) error
	for {
		select {
		case <-wp.ctx.Done():
			return
		case pos, ok := <-wp.jobs:
			if len(wp.errChan) > 1 {
				return
			}
			if !ok {
				return
			}
			var err error
			val, err := wp.sh.Perform(pos)
			if err != nil {
				wp.errChan <- err
				return
			}
			if val == nil {
				continue
			}
			err = wp.sh.Send(val)
			if err != nil {
				wp.errChan <- err
				return
			}
		}
	}

}

func (wp *WorkerPool) receiveClientRequest() {
	defer close(wp.userDone)
	for {
		position, err := wp.sh.Receive()
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

func (wp *WorkerPool) dropOnLoad() {
	for {
		select {
		case <-wp.ctx.Done():
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

func (wp *WorkerPool) handleStop() {
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

// -----------------------------
