package errs

import "errors"

var (
	ErrNotfound           = errors.New("not found")
	ErrUserNotfound       = errors.New("user not found")
	ErrInvalidUserID      = errors.New("invalid user id")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidFieldValue  = errors.New("invalid field value")
)


user, err := s.repo.GetUserById(ctx, userId)
if err != nil {
if errors.Is(err, errs.ErrNotfound) {
return models.User{}, errs.ErrUserNotfound
}
return models.User{}, err
}
return user, nil
}