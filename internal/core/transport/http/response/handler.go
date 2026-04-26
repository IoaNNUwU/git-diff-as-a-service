package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
)

type HTTPResponseHandler struct {
	rw  http.ResponseWriter
	log *logger.Logger
}

func NewHTTPResponseHandler(rw http.ResponseWriter, log *logger.Logger) *HTTPResponseHandler {
	return &HTTPResponseHandler{rw, log}
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError

	err := fmt.Errorf("unexpected panic: %v", p)
	h.log.Error(msg, slog.String("error", err.Error()))

	h.errorResponse(statusCode, err, msg)
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {

	var (
		statusCode int
		log        func(string, ...any)
	)

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		log = h.log.Warn

	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		log = h.log.Debug

	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		log = h.log.Warn

	default:
		statusCode = http.StatusInternalServerError
		log = h.log.Error
	}

	log(msg, slog.String("error", err.Error()))

	h.errorResponse(statusCode, err, msg)
}

func (h *HTTPResponseHandler) errorResponse(statusCode int, err error, msg string) {
	h.rw.WriteHeader(statusCode)

	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	h.JSONResponse(&response, statusCode)
}

func (h *HTTPResponseHandler) JSONResponse(responceBody any, statusCode int) {
	h.rw.WriteHeader(statusCode)

	if err := json.NewEncoder(h.rw).Encode(responceBody); err != nil {
		h.log.Error("unable to encode HTTP response", slog.String("error", err.Error()))
	}
}
