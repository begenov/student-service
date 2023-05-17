package storage

import (
	"context"
	"database/sql"

	"github.com/begenov/test-task-backend/students-app/internal/models"
	"github.com/begenov/test-task-backend/students-app/internal/storage/postgresql"
)

type Students interface {
	CreateStudent(ctx context.Context, student models.Student) error
	GetStudentByID(ctx context.Context) (*models.Student, error)
	Update(ctx context.Context, studentID int) error
	Delete(ctx context.Context, studentID int) error
	ByIDCourses(ctx context.Context, studentID int) ([]models.Student, error)
}

type Storage struct {
	Students Students
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Students: postgresql.NewStudentsStorage(db),
	}
}
