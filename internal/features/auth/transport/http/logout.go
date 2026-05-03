package auth_transport_http

import (
	"net/http"

	core_http_request "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/request"
	core_http_response "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/response"
)

type LogoutRequest struct {
	SessionKey string `json:"session_key"`
}

func (h *AuthHTTPHandler) Logout(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := AuthHTTPTransportLogger(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(rw, log)

	var request LogoutRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "unable to decode and validate HTTP request")
		return
	}

	if err := h.authService.Logout(ctx, request.SessionKey); err != nil {
		responseHandler.ErrorResponse(err, "unable to logout")
		return
	}

	responseHandler.StatusCodeResponse(http.StatusOK)
}
