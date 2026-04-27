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

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		h.log.Debug(msg, slog.String("error", err.Error()))
		h.errorResponse(http.StatusBadRequest, err, msg)

	case errors.Is(err, core_errors.ErrNotFound):
		h.log.Debug(msg, slog.String("error", err.Error()))
		h.errorResponse(http.StatusNotFound, err, msg)

	case errors.Is(err, core_errors.ErrConflict):
		h.log.Warn(msg, slog.String("error", err.Error()))
		h.errorResponse(http.StatusConflict, nil, msg)

	default:
		h.log.Error(msg, slog.String("error", err.Error()))
		h.errorResponse(http.StatusInternalServerError, nil, msg)
	}
}

func (h *HTTPResponseHandler) errorResponse(statusCode int, err error, msg string) {
	h.rw.WriteHeader(statusCode)

	var response map[string]string
	if err != nil {
		response = map[string]string{"message": msg, "error": err.Error()}
	} else {
		response = map[string]string{"message": msg}
	}

	h.JSONResponse(&response, statusCode)
}

func (h *HTTPResponseHandler) JSONResponse(responceBody any, statusCode int) {
	h.rw.WriteHeader(statusCode)

	if err := json.NewEncoder(h.rw).Encode(responceBody); err != nil {
		h.log.Error("unable to encode HTTP response", slog.String("error", err.Error()))
	}
}
