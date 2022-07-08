package service

import "errors"

var (
	ErrInvalidInput = errors.New("input is invalid")
	ErrDatabase     = errors.New("can't connect to databases")
)
