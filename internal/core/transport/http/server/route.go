package core_http_server

import (
	"net/http"

	core_http_middleware "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/middleware"
)

type Route struct {
	Method      string
	Pattern     string
	Handler     http.HandlerFunc
	Middlewares []core_http_middleware.Middleware
}
