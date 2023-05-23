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

func (s *StudentsStorage) GetStudentByID(ctx context.Context, id int) (student *models.Student, err error) {
	stmt := `SELECT * FROM students WHERE id = $1`
	if err := s.db.QueryRowContext(ctx, stmt, id).Scan(&student.ID, student.Email, student.Name, student.GPA, student.Courses); err != nil {
		return nil, err
	}

	return student, nil
}

func (s *StudentsStorage) Update(ctx context.Context, student models.Student) error {

	stmt := `UPDATE email, name, gpa, courses  FROM students VALUES ($1, $2, $3, $4)`
	if _, err := s.db.ExecContext(ctx, stmt, student.Email, student.Name, student.GPA, student.Courses); err != nil {
		return err
	}

	return nil
}

func (s *StudentsStorage) Delete(ctx context.Context, studentID int) error {
	return nil
}

func (s *StudentsStorage) ByIDCourses(ctx context.Context, studentID int) ([]models.Student, error) {
	return nil, nil
}
