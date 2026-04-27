package users_postgres_repository

import "errors"

var (
	ErrTimeout = errors.New("connection timed out")
)
