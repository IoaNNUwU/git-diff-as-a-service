package files_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
)

func (r *FilesRepository) CreateFile(ctx context.Context, file domain.File) (domain.File, error) {

	log := FilesRepositoryPostgresLogger(ctx)

	ctx, cancel := context.WithTimeout(ctx, r.pool.Timeout())
	defer cancel()

	println(file.OwnerID)

	query := `
	INSERT INTO git_diff_app.files (name, created_at, owner_id, expiration, content)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, version, name, content;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		file.FileName,
		time.Now(),
		file.OwnerID,
		time.Now().Add(7 * time.Hour),
		file.Content,
	)

	var fileModel FileModel
	err := row.Scan(
		&fileModel.ID,
		&fileModel.Version,
		&fileModel.FileName,
		&fileModel.Content,
	)
	if err != nil {
		err := fmt.Errorf("scan error: %w", err)
		log.Debug(err.Error())
		return domain.File{}, core_errors.ErrTimeout
	}

	fileDomain := domain.NewFile(
		fileModel.ID,
		fileModel.Version,
		fileModel.FileName,
		fileModel.AuthorID,
		fileModel.Content,
	)

	return fileDomain, nil
}
