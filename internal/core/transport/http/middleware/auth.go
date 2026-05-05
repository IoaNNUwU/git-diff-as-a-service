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

type SessionsRolesService interface {
	GetUserIDAndRoleForActiveSession(ctx context.Context, sessionID string) (int, string, error)
}

func NewAuthMiddleware(sessionsService SessionsRolesService) Middleware {
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

			userID, userRole, err := sessionsService.GetUserIDAndRoleForActiveSession(ctx, sessionKey.Value)
			if err != nil {
				err = fmt.Errorf("authentification failed: %w", core_errors.ErrExpiredSession)
				responseHandler.ErrorResponse(err, "authentification failed")
				return
			}

			ctx = context.WithValue(ctx, core_auth.UserID, userID)
			ctx = context.WithValue(ctx, core_auth.UserRole, userRole)

			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}
