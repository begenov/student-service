package postgresql

import "database/sql"

type StudentsStorage struct {
	db *sql.DB
}

func NewStudentsStorage(db *sql.DB) *StudentsStorage {
	return &StudentsStorage{
		db: db,
	}
}
