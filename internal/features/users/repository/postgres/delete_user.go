package users_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
)

func (r *UsersRepository) DeleteUser(ctx context.Context, id int) error {

	log := UsersRepositoryPostgresLogger(ctx)

	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
	DELETE FROM git_diff_app.users
	WHERE git_diff_app.users.id = $1;
	`

	exec, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		err := fmt.Errorf("execute delete: %w", err)
		log.Debug(err.Error())
		return ErrTimeout
	}
	if exec.RowsAffected() != 1 {
		log.Debug("execute delete", "error", core_errors.ErrNotFound)
		return core_errors.ErrNotFound
	}

	return nil
}
