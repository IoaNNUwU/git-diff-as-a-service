package files_transport_http

import (
	"context"
	"net/http"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	core_http_server "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/server"
)

type FilesHTTPHandler struct {
	usersService FilesService
}

func NewFilesHTTPHandler(usersService FilesService) *FilesHTTPHandler {
	return &FilesHTTPHandler{
		usersService: usersService,
	}
}

type FilesService interface {
	CreateFile(ctx context.Context, file domain.File) (domain.File, error)
	DeleteFile(ctx context.Context, id int) error
}

func (h *FilesHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Pattern: "/files",
			Handler: h.CreateFile,
		},
		{
			Method:  http.MethodDelete,
			Pattern: "/files",
			Handler: h.DeleteFile,
		},
	}
}
