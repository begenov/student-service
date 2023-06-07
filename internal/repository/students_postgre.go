package repository

import (
	"context"
	"database/sql"
	"log"

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
	stmt := `INSERT INTO student(email, name, password_hash, gpa, courses) VALUES($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, stmt, student.Email, student.Name, student.Password, student.GPA, pq.Array(student.Courses))
	if err != nil {
		return err
	}
	return nil
}

func (r *StudentsRepo) GetByEmail(ctx context.Context, email string) (domain.Student, error) {
	var stud domain.Student
	stmt := `SELECT id, email, name, password_hash, gpa, courses FROM student WHERE email = $1`

	if err := r.db.QueryRowContext(ctx, stmt, email).Scan(&stud.ID, &stud.Email, &stud.Name, &stud.Password, &stud.GPA, pq.Array(&stud.Courses)); err != nil {
		log.Printf("error %s", err)
		return domain.Student{}, err
	}
	return stud, nil
}

func (r *StudentsRepo) GetByID(ctx context.Context, id int) (domain.Student, error) {
	var student domain.Student

	stmt := `SELECT id, email, name, gpa, courses, password_hash, created_at FROM student WHERE id = $1`
	var courses pq.StringArray
	err := r.db.QueryRowContext(ctx, stmt, id).Scan(&student.ID, &student.Email, &student.Name, &student.GPA, &courses, &student.Password, &student.ExpiresAt)
	if err != nil {
		return student, err
	}
	student.Courses = []string(courses)
	return student, nil
}

func (r *StudentsRepo) Update(ctx context.Context, student domain.Student) error {
	stmt := `UPDATE student SET email = $1, name = $2, gpa = $3, courses = $4, password_hash = $5 WHERE id = $6`
	_, err := r.db.ExecContext(ctx, stmt, student.Email, student.Name, student.GPA, pq.Array(student.Courses), student.Password, student.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *StudentsRepo) Delete(ctx context.Context, id int) error {

	stmt, err := r.db.PrepareContext(ctx, "DELETE FROM student WHERE id = $1")
	if err != nil {
		return err
	}
	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *StudentsRepo) GetStudentsByCoursesID(ctx context.Context, coursesID string) ([]domain.Student, error) {
	log.Println(coursesID)
	var students []domain.Student
	stmt := `SELECT id, email, name, gpa, courses FROM student WHERE $1 = ANY(courses)`
	rows, err := r.db.QueryContext(ctx, stmt, coursesID)
	if err != nil {
		return students, err
	}
	defer rows.Close()
	for rows.Next() {
		var student domain.Student
		err := rows.Scan(&student.ID, &student.Email, &student.Name, &student.GPA, pq.Array(&student.Courses))
		if err != nil {
			return students, err
		}

		students = append(students, student)
	}
	return students, nil
}
func (r *StudentsRepo) SetSession(ctx context.Context, session domain.Session, id int) error {
	stmt := `UPDATE student SET refresh_token = $1, created_at = $2 WHERE id = $3`
	if _, err := r.db.ExecContext(ctx, stmt, session.RefreshToken, session.ExpiresAt, id); err != nil {
		return err
	}
	return nil
}

func (r *StudentsRepo) GetByRefresh(ctx context.Context, refreshToken string) (domain.Student, error) {
	stmt := `SELECT id, email, password_hash, name, gpa, refresh_token, created_at, courses FROM student WHERE refresh_token = $1`

	var student domain.Student

	err := r.db.QueryRowContext(ctx, stmt, refreshToken).Scan(&student.ID, &student.Email, &student.Password, &student.Name, &student.GPA, &student.RefreshToken, &student.ExpiresAt, pq.Array(&student.Courses))
	if err != nil {
		log.Printf("error refresh student: %s", err)
		return student, err
	}
	return student, nil
}
