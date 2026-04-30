package files_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
)

func (r *FilesRepository) DeleteFile(ctx context.Context, id int) error {

	log := FilesRepositoryPostgresLogger(ctx)

	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
	DELETE FROM git_diff_app.files
	WHERE git_diff_app.files.id = $1;
	`

	exec, err := r.pool.Exec(ctx, query, id)
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
