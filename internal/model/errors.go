package model

import (
	"errors"
)

var (
	ErrorUnverifiedAccount = errors.New("login unverified")
	ErrorUserAlreadyExists = errors.New("user already exists")
	ErrorNoUser            = errors.New("user doesn't exists")

	ErrorWrongOrders        = errors.New("invalid order")
	ErrorOrderAlreadyExists = errors.New("order already exists")
	ErrorOrderConflict      = errors.New("order already register")

	ErrorInsufficientBalance = errors.New("insufficient balance")
)
