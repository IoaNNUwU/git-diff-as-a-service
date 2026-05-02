package core_errors

import "errors"

var (
	ErrUnauthorized    = errors.New("unauthorized")
	ErrExpiredSession  = errors.New("session expired")

	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("invalid argument")

	ErrConflict        = errors.New("conflict")

	ErrTimeout         = errors.New("connection timed out")
	ErrInternal        = errors.New("internal server error")
)
