package users_service

import (
	"context"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
}

func NewUsersService(usersRepository UsersRepository) UsersService {
	return UsersService{
		usersRepository: usersRepository,
	}
}

type UsersRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}