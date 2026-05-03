package auth_users_users_data_postgres_repository

import (
	"context"
	"fmt"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

func (r *UsersDataRepository) CreateUserData(
	ctx context.Context,
	tx pgx.Tx,
	user domain.User,
) (domain.User, error) {

	log := usersDataRepositoryPostgresLogger(ctx)

	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
	INSERT INTO git_diff_app.users (id, full_name, email)
	VALUES ($1, $2, $3)
	RETURNING id, version, full_name, email;
	`

	row := tx.QueryRow(ctx, query, user.ID, user.FullName, user.Email)

	var userModel userDataModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.Email,
	)
	if err != nil {
		err := fmt.Errorf("scan error: %w", err)
		log.Debug(err.Error())
		return domain.User{}, err
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.Email,
	)

	return userDomain, nil
}
