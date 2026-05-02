package auth_service

import (
	"context"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	auth_transport_http "github.com/ioannuwu/git-diff-as-a-service/internal/features/auth/transport/http"
)

// *AuthService implements auth_transport_http.AuthService
var _ auth_transport_http.AuthService = &AuthService{}

type AuthService struct {
	sessionsRepository SessionsRepository
	usersRepository    UsersRepository
}

func NewAuthService(sessionsRepository SessionsRepository, usersRepository UsersRepository) *AuthService {
	return &AuthService{
		sessionsRepository: sessionsRepository,
		usersRepository:    usersRepository,
	}
}

type UsersRepository interface {
	CreateUser(ctx context.Context, user domain.User, credentials domain.Credentials) (domain.User, error)
	GetUser(ctx context.Context, credentials domain.Credentials) (domain.User, error)
}

type SessionsRepository interface {
	OpenSession(ctx context.Context, user_id int) (string, error)
	GetUserIDForActiveSession(ctx context.Context, sessionID string) (int, error)
	CloseSession(ctx context.Context, sessionID string) error
}
