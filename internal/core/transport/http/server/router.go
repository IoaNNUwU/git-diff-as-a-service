package core_http_server

import (
	"fmt"
	"net/http"

	core_http_middleware "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/middleware"
)

type APIVersion string

var (
	V1 = APIVersion("v1")
	V2 = APIVersion("v2")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion APIVersion
}

func NewAPIVersionRouter(apiVersion APIVersion) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVersion: apiVersion,
	}
}

func (r *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Pattern)

		handler := core_http_middleware.Chain(route.Handler, route.Middlewares...)
		
		r.Handle(pattern, handler)
	}
}
