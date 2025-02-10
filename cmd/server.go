package cmd

import (
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/handler"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	"google.golang.org/grpc"
)

func NewGrpcServer() {
	svr := grpc.NewServer()
	user.RegisterUserManagerServer(svr, handler.NewUserHandler())

}
