package auth_users_users_data_postgres_repository

import (
	core_postgres_conn "github.com/ioannuwu/git-diff-as-a-service/internal/core/repository/postgres/conn"
)

type usersDataRepository struct {
	pool core_postgres_conn.Pool
}

func NewUsersDataRepository(pool core_postgres_conn.Pool) *usersDataRepository {
	return &usersDataRepository{
		pool: pool,
	}
}
