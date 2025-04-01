package db

import "context"

type DB interface {
	Create(ctx context.Context, v interface{}) error
	Update(ctx context.Context, v interface{}) error
	Delete(ctx context.Context, v interface{}) error
	Get(ctx context.Context, v interface{}) (interface{}, error)
}
