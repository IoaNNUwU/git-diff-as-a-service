package users_postgres_repository

type UserModel struct {
	ID      int
	Version int

	Role string

	FullName string
	Email    *string
}