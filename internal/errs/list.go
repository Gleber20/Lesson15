package errs

import "errors"

var (
	ErrNotfound           = errors.New("not found")
	ErrProductNotfound    = errors.New("user not found")
	ErrInvalidProductID   = errors.New("invalid user id")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidFieldValue  = errors.New("invalid field value")
)
