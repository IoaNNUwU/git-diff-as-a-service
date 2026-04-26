package core_http_response

import "net/http"

const StatusCodeUninitialized = -1

type ResponseWriterWithStatusCode struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriterWithStatusCode(w http.ResponseWriter) *ResponseWriterWithStatusCode {
	return &ResponseWriterWithStatusCode{
		ResponseWriter: w,
		statusCode:     StatusCodeUninitialized,
	}
}

func (w *ResponseWriterWithStatusCode) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func (w *ResponseWriterWithStatusCode) GetStatusCode() int {
	if w.statusCode == StatusCodeUninitialized {
		panic("unable to get status code: status code wasn't set")
	}
	return w.statusCode
}
