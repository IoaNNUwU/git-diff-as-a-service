package auth_users_users_data_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *usersDataRepository) DeleteUserData(ctx context.Context, tx pgx.Tx, id int) error {

	log := usersDataRepositoryPostgresLogger(ctx)

	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
	DELETE FROM git_diff_app.users
	WHERE git_diff_app.users.id = $1;
	`

	exec, err := tx.Exec(ctx, query, id)
	if err != nil {
		err := fmt.Errorf("execute delete: %w", err)
		log.Debug(err.Error())
		return core_errors.ErrTimeout
	}
	if exec.RowsAffected() != 1 {
		log.Debug("execute delete", "error", core_errors.ErrNotFound)
		return core_errors.ErrNotFound
	}

	return nil
}
