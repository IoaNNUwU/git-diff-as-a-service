package files_service

import (
	"context"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
)

type FilesService struct {
	filesRepository FilesRepository
}

func NewFilesService(usersRepository FilesRepository) FilesService {
	return FilesService{
		filesRepository: usersRepository,
	}
}

type FilesRepository interface {
	CreateFile(ctx context.Context, file domain.File) (domain.File, error)
	DeleteFile(ctx context.Context, id int) error
}
