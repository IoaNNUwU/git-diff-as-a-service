package auth_transport_http

import (
	"context"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
)

func AuthHTTPTransportLogger(ctx context.Context) *logger.Logger {
	log := logger.FromContext(ctx)

	return log.
		With("feature", "auth").
		With("layer", "transport/HTTP")
}
