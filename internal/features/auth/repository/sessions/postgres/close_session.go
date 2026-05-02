package auth_sessions_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
)

func (s *SessionsRepository) CloseSession(ctx context.Context, id string) error {

	log := authSessionsRepositoryPostgresLogger(ctx)
	
	ctx, cancel := context.WithTimeout(ctx, s.pool.Timeout())
	defer cancel()

	query := `
	DELETE FROM git_diff_app.sessions
	WHERE git_diff_app.sessions.id = $1;
	`

	exec, err := s.pool.Exec(ctx, query, id)
	if err != nil {
		err := fmt.Errorf("execute delete: %w", err)
		log.Debug(err.Error())
		return err
	}
	if exec.RowsAffected() != 1 {
		err := fmt.Errorf("execute delete: %w", core_errors.ErrNotFound)
		log.Debug(err.Error())
		return core_errors.ErrNotFound
	}

	return nil
}