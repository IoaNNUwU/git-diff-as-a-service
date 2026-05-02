package auth_sessions_postgres_repository

import (
	"context"
	"fmt"
	"time"
)

func (r *SessionsRepository) OpenSession(ctx context.Context, user_id int) (string, error) {

	log := authSessionsRepositoryPostgresLogger(ctx)

	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
	INSERT INTO git_diff_app.sessions (user_id, created_at, ttl)
	VALUES ($1, $2, $3)
	RETURNING id;
	`

	now := time.Now()
	row := r.pool.QueryRow(ctx, query, user_id, now, now.Add(1*time.Hour))

	var sessionID string
	err := row.Scan(&sessionID)
	if err != nil {
		err := fmt.Errorf("scan error: %w", err)
		log.Debug(err.Error())
		return "", err
	}

	return sessionID, nil
}
