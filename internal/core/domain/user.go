package domain

type User struct {
	ID 		int
	Version int

	FullName string
	Email	 *string
}