package domain

import "errors"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrIncorrectPassword = errors.New("incorect password")
var ErrUserNotFound = errors.New("user not found")
