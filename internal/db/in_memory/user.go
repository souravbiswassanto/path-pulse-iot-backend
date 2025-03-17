package in_memory

import (
	"context"
	"fmt"
	ce "github.com/souravbiswassanto/path-pulse-iot-backend/internal/custom-error"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
)

type UserInMemoryStore struct {
	store *Store[models.UserID, *models.User]
}

func NewUserInMemoryStore(s *Store[models.UserID, *models.User]) *UserInMemoryStore {
	return &UserInMemoryStore{s}
}

func (us *UserInMemoryStore) Create(_ context.Context, v interface{}) error {
	u, ok := v.(*models.User)
	if !ok {
		return fmt.Errorf("expected %T, got %T", &models.User{}, v)
	}
	if u == nil {
		return fmt.Errorf("user is nil")
	}
	_, err := us.store.Get(u.ID)
	if err == nil {
		return fmt.Errorf("user already exists")
	}
	us.store.Create(u.ID, u)

	return nil
}

func (us *UserInMemoryStore) Update(_ context.Context, v interface{}) error {
	u, ok := v.(*models.User)
	if !ok {
		return fmt.Errorf("expected %T, got %T", &models.User{}, v)
	}
	if u == nil {
		return fmt.Errorf("user is nil")
	}
	_, err := us.store.Get(u.ID)
	if err != nil && ce.IsKeyNotFoundErr(err) {
		us.store.Create(u.ID, u)
		return nil
	}
	us.store.Update(u.ID, u)
	return nil
}

func (us *UserInMemoryStore) Delete(_ context.Context, v interface{}) error {
	u, ok := v.(models.UserID)
	if !ok {
		return fmt.Errorf("expected %T, got %T", uint64(0), v)
	}
	_, err := us.store.Get(u)
	if err != nil && ce.IsKeyNotFoundErr(err) {
		return nil
	}
	us.store.Delete(u)
	return nil
}

func (us *UserInMemoryStore) Get(_ context.Context, v interface{}) (interface{}, error) {
	u, ok := v.(models.UserID)
	if !ok {
		return nil, fmt.Errorf("expected %T, got %T", uint64(0), v)
	}
	k, err := us.store.Get(u)
	if err != nil {
		return nil, err
	}
	return k, nil
}
