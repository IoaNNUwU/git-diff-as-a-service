package files_service

import (
	"context"
	"fmt"

	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
)

func (s *FilesService) CreateFile(ctx context.Context, file domain.File) (domain.File, error) {

	log := FilesServiceLogger(ctx)

	if err := file.Validate(); err != nil {
		err := fmt.Errorf("invalid file: %w", err)
		log.Debug(err.Error())
		return domain.File{}, err
	}

	user, err := s.filesRepository.CreateFile(ctx, file)
	if err != nil {
		err := fmt.Errorf("unable to create file: %w", err)
		log.Debug(err.Error())
		return domain.File{}, err
	}

	return user, nil
}
