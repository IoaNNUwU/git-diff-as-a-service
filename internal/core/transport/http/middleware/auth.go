package core_http_middleware

import (
	"context"
	"fmt"
	"net/http"

	core_auth "github.com/ioannuwu/git-diff-as-a-service/internal/core/auth"
	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
	core_http_response "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/response"
)

type SessionsService interface {
	GetUserIDForActiveSession(ctx context.Context, sessionID string) (int, error)
}

func NewAuthMiddleware(sessionsService SessionsService) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			log := logger.FromContext(ctx)

			responseHandler := core_http_response.NewHTTPResponseHandler(rw, log)

			sessionKey, err := r.Cookie(core_auth.SessionCookie)
			if err != nil {
				err = fmt.Errorf("unable to get session key: %w", core_errors.ErrUnauthorized)
				responseHandler.ErrorResponse(err, "unable to get session key")
				return
			}

			userID, err := sessionsService.GetUserIDForActiveSession(ctx, sessionKey.Value)
			if err != nil {
				err = fmt.Errorf("authentification failed: %w", core_errors.ErrExpiredSession)
				responseHandler.ErrorResponse(err, "authentification failed")
				return
			}

			ctx = context.WithValue(ctx, core_auth.UserID, userID)
			ctx = context.WithValue(ctx, core_auth.UserRole, core_auth.RoleUser)

			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}
