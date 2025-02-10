package custom_error

import "errors"

type CustomError struct {
	msg string
}

func (e *CustomError) Error() string {
	return e.msg
}

var (
	ErrKeyNotFound = &CustomError{"key not found"}
)

func IsKeyNotFoundErr(err error) bool {
	return errors.Is(err, ErrKeyNotFound)
}
