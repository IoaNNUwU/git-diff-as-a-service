package auth_sessions_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
)

func (r *SessionsRepository) runCron(ctx context.Context, log *logger.Logger) {

	log = log.With("feature", "auth").
		With("layer", "repository/sessions/postgres")

	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		log.Debug("cron job started: clear expired sessions")
		defer log.Debug("cron job stopped: clear expired sessions")
		for {
			select {
			case <-ticker.C:
				r.clearExpiredSessions(ctx, log)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (r *SessionsRepository) clearExpiredSessions(ctx context.Context, log *logger.Logger) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
	DELETE FROM git_diff_app.sessions
	WHERE git_diff_app.sessions.ttl < now();
	`

	tag, err := r.pool.Exec(ctx, query)
	if err != nil {
		err := fmt.Errorf("execute delete: %w", err)
		log.Debug(err.Error())
	}
	log.Debug("cron job: expired sessions deleted sucessfuly", "closed_sessions", tag.RowsAffected())
}
