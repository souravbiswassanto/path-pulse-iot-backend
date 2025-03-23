package postgres

import (
	"context"
	"fmt"
	"github.com/souravbiswassanto/path-pulse-iot-backend/internal/models"
	custom_error "github.com/souravbiswassanto/path-pulse-iot-backend/pkg/lib"
)

type EventSqlDB struct {
	*PostgresClient
}

func NewEventSqlDB(pc *PostgresClient) (*EventSqlDB, error) {
	if pc == nil || pc.xc == nil {
		return nil, fmt.Errorf("client can't be nil")
	}
	return &EventSqlDB{pc}, nil
}

func (db *EventSqlDB) SetupTable() error {
	// this should create the table
	return db.xc.Sync(new(models.Event))
}

func (db *EventSqlDB) Get(ctx context.Context, v interface{}) (interface{}, error) {
	eventID, ok := v.(uint64)
	if !ok {
		return nil, fmt.Errorf("expected %T object, but got %T object", uint64(0), v)
	}
	if eventID == 0 {
		return nil, fmt.Errorf("eventID can't be 0")
	}
	event := new(models.Event)
	db.xc.SetDefaultContext(ctx)
	exists, err := db.xc.ID(eventID).Get(event)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, custom_error.ErrUserNotFound
	}
	return event, nil
}

func (db *EventSqlDB) Create(ctx context.Context, v interface{}) error {
	event, ok := v.(*models.Event)
	if !ok {
		return fmt.Errorf("expected %T object, but got %T object", &models.Event{}, v)
	}
	if event.EventID == 0 {
		return fmt.Errorf("eventID can't be 0")
	}
	_, err := db.Get(ctx, event.EventID)
	if err == nil {
		return custom_error.ErrUserAlreadyExists
	}
	db.xc.SetDefaultContext(ctx)
	_, err = db.xc.Insert(event)
	return err
}

func (db *EventSqlDB) Update(ctx context.Context, v interface{}) error {
	event, ok := v.(*models.Event)
	if !ok {
		return fmt.Errorf("expected %T object, but got %T object", &models.Event{}, v)
	}
	if event.EventID == 0 {
		return fmt.Errorf("eventID can't be 0")
	}
	_, err := db.Get(ctx, event.EventID)
	if err != nil && custom_error.IsUserNotFoundErr(err) {
		return db.Create(ctx, event)
	} else if err != nil {
		return fmt.Errorf("can't get the record with %v, err: %v", event.EventID, err)
	}
	db.xc.SetDefaultContext(ctx)
	_, err = db.xc.ID(event.EventID).Update(event)

	return err
}

func (db *EventSqlDB) Delete(ctx context.Context, v interface{}) error {
	eventID, ok := v.(int64)
	if !ok {
		return fmt.Errorf("expected %T object, but got %T object", uint64(0), v)
	}
	if eventID == 0 {
		return fmt.Errorf("eventID can't be 0")
	}
	event := new(models.Event)
	db.xc.SetDefaultContext(ctx)

	_, err := db.xc.ID(eventID).Delete(event)
	return err
}
