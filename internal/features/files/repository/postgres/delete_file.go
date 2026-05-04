package files_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
	"github.com/jackc/pgx/v5"
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


func (r *FilesRepository) DeleteFileByAuthor(ctx context.Context, id int, authorID int) error {

	log := FilesRepositoryPostgresLogger(ctx)

	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		err := fmt.Errorf("begin transaction failed: %w", err)
		log.Debug(err.Error())
		return core_errors.ErrInternal
	}
	defer tx.Rollback(ctx)

	query := `
	SELECT author_id
	FROM git_diff_app.files
	WHERE git_diff_app.files.id = $1
	`

	row := tx.QueryRow(ctx, query, id)

	var fileAuthorID int
	err = row.Scan(&authorID)

	if errors.Is(err, pgx.ErrNoRows) {
		err := fmt.Errorf("execute delete: %w", core_errors.ErrNotFound)
		log.Debug(err.Error())
		return core_errors.ErrNotFound
	}
	if err != nil {
		err := fmt.Errorf("execute delete: %w", err)
		log.Debug(err.Error())
		return core_errors.ErrInternal
	}
	if fileAuthorID != authorID {
		err := fmt.Errorf("unable to delete file owned by another user: %w", core_errors.ErrPermission)
		log.Debug(err.Error())
		return core_errors.ErrPermission
	}

	query = `
	DELETE FROM git_diff_app.files
	WHERE git_diff_app.files.id = $1;
	`

	exec, err := tx.Exec(ctx, query, id)
	if err != nil {
		err := fmt.Errorf("execute delete: %w", err)
		log.Debug(err.Error())
		return core_errors.ErrInternal
	}
	if exec.RowsAffected() != 1 {
		log.Debug("execute delete", "error", core_errors.ErrNotFound)
		return core_errors.ErrNotFound
	}
	if err = tx.Commit(ctx); err != nil {
		err := fmt.Errorf("unable to commit transaction: %w", err)
		log.Debug(err.Error())
		return core_errors.ErrInternal
	}

	return nil
}
