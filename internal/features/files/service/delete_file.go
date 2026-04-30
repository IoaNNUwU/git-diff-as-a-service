package files_service

import (
	"context"
	"fmt"

	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
)

func (s *FilesService) DeleteFile(ctx context.Context, id int) error {

	log := FilesServiceLogger(ctx)

	if id < 0 {
		return fmt.Errorf("file id cannot be negative: %w", core_errors.ErrInvalidArgument)
	}

	if err := s.filesRepository.DeleteFile(ctx, id); err != nil {
		err := fmt.Errorf("unable to delete user: %w", err)
		log.Debug(err.Error())
		return err
	}

	return nil
}
