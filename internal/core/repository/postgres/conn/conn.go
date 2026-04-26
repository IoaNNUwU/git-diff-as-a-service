package core_postgres_conn

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnectionPool struct {
	*pgxpool.Pool
}

func MustNewConnectionPool(config Config, ctx context.Context) *ConnectionPool {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	pgxconfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		panic(fmt.Sprintf("unable to parse connection string %q: %s", connectionString, err))
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxconfig)
	if err != nil {
		panic(fmt.Sprintf("unable to create pgx pool: %s", err))
	}

	if err := pool.Ping(ctx); err != nil {
		panic(fmt.Sprintf("unable to ping database: %s", err))
	}

	return &ConnectionPool{Pool: pool}
}
