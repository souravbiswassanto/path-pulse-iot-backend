package in_memory

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
)

type GroupInMemoryStore struct {
	store *Store[uint64, *models.Group]
}

func NewGroupInMemoryStore(s *Store[uint64, *models.Group]) *GroupInMemoryStore {
	return &GroupInMemoryStore{s}
}

func (gs *GroupInMemoryStore) Create(_ context.Context, v interface{}) error {
	g, ok := v.(*models.Group)
	if !ok {
		return fmt.Errorf("expected %T, got %T", &models.Group{}, v)
	}
	if g == nil || g.GID == 0 {
		return fmt.Errorf("group can't be nil")
	}

	gs.store.Create(g.GID, g)
	return nil
}

func (gs *GroupInMemoryStore) Update(_ context.Context, v interface{}) error {
	g, ok := v.(*models.Group)
	if !ok {
		return fmt.Errorf("expected %T, got %T", &models.Group{}, v)
	}
	if g == nil || g.GID == 0 {
		return fmt.Errorf("group can't be nil")
	}
	gs.store.Update(g.GID, g)

	return nil
}

func (gs *GroupInMemoryStore) Delete(_ context.Context, v interface{}) error {
	groupID, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("got %v, expected uint64 eventID", v)
	}
	if groupID == 0 {
		return fmt.Errorf("groupID can't be 0")
	}
	gs.store.Delete(groupID)
	return nil
}

func (gs *GroupInMemoryStore) Get(_ context.Context, v interface{}) (interface{}, error) {
	groupID, ok := v.(uint64)
	if !ok {
		return nil, fmt.Errorf("got %v, expected uint64 eventID", v)
	}
	if groupID == 0 {
		return nil, fmt.Errorf("groupID can't be 0")
	}
	g, err := gs.store.Get(groupID)
	if err != nil {
		return nil, err
	}
	return g, nil
}
