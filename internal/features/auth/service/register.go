package auth_service

import (
	"context"
	"fmt"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
)

func (s *AuthService) Register(ctx context.Context, credentials domain.Credentials, user domain.User) (string, error) {

	log := AuthServiceLogger(ctx)

	user, err := s.usersRepository.CreateUser(ctx, user, credentials)
	if err != nil {
		err := fmt.Errorf("unable to create user: %w", err)
		log.Debug(err.Error())
		return "", err
	}

	token, err := s.sessionsRepository.OpenSession(ctx, user.ID)
	if err != nil {
		err := fmt.Errorf("unable to open session: %w", err)
		log.Debug(err.Error())
		return "", err
	}

	return token, nil
}