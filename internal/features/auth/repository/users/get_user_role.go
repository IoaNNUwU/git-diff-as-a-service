package auth_repository

import (
	"context"
	"fmt"
)

func (r *UsersRepository) GetUserRole(ctx context.Context, userID int) (string, error) {

	log := AuthRepositoryLogger(ctx)

	user, err := r.usersRepository.GetUserData(ctx, userID)
	if err != nil {
		err := fmt.Errorf("unable to get user: %w", err)
		log.Debug(err.Error())
		return "", err
	}

	return user.Role, nil
}
