package handler

import (
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	proto "github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
)

func protoToModel(user *proto.User) *models.User {
	return &models.User{
		ID:   user.Id,
		Name: user.Name,
		Age:  user.Age,
		ContactInfo: models.ContactInfo{
			UserID:  user.Id,
			Email:   user.Email,
			Phone:   user.PhoneNo,
			Address: user.Address,
		},
		Factors: models.Factors{
			UserID:        user.Id,
			Height:        user.Height,
			Weight:        user.Weight,
			DiabeticLevel: user.DiabeticLevel,
			BP:            user.Bp,
		},
		Gender: user.Gender.String(),
	}
}

func modelToProto(user *models.User) *proto.User {
	return &proto.User{
		Id:            user.ID,
		Name:          user.Name,
		Age:           user.Age,
		Email:         user.ContactInfo.Email,
		PhoneNo:       user.ContactInfo.Phone,
		Address:       user.ContactInfo.Address,
		Height:        user.Factors.Height,
		Weight:        user.Factors.Weight,
		DiabeticLevel: user.Factors.DiabeticLevel,
		Bp:            user.Factors.BP,
		Gender:        proto.Gender(proto.Gender_value[user.Gender]),
	}
}
