package handler

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/service"
	proto "github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	"google.golang.org/grpc"
	"log"
)

type UserServerHandlerSer struct {
	svc *service.UserService
	proto.UnimplementedUserManagerServer
}

func NewUserServerHandler(svc *service.UserService) *UserServerHandlerSer {
	////ctx, cancel := context.WithCancel(context.TODO())
	////defer cancel()
	//// TODO: need to redesign here
	//sqlDb, err := postgres.NewUseSqlDB(nil)
	//if err != nil {
	//	return nil, err
	//}
	//cacheDb := in_memory.NewEventInMemoryStore(nil)
	//svc := service.NewUserService(sqlDb, cacheDb)
	return &UserServerHandlerSer{
		svc: svc,
	}
}

func (ush *UserServerHandlerSer) GetUser(ctx context.Context, uid *proto.UserID) (*proto.User, error) {

	if uid == nil {
		return nil, fmt.Errorf("given user id can't be nil")
	}
	user, err := ush.svc.GetUser(ctx, (models.UserID)(uid.Id))
	if err != nil {
		return nil, err
	}
	return userModelToProto(user), nil
}

func (ush *UserServerHandlerSer) CreateUser(ctx context.Context, user *proto.User) (*proto.Empty, error) {
	if user == nil {
		return &proto.Empty{}, fmt.Errorf("given user is nil")
	}
	err := ush.svc.CreateUser(ctx, userProtoToModel(user))
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

func (ush *UserServerHandlerSer) UpdateUser(ctx context.Context, user *proto.User) (*proto.Empty, error) {
	if user == nil {
		return nil, fmt.Errorf("upating user can't be nil")
	}
	err := ush.svc.UpdateUser(ctx, userProtoToModel(user))
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

func (ush *UserServerHandlerSer) DeleteUser(ctx context.Context, uid *proto.UserID) (*proto.Empty, error) {
	if uid == nil {
		return nil, fmt.Errorf("the request user id can't be nil")
	}
	err := ush.svc.DeleteUser(ctx, (models.UserID)(uid.Id))
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

type UserClientHandler struct {
	cc proto.UserManagerClient
}

func NewUserManagerClientHandler(cc grpc.ClientConnInterface) *UserClientHandler {
	return &UserClientHandler{
		cc: proto.NewUserManagerClient(cc),
	}
}

func (uch *UserClientHandler) CreateUser(user *proto.User) error {
	if user == nil || user.Id == nil {
		return fmt.Errorf("cant't create user, user or userID is nil")
	}
	_, err := uch.cc.CreateUser(context.TODO(), user)
	if err != nil {
		return err
	}
	log.Println("successfully created user with UserID: ", user.Id)
	return err
}

func (uch *UserClientHandler) GetUser(userID uint64) (*proto.User, error) {

	user, err := uch.cc.GetUser(context.TODO(), &proto.UserID{Id: userID})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uch *UserClientHandler) UpdateUser(user *proto.User) error {
	if user == nil {
		return fmt.Errorf("updating user can't be nil")
	}
	_, err := uch.cc.UpdateUser(context.TODO(), user)
	if err != nil {
		return err
	}
	log.Println("successfully updated user")
	return nil
}

func (uch *UserClientHandler) DeleteUser(userId uint64) error {
	_, err := uch.cc.DeleteUser(context.TODO(), &proto.UserID{
		Id: userId,
	})
	return err
}
