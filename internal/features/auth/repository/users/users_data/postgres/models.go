package auth_users_users_data_postgres_repository

type userDataModel struct {
	ID      int
	Version int

	Role string

	FullName string
	Email    *string
}
