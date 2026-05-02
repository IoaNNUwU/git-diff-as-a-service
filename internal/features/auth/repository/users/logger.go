package auth_repository

import (
	"context"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
)

func AuthRepositoryLogger(ctx context.Context) *logger.Logger {
	log := logger.FromContext(ctx)

	return log.
		With("feature", "auth").
		With("layer", "repository")
}
