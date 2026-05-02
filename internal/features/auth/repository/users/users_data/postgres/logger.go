package auth_users_users_data_postgres_repository

import (
	"context"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
)

func usersDataRepositoryPostgresLogger(ctx context.Context) *logger.Logger {
	log := logger.FromContext(ctx)

	return log.
		With("feature", "auth").
		With("layer", "repository/users/users/postgres")
}
