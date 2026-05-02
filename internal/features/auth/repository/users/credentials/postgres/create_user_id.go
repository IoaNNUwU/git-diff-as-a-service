package auth_credentials_postgres_repository

import (
	"context"
	"fmt"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

func (r *credentialsRepository) CreateUserIDFromCredentials(ctx context.Context, tx pgx.Tx, credentials domain.Credentials) (int, error) {
	
	log := authCredentialsRepositoryPostgresLogger(ctx)

	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
	WITH new_salt AS (SELECT gen_salt('bf', 8))
	INSERT INTO git_diff_app.credentials (login, salt, password_hash)
	VALUES ($1, new_salt, crypt($2, new_salt))
	RETURNING user_id;
	`

	row := tx.QueryRow(ctx, query, credentials.Login, credentials.Password)

	var userID int
	err := row.Scan(&userID)
	if err != nil {
		err := fmt.Errorf("scan error: %w", err)
		log.Debug(err.Error())
		return domain.UninitializedID, err
	}

	return userID, nil
}