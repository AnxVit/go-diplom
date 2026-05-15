package repository

import (
	"errors"
)

var (
	ErrorLoginAlreadyExists = errors.New("login already exists")
	ErrorNoUser             = errors.New("user not exists")

	ErrorOrderAlreadyExists = errors.New("order already exists")
	ErrorOrderConflict      = errors.New("order already register")

	ErrorInsufficientBalance = errors.New("insufficient balance")
)
