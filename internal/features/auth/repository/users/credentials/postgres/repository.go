package auth_credentials_postgres_repository

import (
	core_postgres_conn "github.com/ioannuwu/git-diff-as-a-service/internal/core/repository/postgres/conn"
)

type CredentialsRepository struct {
	pool core_postgres_conn.Pool
}

func NewCredentialsRepository(pool core_postgres_conn.Pool) *CredentialsRepository {
	return &CredentialsRepository{
		pool: pool,
	}
}
