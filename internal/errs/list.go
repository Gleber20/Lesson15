package errs

import "errors"

var (
	ErrNotfound           = errors.New("not found")
	ErrUserNotfound       = errors.New("user not found")
	ErrInvalidUserID      = errors.New("invalid user id")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidFieldValue  = errors.New("invalid field value")
)
