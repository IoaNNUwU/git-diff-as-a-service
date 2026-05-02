package auth_service

import (
	"context"
	"fmt"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
)

func (s *AuthService) Login(ctx context.Context, credentials domain.Credentials) (string, error) {

	log := AuthServiceLogger(ctx)

	user, err := s.usersRepository.GetUser(ctx, credentials)
	if err != nil {
		err := fmt.Errorf("unable to get user: %w", err)
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