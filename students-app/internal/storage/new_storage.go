package storage

import (
	"database/sql"

	"github.com/begenov/test-task-backend/students-app/internal/storage/postgresql"
)

type Students interface {
}

type Storage struct {
	Students Students
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Students: postgresql.NewStudentsStorage(db),
	}
}
