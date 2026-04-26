package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
	core_http_server "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/server"
	users_transport_http "github.com/ioannuwu/git-diff-as-a-service/internal/features/users/transport/http"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	log := logger.MustNewLogger(logger.MustNewConfig())
	defer log.Close()

	log.Debug("Starting git-diff-app")

	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(nil)
	usersTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.V1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)

	httpServer := core_http_server.NewHTTPServer(log, core_http_server.MustNewConfig())

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		log.Error("HTTP server failed: " + err.Error())
	}
}
