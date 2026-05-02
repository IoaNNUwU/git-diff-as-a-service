package auth_transport_http

import (
	"net/http"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	core_http_request "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/request"
	core_http_response "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/response"
)

type LoginRequest struct {
	Username string `json:"email" validate:"required,min=5,max=100"`
	Password string `json:"password" validate:"required,min=10,max=100"`
}

type LoginResponse struct {
	SessionKey string `json:"session_key"`
}

func (h *AuthHTTPHandler) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := AuthHTTPTransportLogger(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(rw, log)

	var request LoginRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "unable to decode and validate HTTP request")
		return
	}

	sessionKey, err := h.authService.Login(ctx, domain.Credentials{
		Login:    request.Username,
		Password: request.Password,
	})
	if err != nil {
		responseHandler.ErrorResponse(err, "unable to login")
		return
	}

	response := LoginResponse{
		SessionKey: sessionKey,
	}

	responseHandler.JSONResponse(&response, http.StatusCreated)
}