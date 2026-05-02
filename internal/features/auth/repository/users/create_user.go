package auth_repository

import (
	"context"
	"fmt"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
	credentials domain.Credentials,
) (domain.User, error) {

	log := AuthRepositoryLogger(ctx)

	ctx, cancel := context.WithTimeout(ctx, r.pgPool.Timeout())
	defer cancel()

	tx, err := r.pgPool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	defer tx.Rollback(ctx)

	if err != nil {
		err := fmt.Errorf("unable to begin transaction: %w", err)
		log.Debug(err.Error())
		return domain.User{}, err
	}

	userID, err := r.credentialsRepository.CreateUserIDFromCredentials(ctx, tx, credentials)
	if err != nil {
		err := fmt.Errorf("unable to save user credentials: %w", err)
		log.Debug(err.Error())
		return domain.User{}, err
	}
	user.ID = userID

	user, err = r.usersRepository.CreateUserData(ctx, tx, user)
	if err != nil {
		err := fmt.Errorf("unable to save user info: %w", err)
		log.Debug(err.Error())
		return domain.User{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		err := fmt.Errorf("commit transaction failed: %w", err)
		log.Debug(err.Error())
		return domain.User{}, err
	}

	return user, nil
}
