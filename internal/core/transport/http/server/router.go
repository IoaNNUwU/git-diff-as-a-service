package core_http_server

import (
	"fmt"
	"net/http"
)

type ApiVersion string

var (
	V1 = ApiVersion("v1")
	V2 = ApiVersion("v2")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
}

func NewAPIVersionRouter(apiVersion ApiVersion) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVersion: apiVersion,
	}
}

func (r *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Pattern)

		r.Handle(pattern, route.Handler)
	}
}
