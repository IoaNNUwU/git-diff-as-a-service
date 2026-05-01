package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
	core_http_middleware "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/middleware"
)

type HTTPServer struct {
	log    *logger.Logger
	mux    *http.ServeMux
	config Config

	middlewares []core_http_middleware.Middleware
}

func NewHTTPServer(
	log *logger.Logger,
	config Config,
	middlewares ...core_http_middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		log:    log,
		mux:    http.NewServeMux(),
		config: config,

		middlewares: middlewares,
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {

	mux := core_http_middleware.Chain(h.mux, h.middlewares...)

	server := &http.Server{
		Addr:    h.config.Addr,
		Handler: mux,
	}

	ch := make(chan error)

	go func() {
		defer close(ch)

		var err error
		if h.config.HTTPS {
			h.log.Warn("HTTPS server starting on " + h.config.Addr)
			err = server.ListenAndServeTLS("keys/server.crt", "keys/server.key")
		} else {
			h.log.Warn("HTTP server starting on " + h.config.Addr)
			err = server.ListenAndServe()
		}

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("http server failed: %w", err)
		}
	case <-ctx.Done():
		h.log.Warn("HTTP server is shutting down")

		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			h.config.ShutdownTimeout,
		)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("unable to shutdown HTTP server: %w", err)
		}

		h.log.Warn("HTTP server stopped")
	}
	return nil
}

func (h *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {

		prefix := "/api/" + string(router.apiVersion)

		h.mux.Handle(prefix+"/", http.StripPrefix(prefix, router))
	}
}
