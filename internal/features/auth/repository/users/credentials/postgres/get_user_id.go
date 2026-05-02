package auth_credentials_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *credentialsRepository) GetUserID(
	ctx context.Context,
	credentials domain.Credentials,
) (int, error) {

	log := authCredentialsRepositoryPostgresLogger(ctx)

	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.RepeatableRead,
		AccessMode: pgx.ReadOnly,
	})
	if err != nil {
		err := fmt.Errorf("unable to begin transaction: %w", err)
		log.Debug(err.Error())
		return domain.UninitializedID, err
	}
	defer tx.Rollback(ctx)

	query := `
	SELECT
	FROM git_diff_app.credentials
	WHERE login = $1;
	`
	_, err = tx.Exec(ctx, query, credentials.Login)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.UninitializedID, core_errors.ErrNotFound
	}
	if err != nil {
		err := fmt.Errorf("unable to get credentials: %w", err)
		log.Debug(err.Error())
		return domain.UninitializedID, err
	}

	query = `
	SELECT user_id
	FROM git_diff_app.credentials
	WHERE login = $1 
	AND password_hash = crypt($2, salt);
	`
	row, err := tx.Query(ctx, query, credentials.Login, credentials.Password)
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.UninitializedID, core_errors.ErrUnauthorized
	}
	if err != nil {
		err := fmt.Errorf("unable to get credentials: %w", err)
		log.Debug(err.Error())
		return domain.UninitializedID, err
	}

	var user_id int
	if err := row.Scan(&user_id); err != nil {
		err := fmt.Errorf("unable to get credentials: %w", err)
		log.Debug(err.Error())
		return domain.UninitializedID, err
	}

	if err = tx.Commit(ctx); err != nil {
		err := fmt.Errorf("unable to commit transaction: %w", err)
		log.Debug(err.Error())
		return domain.UninitializedID, err
	}

	return user_id, nil
}
