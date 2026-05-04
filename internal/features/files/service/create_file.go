package files_service

import (
	"context"
	"fmt"

	core_auth "github.com/ioannuwu/git-diff-as-a-service/internal/core/auth"
	"github.com/ioannuwu/git-diff-as-a-service/internal/core/domain"
	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
)

func (s *FilesService) CreateFile(ctx context.Context, file domain.File) (domain.File, error) {

	log := FilesServiceLogger(ctx)

	userID := core_auth.UserIDFromContext(ctx)
	if userID != file.OwnerID {
		err := fmt.Errorf("unable to add file owned by another user: %w", core_errors.ErrPermission)
		log.Debug(err.Error())
		return domain.File{}, err
	}

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
