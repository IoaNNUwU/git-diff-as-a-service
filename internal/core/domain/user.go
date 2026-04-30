package domain

import (
	"fmt"

	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	FullName string
	Email    *string
}

func NewUser(id int, version int, fullName string, email *string) User {
	return User{
		ID:       id,
		Version:  version,
		FullName: fullName,
		Email:    email,
	}
}

func NewUserUninitialized(fullName string, email *string) User {
	return NewUser(UninitializedID, UninitializedVersion, fullName, email)
}

func (u *User) Validate() error {
	
	if u.ID == UninitializedID || u.Version == UninitializedVersion {
		return fmt.Errorf("user wasn't properly initialized")
	}

	fullNameLength := len([]rune(u.FullName))

	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf(
			"invalid `full_name` length: %d: %w", fullNameLength, core_errors.ErrInvalidArgument,
		)
	}

	if u.Email != nil {
		emailLen := len([]rune(*u.Email))
		if emailLen < 10 || emailLen > 20 {
			return fmt.Errorf(
				"invalid `email` length: %d: %w", emailLen, core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}