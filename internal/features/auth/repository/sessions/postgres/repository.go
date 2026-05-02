package auth_sessions_postgres_repository

import (
	core_postgres_conn "github.com/ioannuwu/git-diff-as-a-service/internal/core/repository/postgres/conn"
	auth_service "github.com/ioannuwu/git-diff-as-a-service/internal/features/auth/service"
)

// *SessionsRepository implements auth_service.SessionsRepository
var _ auth_service.SessionsRepository = &SessionsRepository{}

type SessionsRepository struct {
	pool core_postgres_conn.Pool
}

func NewSessionsRepository(pool core_postgres_conn.Pool) *SessionsRepository {
	return &SessionsRepository{
		pool: pool,
	}
}
