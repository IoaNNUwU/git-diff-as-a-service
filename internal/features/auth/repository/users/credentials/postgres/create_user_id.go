package auth_credentials_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *CredentialsRepository) CreateUserIDFromCredentials(ctx context.Context, tx pgx.Tx, credentials domain.Credentials) (int, error) {

	log := authCredentialsRepositoryPostgresLogger(ctx)

	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `

	WITH new_salt AS (SELECT gen_salt('bf', 8) AS new_salt)

	INSERT INTO git_diff_app.credentials (login, salt, password_hash)
	SELECT $1, new_salt, crypt($2, new_salt) FROM new_salt
	RETURNING user_id;

	`

	row := tx.QueryRow(ctx, query, credentials.Login, credentials.Password)

	var userID int
	err := row.Scan(&userID)

	if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
		if pgErr.Code == "23505" {
			err = fmt.Errorf("unable to create login: %w", core_errors.ErrAlreadyExists)
			log.Debug(err.Error())
			return domain.UninitializedID, err
		}
	}
	if err != nil {
		err := fmt.Errorf("scan error: %w", err)
		log.Debug(err.Error())
		return domain.UninitializedID, err
	}

	return userID, nil
}
