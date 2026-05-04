package auth_sessions_postgres_repository

import (
	"context"
	"fmt"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
)

func (s *SessionsRepository) GetUserIDForActiveSession(ctx context.Context, sessionID string) (int, error) {

	log := authSessionsRepositoryPostgresLogger(ctx)

	ctx, cancel := context.WithTimeout(ctx, s.pool.Timeout())
	defer cancel()

	query := `
	SELECT user_id, (SELECT git_diff_app.sessions.ttl < now()) AS expired
	FROM git_diff_app.sessions
	WHERE git_diff_app.sessions.id = $1;
	`

	row := s.pool.QueryRow(ctx, query, sessionID)

	var expired bool
	var userID int
	err := row.Scan(&userID, &expired)
	if expired {
		err := fmt.Errorf("expired: %s", sessionID)
		log.Debug(err.Error())
		return domain.UninitializedID, err
	}
	if err != nil {
		err := fmt.Errorf("scan error: %w", err)
		log.Debug(err.Error())
		return domain.UninitializedID, err
	}

	return userID, nil
}
