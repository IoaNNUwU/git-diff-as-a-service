package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
)

type HTTPServer struct {
	log    *logger.Logger
	mux    *http.ServeMux
	config Config
}

func NewHTTPServer(
	log *logger.Logger,
	config Config,
) *HTTPServer {
	return &HTTPServer{
		mux:    http.NewServeMux(),
		config: config,
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:    h.config.Addr,
		Handler: h.mux,
	}

	ch := make(chan error)

	go func() {
		defer close(ch)

		h.log.Warn("HTTP server is starting")

		err := server.ListenAndServe()

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

		h.mux.Handle(prefix + "/", http.StripPrefix(prefix, router))
	}
}