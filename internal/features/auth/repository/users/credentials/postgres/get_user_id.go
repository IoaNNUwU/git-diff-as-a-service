package auth_credentials_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *CredentialsRepository) GetUserID(
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
	WHERE
		login = $1
		AND password_hash = public.crypt($2, salt);
	`

	row := tx.QueryRow(ctx, query, credentials.Login, credentials.Password)

	var userID int
	err = row.Scan(&userID)
	if errors.Is(err, pgx.ErrNoRows) {
		err := fmt.Errorf("unable to get credentials: %w", core_errors.ErrWrongPassword)
		log.Debug(err.Error())
		return domain.UninitializedID, err
	}
	if err != nil {
		err := fmt.Errorf("unable to get credentials: %w", err)
		log.Debug(err.Error())
		return domain.UninitializedID, err
	}

	if err = tx.Commit(ctx); err != nil {
		err := fmt.Errorf("commit transaction failed: %w", err)
		log.Debug(err.Error())
		return domain.UninitializedID, err
	}

	return userID, nil
}
