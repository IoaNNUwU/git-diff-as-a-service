package auth_repository

import (
	"context"
	"fmt"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
)

func (r *UsersRepository) GetUser(
	ctx context.Context,
	credentials domain.Credentials,
) (domain.User, error) {

	log := AuthRepositoryLogger(ctx)

	ctx, cancel := context.WithTimeout(ctx, r.pgPool.Timeout())
	defer cancel()

	userID, err := r.credentialsRepository.GetUserID(ctx, credentials)
	if err != nil {
		err := fmt.Errorf("unable to get user: %w", err)
		log.Debug(err.Error())
		return domain.User{}, err
	}

	user, err := r.usersRepository.GetUserData(ctx, userID)
	if err != nil {
		err := fmt.Errorf("unable to get user: %w", err)
		log.Debug(err.Error())
		return domain.User{}, err
	}

	return user, nil
}