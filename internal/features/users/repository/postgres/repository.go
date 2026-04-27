package users_postgres_repository

import core_postgres_conn "github.com/ioannuwu/git-diff-as-a-service/internal/core/repository/postgres/conn"

type UsersRepository struct {
	pool core_postgres_conn.Pool
}

func NewUsersRepository(pool core_postgres_conn.Pool) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
