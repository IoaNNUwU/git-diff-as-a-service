package files_service

import (
	"context"
	"fmt"

	core_auth "github.com/ioannuwu/git-diff-as-a-service/internal/core/auth"
	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
)

func (s *FilesService) DeleteFile(ctx context.Context, id int) error {

	log := FilesServiceLogger(ctx)

	if id < 0 {
		return fmt.Errorf("file id cannot be negative: %w", core_errors.ErrInvalidArgument)
	}

	userID := core_auth.UserIDFromContext(ctx)
	userRole := core_auth.RoleFromContext(ctx)

	var errMsg string
	var err error
	if userRole == core_auth.RoleAdmin {
		err = s.filesRepository.DeleteFile(ctx, id)
		errMsg = "unable to delete file"
	} else {
		err = s.filesRepository.DeleteFileByAuthor(ctx, id, userID)
		errMsg = "unable to delete file owned by another user"
	}
	if err != nil {
		err := fmt.Errorf("%s: %w", errMsg, err)
		log.Debug(err.Error())
		return err
	}

	if err := s.filesRepository.DeleteFile(ctx, id); err != nil {
		err := fmt.Errorf("unable to delete user: %w", err)
		log.Debug(err.Error())
		return err
	}

	return nil
}
