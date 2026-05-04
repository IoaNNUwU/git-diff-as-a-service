package core_errors

import "errors"

var (
	ErrUnauthorized    = errors.New("unauthorized")
	ErrExpiredSession  = errors.New("session expired")
	ErrWrongPassword   = errors.New("wrong password")
	ErrPermission      = errors.New("insufficient permission")

	ErrNotFound        = errors.New("not found")
	ErrAlreadyExists   = errors.New("already exists")

	ErrInvalidArgument = errors.New("invalid argument")

	ErrConflict        = errors.New("conflict")

	ErrTimeout         = errors.New("connection timed out")
	ErrInternal        = errors.New("internal server error")
)
