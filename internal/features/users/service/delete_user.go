package users_service

import (
	"context"
	"fmt"

	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
)

func (s *UsersService) DeleteUser(ctx context.Context, id int) error {

	println(id)

	log := UserServiceLogger(ctx)

	if id < 0 {
		return fmt.Errorf("user id cannot be negative: %w", core_errors.ErrInvalidArgument)
	}

	if err := s.usersRepository.DeleteUser(ctx, id); err != nil {
		err := fmt.Errorf("unable to delete user: %w", err)
		log.Debug(err.Error())
		return err
	}

	return nil
}
