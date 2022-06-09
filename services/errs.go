package services

import "errors"

var (
	ErrInvalidAmount = errors.New("invalid amount")
	ErrRepository    = errors.New("repository error")
)
