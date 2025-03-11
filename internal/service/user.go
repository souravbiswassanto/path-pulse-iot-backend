package service

import (
	"context"
	"fmt"
	ce "github.com/souravbiswassanto/path-pulse-iot-backend/internal/custom-error"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/db/in_memory"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
)

type UserService struct {
	userDb in_memory.UserStore[models.UserID, *models.User]
}

type US struct {
	db db.DB
}

func (us *US) CreateUser() {
	return us.db.Create()
}

type sql struct {
}

func (a *sql) Create(ctx context.Context, v interface{}) error {}

type inmem struct {
}

func NewUserService() *UserService {
	// eikhane db type ta lagbe
	userDb := in_memory.NewUserStore[models.UserID, *models.User]()
	id := uint64(1)
	userDb.Create((models.UserID)(id), &models.User{
		ID: (models.UserID)(id),
	})
	return &UserService{
		userDb: userDb,
	}
}

func (us *UserService) CreateUser(_ context.Context, u *models.User) error {
	if u == nil {
		return fmt.Errorf("user is nil")
	}
	_, err := us.userDb.Get(u.ID)
	if err == nil {
		return fmt.Errorf("user already exists")
	}
	us.userDb.Create(u.ID, u)

	return nil
}

func (us *UserService) UpdateUser(ctx context.Context, u *models.User) error {
	if u == nil {
		fmt.Errorf("user is nil")
	}
	_, err := us.userDb.Get(u.ID)
	if err != nil && ce.IsKeyNotFoundErr(err) {
		return us.CreateUser(ctx, u)
	}
	us.userDb.Update(u.ID, u)
	return nil
}

func (us *UserService) DeleteUser(_ context.Context, u models.UserID) error {
	_, err := us.userDb.Get(u)
	if err != nil && ce.IsKeyNotFoundErr(err) {
		return nil
	}
	us.userDb.Delete(u)
	return nil
}

func (us *UserService) GetUser(_ context.Context, u models.UserID) (*models.User, error) {
	k, err := us.userDb.Get(u)
	if err != nil {
		return nil, err
	}
	return k, nil
}
