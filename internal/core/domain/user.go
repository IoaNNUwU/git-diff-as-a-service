package domain

import (
	"fmt"

	core_auth "github.com/ioannuwu/git-diff-as-a-service/internal/core/auth"
	core_errors "github.com/ioannuwu/git-diff-as-a-service/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	Role string // this field can only be changed inside database itself

	FullName string
	Email    *string
}

func NewUser(id int, version int, role string, fullName string, email *string) User {
	return User{
		ID:       id,
		Version:  version,
		Role:     role,
		FullName: fullName,
		Email:    email,
	}
}

func NewUserUninitialized(fullName string, email *string) User {
	return NewUser(UninitializedID, UninitializedVersion, core_auth.RoleUser, fullName, email)
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
