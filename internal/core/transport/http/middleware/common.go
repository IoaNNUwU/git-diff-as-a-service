package core_http_middleware

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
	core_http_response "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/response"
)

const RequestIDHeader = "X-Request-ID"

func AddRequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(RequestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}
			r.Header.Set(RequestIDHeader, requestID)
			w.Header().Set(RequestIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func AddLogger(log *logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(RequestIDHeader)

			slogger := log.With(
				"request_id", requestID,
				"url", r.URL.String(),
			)

			log = &logger.Logger{
				Logger: *slogger,
			}

			ctx := context.WithValue(r.Context(), "log", log)

			log.Warn("Usable!!!!")

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RecoverPanic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			println("RECOVER")
			ctx := r.Context()
			log := logger.FromContext(ctx)

			log.Warn("HELLO")

			respHandler := core_http_response.NewHTTPResponseHandler(w, log)

			defer func() {
				if p := recover(); p != nil {
					respHandler.PanicResponse(p, "unexpected panic in HTTP handler")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := logger.FromContext(ctx)
			rw := core_http_response.NewResponseWriterWithStatusCode(w)

			before := time.Now()
			log.Debug("New HTTP request", slog.Time("time", before.UTC()))

			next.ServeHTTP(rw, r)

			statusCode := rw.GetStatusCode()

			after := time.Now()
			log.Debug("Finished HTTP request",
				slog.Time("time", after.UTC()),
				slog.Duration("time_spent", after.Sub(before)),
				slog.Int("status_code", statusCode),
			)
		})
	}
}
