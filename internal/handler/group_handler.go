package handler

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/service"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/group"
	"github.com/souravbiswassanto/path-pulse-iot-backend/protogen/golang/iot/user"
	"google.golang.org/grpc"
	"log"
)

type GroupServerHandler struct {
	svc *service.GroupService
	group.UnimplementedGroupManagerServer
}

func NewGroupServerHandler(svc *service.GroupService) *GroupServerHandler {
	return &GroupServerHandler{
		svc: svc,
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

type GroupClientHandler struct {
	cc group.GroupManagerClient
}

func NewGroupManagerClientHandler(cc grpc.ClientConnInterface) *GroupClientHandler {
	return &GroupClientHandler{
		cc: group.NewGroupManagerClient(cc),
	}
}

func (gch *GroupClientHandler) CreateGroup(group *group.Group) error {
	if group == nil || group.GId == 0 {
		return fmt.Errorf("cant't create group, group or groupID is nil")
	}
	_, err := gch.cc.CreateGroup(context.TODO(), group)
	if err != nil {
		return err
	}
	log.Println("successfully created group with GroupID: ", group.GId)
	return err
}

func (gch *GroupClientHandler) GetGroup(groupID uint64) (*group.Group, error) {

	group, err := gch.cc.GetGroup(context.TODO(), &group.GroupId{GId: groupID})
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (gch *GroupClientHandler) UpdateGroup(group *group.Group) error {
	if group == nil {
		return fmt.Errorf("updating group can't be nil")
	}
	_, err := gch.cc.UpdateGroup(context.TODO(), group)
	if err != nil {
		return err
	}
	log.Println("successfully updated group")
	return nil
}

func (gch *GroupClientHandler) DeleteGroup(groupGId uint64) error {
	_, err := gch.cc.DeleteGroup(context.TODO(), &group.GroupId{
		GId: groupGId,
	})
	return err
}

func (gch *GroupClientHandler) AddUserToGroup(ctx context.Context, ua *models.UserAdd) error {
	_, err := gch.cc.AddUserToGroup(ctx, &group.UserAdd{
		GroupId: ua.Gid,
		UserId:  uint64(ua.Uid),
	})

	return err
}
