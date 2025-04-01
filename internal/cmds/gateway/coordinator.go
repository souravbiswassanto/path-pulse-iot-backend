package gateway

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/event"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/group"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/tracker"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type Coordinator struct {
	*Config
	ctx context.Context
	log logr.Logger
}

func (c *Coordinator) SetupGateway() error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := user.RegisterUserManagerHandlerFromEndpoint(c.ctx, mux, c.GrpcAddr, opts)
	if err != nil {
		return err
	}
	err = event.RegisterEventManagerHandlerFromEndpoint(c.ctx, mux, c.GrpcAddr, opts)
	if err != nil {
		return err
	}
	err = group.RegisterGroupManagerHandlerFromEndpoint(c.ctx, mux, c.GrpcAddr, opts)
	if err != nil {
		return err
	}
	err = tracker.RegisterTrackerHandlerFromEndpoint(c.ctx, mux, c.GrpcAddr, opts)
	if err != nil {
		return err
	}
	if err := http.ListenAndServe(c.Addr, mux); err != nil {
		c.log.Error(err, "failed to listen and serve on addr %v", c.Addr)
	}

	return nil
}
