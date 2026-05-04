package files_transport_http

import (
	"context"
	"net/http"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	core_http_middleware "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/middleware"
	core_http_server "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/server"
)

type FilesHTTPHandler struct {
	filesService FilesService
}

func NewFilesHTTPHandler(usersService FilesService) *FilesHTTPHandler {
	return &FilesHTTPHandler{
		filesService: usersService,
	}
}

type FilesService interface {
	CreateFile(ctx context.Context, file domain.File) (domain.File, error)
	DeleteFile(ctx context.Context, id int) error
}

func (h *FilesHTTPHandler) Routes(authMiddleware core_http_middleware.Middleware) []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Pattern: "/files",
			Handler: h.CreateFile,
			Middlewares: []core_http_middleware.Middleware{
				authMiddleware,
			},
		},
		{
			Method:  http.MethodDelete,
			Pattern: "/files",
			Handler: h.DeleteFile,
			Middlewares: []core_http_middleware.Middleware{
				authMiddleware,
			},
		},
	}
}
