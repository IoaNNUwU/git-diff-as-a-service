package domain

import (
	"fmt"

	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
)

type File struct {
	ID      int
	Version int

	FileName string
	OwnerID  int
	Content  string
}

func NewFile(id int, version int, fileName string, ownerID int, content string) File {
	return File{
		ID:       id,
		Version:  version,
		FileName: fileName,
		Content:  content,
		OwnerID: ownerID,
	}
}

func NewFileUninitialized(fileName string, ownerID int, content string) File {
	return NewFile(UninitializedID, UninitializedVersion, fileName, ownerID, content)
}

func (f *File) Validate() error {

	if f.ID == UninitializedID || f.Version == UninitializedVersion {
		return fmt.Errorf("file wasn't properly initialized")
	}

	fullNameLength := len([]rune(f.FileName))

	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf(
			"invalid `file_name` length: %d: %w", fullNameLength, core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
