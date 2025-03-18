package handler

import (
	"context"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/service"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/group"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
)

type GroupServerHandler struct {
	svc *service.GroupService
	group.UnimplementedGroupManagerServer
}

func NewGroupServerHandler() *GroupServerHandler {
	return &GroupServerHandler{
		// TODO: need to fix
		svc: service.NewGroupService(nil, nil),
	}
}

func (gsh *GroupServerHandler) CreateGroup(ctx context.Context, group *group.Group) (*user.Empty, error) {
	return &user.Empty{}, gsh.svc.AddGroup(ctx, groupProtoToModel(group))
}

func (gsh *GroupServerHandler) UpdateGroup(ctx context.Context, group *group.Group) (*user.Empty, error) {
	return &user.Empty{}, gsh.svc.UpdateGroup(ctx, groupProtoToModel(group))
}

func (gsh *GroupServerHandler) DeleteGroup(ctx context.Context, groupId *group.GroupId) (*user.Empty, error) {
	return &user.Empty{}, gsh.svc.DeleteGroup(ctx, groupId.GId)
}

func (gsh *GroupServerHandler) GetGroup(ctx context.Context, groupId *group.GroupId) (*group.Group, error) {
	se, err := gsh.svc.GetSingleGroupDetails(ctx, groupId.GId)
	if err != nil {
		return nil, err
	}
	return groupModelToProto(se), nil
}

func (gsh *GroupServerHandler) AddUserToGroup(ctx context.Context, ua *group.UserAdd) (*user.Empty, error) {
	return nil, gsh.svc.AddUserToGroup(ctx, models.UserAdd{Uid: models.UserID(ua.UserId), Gid: ua.GroupId})
}
