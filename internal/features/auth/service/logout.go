package auth_service

import (
	"context"
	"fmt"
)

func (s *AuthService) Logout(ctx context.Context, sessionKey string) error {
	
	log := AuthServiceLogger(ctx)

	if err := s.sessionsRepository.CloseSession(ctx, sessionKey); err != nil {
		err := fmt.Errorf("unable to close session: %w", err)
		log.Debug(err.Error())
		return err
	}

	return nil
}