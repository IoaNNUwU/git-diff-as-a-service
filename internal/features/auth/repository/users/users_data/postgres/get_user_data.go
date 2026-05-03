package auth_users_users_data_postgres_repository

import (
	"context"
	"fmt"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
)

func (r *UsersDataRepository) GetUserData(ctx context.Context, id int) (domain.User, error) {

	log := usersDataRepositoryPostgresLogger(ctx)

	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	query := `
	SELECT id, version, full_name, email
	FROM git_diff_app.users
	WHERE git_diff_app.users.id = $1;
	`

	row := r.pool.QueryRow(ctx, query, id)

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
