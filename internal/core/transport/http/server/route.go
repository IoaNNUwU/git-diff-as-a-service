package core_http_server

import "net/http"

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

func NewRoute(
	method string,
	path string,
	handler http.HandlerFunc,
) Route {
	return Route{
		Method:  method,
		Pattern: path,
		Handler: handler,
	}
}
