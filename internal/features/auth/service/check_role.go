package auth_service

import (
	"context"
	"fmt"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
)

func (s *AuthService) GetUserIDAndRoleForActiveSession(ctx context.Context, sessionID string) (int, string, error) {

	log := AuthServiceLogger(ctx)

	userID, err := s.sessionsRepository.GetUserIDForActiveSession(ctx, sessionID)
	if err != nil {
		err := fmt.Errorf("unable to get user id: %w", err)
		log.Debug(err.Error())
		return domain.UninitializedID, "", err
	}

	userRole, err := s.usersRepository.GetUserRole(ctx, userID)
	if err != nil {
		err := fmt.Errorf("unable to get user role: %w", err)
		log.Debug(err.Error())
		return domain.UninitializedID, "", err
	}

	return userID, userRole, nil
}
