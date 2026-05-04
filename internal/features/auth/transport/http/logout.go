package auth_transport_http

import (
	"net/http"

	core_auth "github.com/ioannuwu/git-diff-as-a-service/internal/core/auth"
	core_http_response "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/response"
)

func (h *AuthHTTPHandler) Logout(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := AuthHTTPTransportLogger(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(rw, log)

	sessionKey, err := r.Cookie(core_auth.SessionCookie)
	if err != nil {
		responseHandler.ErrorResponse(err, "unable to decode and validate HTTP request")
		return
	}

	if err := h.authService.Logout(ctx, sessionKey.Value); err != nil {
		responseHandler.ErrorResponse(err, "unable to logout")
		return
	}

	responseHandler.StatusCodeResponse(http.StatusOK)
}
