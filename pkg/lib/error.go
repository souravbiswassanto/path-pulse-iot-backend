package lib

import "errors"

type CustomError struct {
	msg string
}

func (e *CustomError) Error() string {
	return e.msg
}

var (
	ErrKeyNotFound       = &CustomError{"keys not found"}
	ErrUserNotFound      = &CustomError{"User not found"}
	ErrUserAlreadyExists = &CustomError{"User already exists"}
	ErrFileNotExists     = &CustomError{"File does not exits"}
)

func IsKeyNotFoundErr(err error) bool {
	return errors.Is(err, ErrKeyNotFound)
}

func IsUserNotFoundErr(err error) bool {
	return errors.Is(err, ErrUserNotFound)
}

func IsFileNotExistsErr(err error) bool {
	return errors.Is(err, ErrFileNotExists)
}
