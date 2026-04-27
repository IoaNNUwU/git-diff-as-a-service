package users_service

import (
	"context"
	"fmt"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
)

func (s *UsersService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {

	log := UserServiceLogger(ctx)

	if err := user.Validate(); err != nil {
		err := fmt.Errorf("invalid user: %w", err)
		log.Debug(err.Error())
		return domain.User{}, err
	}

	user, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		err := fmt.Errorf("unable to create user: %w", err)
		log.Debug(err.Error())
		return domain.User{}, err
	}

	return user, nil
}
