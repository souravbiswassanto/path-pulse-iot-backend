package service

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
)

type GroupService struct {
	groupDb db.DB
	cache   db.DB
}

func NewGroupService(db db.DB, cache db.DB) *GroupService {
	return &GroupService{
		groupDb: db,
		cache:   cache,
	}
}

func (gs *GroupService) AddGroup(ctx context.Context, group *models.Group) error {
	if group == nil || group.GID == 0 {
		return fmt.Errorf("group can't be nil")
	}
	return gs.groupDb.Create(ctx, group)
}

func (gs *GroupService) UpdateGroup(ctx context.Context, group *models.Group) error {
	if group == nil || group.GID == 0 {
		return fmt.Errorf("group can't be nil")
	}
	err := gs.groupDb.Update(ctx, group)
	if err != nil {
		return err
	}
	_ = gs.cache.Update(ctx, group)
	return nil
}

func (gs *GroupService) DeleteGroup(ctx context.Context, groupId uint64) error {
	if groupId == 0 {
		return fmt.Errorf("groupId can't be 0")
	}
	err := gs.groupDb.Delete(ctx, groupId)
	if err != nil {
		return err
	}
	_ = gs.cache.Delete(ctx, groupId)
	return nil
}

func (gs *GroupService) GetSingleGroupDetails(ctx context.Context, groupId uint64) (*models.Group, error) {
	v, err := gs.cache.Get(ctx, groupId)
	if err == nil {
		group, ok := v.(*models.Group)
		if ok {
			return group, nil
		}
		// need a log here
	}
	k, err := gs.groupDb.Get(ctx, groupId)
	if err != nil {
		return nil, err
	}
	group, ok := k.(*models.Group)
	if !ok {
		return nil, fmt.Errorf("can't decode group, expected %T, got %T", &models.Group{}, k)
	}
	_ = gs.cache.Create(ctx, group)
	return group, nil
}

func (gs *GroupService) AddUserToGroup(ctx context.Context, ua models.UserAdd) error {
	g, err := gs.GetSingleGroupDetails(ctx, ua.Gid)
	if err != nil {
		return err
	}
	found := false
	for _, v := range g.Members {
		if *v == ua.Uid {
			found = true
			break
		}
	}
	if !found {
		g.Members = append(g.Members, &ua.Uid)
		err := gs.UpdateGroup(ctx, g)
		if err != nil {
			return err
		}
	}
	return nil
}
