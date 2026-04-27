package users_transport_http

import (
	"context"
	"net/http"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	core_http_server "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

type UsersService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Pattern: "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodDelete,
			Pattern: "/users",
			Handler: h.DeleteUser,
		},
	}
}
