package errors

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrPassswordNotMatch = errors.New("password not match")
	ErrUsernameExist     = errors.New("username already taken")
)
