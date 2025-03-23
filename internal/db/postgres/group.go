package postgres

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	custom_error "github.com/souravbiswassanto/path-pulse-iot-backend/pkg/lib"
)

type GroupSqlDB struct {
	*PostgresClient
}

func NewGroupSqlDB(pc *PostgresClient) (*GroupSqlDB, error) {
	if pc == nil || pc.xc == nil {
		return nil, fmt.Errorf("client can't be nil")
	}
	return &GroupSqlDB{pc}, nil
}

func (db *GroupSqlDB) SetupTable() error {
	// this should create the table
	return db.xc.Sync(new(models.Group))
}

func (db *GroupSqlDB) Get(ctx context.Context, v interface{}) (interface{}, error) {
	GID, ok := v.(uint64)
	if !ok {
		return nil, fmt.Errorf("expected %T object, but got %T object", uint64(0), v)
	}
	if GID == 0 {
		return nil, fmt.Errorf("GID can't be 0")
	}
	group := new(models.Group)
	db.xc.SetDefaultContext(ctx)
	exists, err := db.xc.ID(GID).Get(group)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, custom_error.ErrUserNotFound
	}
	return group, nil
}

func (db *GroupSqlDB) Create(ctx context.Context, v interface{}) error {
	group, ok := v.(*models.Group)
	if !ok {
		return fmt.Errorf("expected %T object, but got %T object", &models.Group{}, v)
	}
	if group.GID == 0 {
		return fmt.Errorf("GID can't be 0")
	}
	_, err := db.Get(ctx, group.GID)
	if err == nil {
		return custom_error.ErrUserAlreadyExists
	}
	db.xc.SetDefaultContext(ctx)
	_, err = db.xc.Insert(group)
	return err
}

func (db *GroupSqlDB) Update(ctx context.Context, v interface{}) error {
	group, ok := v.(*models.Group)
	if !ok {
		return fmt.Errorf("expected %T object, but got %T object", &models.Group{}, v)
	}
	if group.GID == 0 {
		return fmt.Errorf("GID can't be 0")
	}
	_, err := db.Get(ctx, group.GID)
	if err != nil && custom_error.IsUserNotFoundErr(err) {
		return db.Create(ctx, group)
	} else if err != nil {
		return fmt.Errorf("can't get the record with %v, err: %v", group.GID, err)
	}
	db.xc.SetDefaultContext(ctx)
	_, err = db.xc.ID(group.GID).Update(group)

	return err
}

func (db *GroupSqlDB) Delete(ctx context.Context, v interface{}) error {
	GID, ok := v.(int64)
	if !ok {
		return fmt.Errorf("expected %T object, but got %T object", uint64(0), v)
	}
	if GID == 0 {
		return fmt.Errorf("GID can't be 0")
	}
	group := new(models.Group)
	db.xc.SetDefaultContext(ctx)

	_, err := db.xc.ID(GID).Delete(group)
	return err
}
