package files_postgres_repository

type FileModel struct {
	ID      int
	Version int

	AuthorID int

	FileName string
	Content  string
}
