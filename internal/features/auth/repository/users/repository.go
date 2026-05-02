package auth_repository

import (
	"context"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	core_postgres_conn "github.com/ioannuwu/git-diff-as-a-service/internal/core/repository/postgres/conn"
	credentials_postgres_repository "github.com/ioannuwu/git-diff-as-a-service/internal/features/auth/repository/users/credentials/postgres"
	users_data_postgres_repository "github.com/ioannuwu/git-diff-as-a-service/internal/features/auth/repository/users/users_data/postgres"
	auth_service "github.com/ioannuwu/git-diff-as-a-service/internal/features/auth/service"
	"github.com/jackc/pgx/v5"
)

// *UsersRepository implements auth_service.UsersRepository
var _ auth_service.UsersRepository = &UsersRepository{}

type UsersRepository struct {
	credentialsRepository credentialsRepository
	usersRepository       usersDataRepository

	pgPool core_postgres_conn.Pool
}

func NewUsersRepository(
	credentialsRepository credentialsRepository,
	usersRepository usersDataRepository,
	pgPool core_postgres_conn.Pool,
) *UsersRepository {
	return &UsersRepository{
		credentialsRepository: credentialsRepository,
		usersRepository:       usersRepository,
		pgPool:                pgPool,
	}
}

func DefaultUsersRepository(pgPool core_postgres_conn.Pool) *UsersRepository {

	usersRepository := users_data_postgres_repository.NewUsersDataRepository(pgPool)
	credentialsRepository := credentials_postgres_repository.NewCredentialsRepository(pgPool)

	return NewUsersRepository(
		credentialsRepository,
		usersRepository,
		pgPool,
	)
}

type credentialsRepository interface {
	CreateUserIDFromCredentials(ctx context.Context, tx pgx.Tx, credentials domain.Credentials) (int, error)
	GetUserID(ctx context.Context, credentials domain.Credentials) (int, error)
}

type usersDataRepository interface {
	CreateUserData(ctx context.Context, tx pgx.Tx, user domain.User) (domain.User, error)
	DeleteUserData(ctx context.Context, tx pgx.Tx, id int) error
	GetUserData(ctx context.Context, id int) (domain.User, error)
}
