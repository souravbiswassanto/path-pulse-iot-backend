package handler

import (
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"runtime"
	"sync"
)

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
