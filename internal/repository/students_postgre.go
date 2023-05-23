package repository

import (
	"context"
	"database/sql"

	"github.com/begenov/student-service/internal/domain"
	"github.com/lib/pq"
)

type StudentsRepo struct {
	db *sql.DB
}

func NewStudentsRepo(db *sql.DB) *StudentsRepo {
	return &StudentsRepo{
		db: db,
	}
}

func (r *StudentsRepo) Create(ctx context.Context, student domain.Student) error {
	stmt := `INSERT INTO student(email, name, gpa, courses) VALUES($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, stmt, student.Email, student.Name, student.GPA, pq.Array(student.Courses))
	if err != nil {
		return err
	}
	return nil
}

func (r *StudentsRepo) GetByID(ctx context.Context, id int) (domain.Student, error) {
	var student domain.Student

	stmt := `SELECT id, email, name, gpa, courses FROM student WHERE id = $1`
	var courses pq.StringArray
	err := r.db.QueryRowContext(ctx, stmt, id).Scan(&student.ID, &student.Email, &student.Name, &student.GPA, &courses)
	if err != nil {
		return student, err
	}
	student.Courses = []string(courses)
	return student, nil
}

func (r *StudentsRepo) Update(ctx context.Context, student domain.Student) error {
	stmt := `UPDATE student SET email = $1, name = $2, gpa = $3, courses = $4 WHERE id = $5`
	_, err := r.db.ExecContext(ctx, stmt, student.Email, student.Name, student.GPA, pq.Array(student.Courses), student.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *StudentsRepo) Delete(ctx context.Context, id int) error {
	stmt := `DELETE FROM student WHERE id = $1`
	_, err := r.db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *StudentsRepo) GetStudentsByCoursesID(ctx context.Context, coursesID string) ([]domain.Student, error) {
	var students []domain.Student
	stmt := `SELECT id, email, name, gpa, courses FROM student WHERE $1 = ANY(courses)`
	rows, err := r.db.QueryContext(ctx, stmt, coursesID)
	if err != nil {
		return students, err
	}
	defer rows.Close()

	for rows.Next() {
		var student domain.Student
		err := rows.Scan(&student.ID, &student.Email, &student.Name, &student.GPA, pq.Array(student.Courses))
		if err != nil {
			return students, err
		}
		students = append(students, student)
	}
	return students, nil
}
