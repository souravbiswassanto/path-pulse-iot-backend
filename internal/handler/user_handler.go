package handler

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/service"
	proto "github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	"gomodules.xyz/pointer"
)

type UserHandlerSer struct {
	svc *service.UserService
	proto.UnimplementedUserManagerServer
}

func NewUserHandler() *UserHandlerSer {
	return &UserHandlerSer{
		svc: service.NewUserService(),
	}
}

func (uh *UserHandlerSer) GetUser(ctx context.Context, uid *proto.UserID) (*proto.User, error) {
	if uid == nil {
		return nil, fmt.Errorf("given user id can't be nil")
	}
	user, err := uh.svc.GetUser(ctx, (*models.UserID)(&uid.Id))
	if err != nil {
		return nil, err
	}
	return modelToProto(user), nil
}

func (uh *UserHandlerSer) CreateUser(ctx context.Context, user *proto.User) (*proto.Empty, error) {
	if user == nil {
		return &proto.Empty{}, fmt.Errorf("given user is nil")
	}
	err := uh.svc.CreateUser(ctx, protoToModel(user))
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

func (uh *UserHandlerSer) UpdateUser(ctx context.Context, user *proto.User) (*proto.Empty, error) {
	if user == nil {
		return nil, fmt.Errorf("upating user can't be nil")
	}
	err := uh.svc.UpdateUser(ctx, protoToModel(user))
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

func (uh *UserHandlerSer) DeleteUser(ctx context.Context, uid *proto.UserID) (*proto.Empty, error) {
	if uid == nil {
		return nil, fmt.Errorf("the request user id can't be nil")
	}
	err := uh.svc.DeleteUser(ctx, (*models.UserID)(&uid.Id))
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

func protoToModel(user *proto.User) *models.User {
	return &models.User{
		ID:   (*models.UserID)(pointer.Uint64P(user.Id.GetId())),
		Name: user.Name,
		Age:  user.Age,
		ContactInfo: models.ContactInfo{
			UserID:  (*models.UserID)(pointer.Uint64P(user.Id.GetId())),
			Email:   user.Email,
			Phone:   user.PhoneNo,
			Address: user.Address,
		},
		Factors: models.Factors{
			UserID:        (*models.UserID)(pointer.Uint64P(user.Id.GetId())),
			Height:        user.Height,
			Weight:        user.Weight,
			DiabeticLevel: user.DiabeticLevel,
			BP: models.BloodPressure{
				Systolic:  user.Bp.GetSystolic(),
				Diastolic: user.Bp.GetDiastolic(),
			},
		},
		Gender: user.Gender.String(),
	}
}

func modelToProto(user *models.User) *proto.User {
	return &proto.User{
		Id:            &proto.UserID{Id: uint64(*user.ID)},
		Name:          user.Name,
		Age:           user.Age,
		Email:         user.ContactInfo.Email,
		PhoneNo:       user.ContactInfo.Phone,
		Address:       user.ContactInfo.Address,
		Height:        user.Factors.Height,
		Weight:        user.Factors.Weight,
		DiabeticLevel: user.Factors.DiabeticLevel,
		Bp: &proto.BloodPressure{
			Systolic:  user.Factors.BP.Systolic,
			Diastolic: user.Factors.BP.Diastolic,
		},
		Gender: proto.Gender(proto.Gender_value[user.Gender]),
	}
}
