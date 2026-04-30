package files_postgres_repository

import core_postgres_conn "github.com/ioannuwu/git-diff-as-a-service/internal/core/repository/postgres/conn"

type FilesRepository struct {
	pool core_postgres_conn.Pool
}

func NewFilesRepository(pool core_postgres_conn.Pool) *FilesRepository {
	return &FilesRepository{
		pool: pool,
	}
}
