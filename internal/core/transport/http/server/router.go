package core_http_server

import (
	"fmt"
	"net/http"
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

		r.Handle(pattern, route.Handler)
	}
}
