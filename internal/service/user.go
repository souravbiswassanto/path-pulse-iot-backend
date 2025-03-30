package service

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
)

type UserService struct {
	db    db.DB
	cache db.DB
}

func NewUserService(db db.DB, cache db.DB) *UserService {
	//ctx context.Context, maxCacheStoreLimit int, options ...func(*postgres.ClientOptions)
	//co := &postgres.ClientOptions{}
	//co.SetDefaultConnectionPooling()
	//for _, opt := range options {
	//	opt(co)
	//}
	//client := &postgres.PostgresClient{
	//	ClientOptions: co,
	//}
	//db, err := postgres.NewUserSqlDB(client)
	//if err != nil {
	//	return nil, err
	//}
	return &UserService{
		db:    db,
		cache: cache,
		// cache: in_memory.NewUserInMemoryStore(ctx, maxCacheStoreLimit),
	}
}

func (us *UserService) CreateUser(ctx context.Context, u *models.User) error {
	if u == nil {
		return fmt.Errorf("user is nil")
	}
	return us.db.Create(ctx, u)
}

func (us *UserService) UpdateUser(ctx context.Context, u *models.User) error {
	if u == nil {
		return fmt.Errorf("user is nil")
	}
	err := us.db.Update(ctx, u)
	if err != nil {
		return err
	}
	_ = us.cache.Update(ctx, u)
	return nil
}

func (us *UserService) DeleteUser(ctx context.Context, u models.UserID) error {
	err := us.db.Delete(ctx, u)
	if err != nil {
		return err
	}
	_ = us.cache.Delete(ctx, u)
	return nil
}

func (us *UserService) GetUser(ctx context.Context, u models.UserID) (*models.User, error) {
	v, err := us.cache.Get(ctx, u)
	if err == nil {
		user, ok := v.(*models.User)
		if ok {
			return user, nil
		}
	}
	k, err := us.db.Get(ctx, u)
	if err != nil {
		return nil, err
	}
	user, ok := k.(*models.User)
	if !ok {
		return nil, fmt.Errorf("can't decode user")
	}
	_ = us.cache.Create(ctx, user)
	return user, nil
}
