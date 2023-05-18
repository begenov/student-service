package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/begenov/test-task-backend/students-app/internal/models"
)

type StudentsStorage struct {
	db *sql.DB
}

func NewStudentsStorage(db *sql.DB) *StudentsStorage {
	return &StudentsStorage{
		db: db,
	}
}

func (s *StudentsStorage) CreateStudent(ctx context.Context, student models.Student) error {
	stmt := `INSERT INTO students (email, name, gpa, courses) VALUES ($1, $2, $3)`
	if _, err := s.db.ExecContext(ctx, stmt, student.Email, student.GPA, student.Courses); err != nil {
		log.Printf("Error executing SQL statement: %s", err.Error())
		return fmt.Errorf("failed to create student: %w", err)
	}

	return nil
}

func (s *StudentsStorage) GetStudentByID(ctx context.Context) (*models.Student, error) {
	return nil, nil
}

func (s *StudentsStorage) Update(ctx context.Context, studentID int) error {
	return nil
}

func (s *StudentsStorage) Delete(ctx context.Context, studentID int) error {
	return nil
}

func (s *StudentsStorage) ByIDCourses(ctx context.Context, studentID int) ([]models.Student, error) {
	return nil, nil
}
