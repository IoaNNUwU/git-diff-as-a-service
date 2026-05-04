package auth_transport_http

import (
	"context"
	"net/http"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	core_http_server "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/server"
)

var (
	SessionKeyCookie = "session_key"
)

type AuthHTTPHandler struct {
	authService AuthService
	useHTTPS bool
}

func NewAuthHTTPHandler(authService AuthService, useHTTPS bool) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		authService: authService,
		useHTTPS: useHTTPS,
	}
}

type AuthService interface {
	Register(ctx context.Context, credentials domain.Credentials, user domain.User) (string, error)
	Login(ctx context.Context, credentials domain.Credentials) (string, error)
	Logout(ctx context.Context, sessionKey string) error
}

func (h *AuthHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Pattern: "/auth/login",
			Handler: h.Login,
		},
		{
			Method:  http.MethodDelete,
			Pattern: "/auth/logout",
			Handler: h.Logout,
		},
		{
			Method:  http.MethodPost,
			Pattern: "/auth/register",
			Handler: h.Register,
		},
	}
}
