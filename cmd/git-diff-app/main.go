package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/logger"
	core_postgres_conn "github.com/ioannuwu/git-diff-as-a-service/internal/core/repository/postgres/conn"
	core_http_middleware "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/middleware"
	core_http_server "github.com/ioannuwu/git-diff-as-a-service/internal/core/transport/http/server"
	auth_sessions_postgres_repository "github.com/ioannuwu/git-diff-as-a-service/internal/features/auth/repository/sessions/postgres"
	auth_repository "github.com/ioannuwu/git-diff-as-a-service/internal/features/auth/repository/users"
	// auth_users_users_data_postgres_repository "github.com/ioannuwu/git-diff-as-a-service/internal/features/auth/repository/users/users_data/postgres"
	auth_service "github.com/ioannuwu/git-diff-as-a-service/internal/features/auth/service"
	auth_transport_http "github.com/ioannuwu/git-diff-as-a-service/internal/features/auth/transport/http"
	files_postgres_repository "github.com/ioannuwu/git-diff-as-a-service/internal/features/files/repository/postgres"
	files_transport_http "github.com/ioannuwu/git-diff-as-a-service/internal/features/files/transport/http"
	// users_postgres_repository "github.com/ioannuwu/git-diff-as-a-service/internal/features/users/repository/postgres"
	// users_transport_http "github.com/ioannuwu/git-diff-as-a-service/internal/features/users/transport/http"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	log := logger.MustNewLogger(logger.MustNewConfig())
	defer log.Close()

	log.Debug("Starting git-diff-app")

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.V1)

	poolConf := core_postgres_conn.MustNewConfig()
	pool := core_postgres_conn.MustNewConnectionPool(ctx, poolConf)

	/*
	usersRepo := users_postgres_repository.NewUsersRepository(pool)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersRepo)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	*/

	usersRepo := auth_repository.DefaultUsersRepository(pool)
	sessionsRepo := auth_sessions_postgres_repository.NewSessionsRepository(ctx, log, pool)

	authService := auth_service.NewAuthService(sessionsRepo, usersRepo)

	coreHTTPServerConfig := core_http_server.MustNewConfig()

	authTransportHTTP := auth_transport_http.NewAuthHTTPHandler(authService, coreHTTPServerConfig.HTTPS)
	apiVersionRouter.RegisterRoutes(authTransportHTTP.Routes()...)

	filesRepo := files_postgres_repository.NewFilesRepository(pool)
	filesTransportHTTP := files_transport_http.NewFilesHTTPHandler(filesRepo)
	apiVersionRouter.RegisterRoutes(filesTransportHTTP.Routes()...)

	httpServer := core_http_server.NewHTTPServer(
		log,
		core_http_server.MustNewConfig(),

		core_http_middleware.AddRequestID(),
		core_http_middleware.AddLogger(log),
		core_http_middleware.RecoverPanic(),
		core_http_middleware.Trace(),
	)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		log.Error("HTTP server failed: " + err.Error())
	}
}
