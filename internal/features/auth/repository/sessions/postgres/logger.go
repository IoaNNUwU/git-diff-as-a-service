package auth_sessions_postgres_repository

import (
	"context"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
)

func authSessionsRepositoryPostgresLogger(ctx context.Context) *logger.Logger {
	
	log := logger.FromContext(ctx)

	return log.
		With("feature", "auth").
		With("layer", "repository/sessions/postgres")
}
