package core_http_response

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	h.log.Error(msg + ": " + err.Error())
	h.rw.WriteHeader(statusCode)

	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	if err := json.NewEncoder(h.rw).Encode(response); err != nil {
		h.log.Error("unable to encode http response: " + err.Error())
	}
}
