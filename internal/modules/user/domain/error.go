package domain

import (
	
	"errors"
)

var (
	ErrUserNotFound = errors.New("User not found")
	ErrEmailAlreadyExits = errors.New("email already exits")
	ErrInvalidCredintials = errors.New("inavlid user or password")
	ErrUnauthorized = errors.New("Unauthorized")
	ErrInvalidToken = errors.New("invalid token")
)