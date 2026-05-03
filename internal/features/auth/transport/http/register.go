package auth_transport_http

import (
	"net/http"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	core_http_request "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/request"
	core_http_response "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/response"
)

type registerRequest struct {
	Login    string `json:"login" validate:"required,min=5,max=100"`
	Password string `json:"password" validate:"required,min=10,max=100"`

	FullName string  `json:"full_name" validate:"required,min=3,max=50"`
	Email    *string `json:"email" validate:"min=5,max=50"`
}

type registerResponse struct {
	SessionKey string `json:"session_key"`
}

func (h *AuthHTTPHandler) Register(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := AuthHTTPTransportLogger(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(rw, log)

	var request registerRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "unable to decode and validate HTTP request")
		return
	}

	credentials := domain.Credentials{
		Login:    request.Login,
		Password: request.Password,
	}

	user := domain.User{
		FullName: request.FullName,
		Email:    request.Email,
	}

	sessionKey, err := h.authService.Register(ctx, credentials, user)
	if err != nil {
		responseHandler.ErrorResponse(err, "unable to register")
		return
	}

	response := registerResponse{
		SessionKey: sessionKey,
	}

	responseHandler.JSONResponse(&response, http.StatusCreated)
}
