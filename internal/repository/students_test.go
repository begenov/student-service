package repository_test

import (
	"context"
	"testing"

	"github.com/begenov/student-service/internal/domain"
)

func TestStudentCreate(t *testing.T) {
	student := domain.Student{
		Email:   "test@test.com",
		Name:    "test",
		GPA:     3.5,
		Courses: []string{"1", "2", "3"},
	}
	_, err := createStudent(context.Background(), student)
	if err != nil {
		t.Errorf("error testing %v", err)
	}

}

func TestStudentGetByID(t *testing.T) {
	student := domain.Student{
		Email:   "test@test.com",
		Name:    "test",
		GPA:     3.5,
		Courses: []string{"1", "2", "3"},
	}
	student, err := createStudent(context.Background(), student)
	if err != nil {
		t.Errorf("error testing %v", err)
	}

	s, err := repo.Students.GetByID(context.Background(), 2)
	if err != nil {
		t.Errorf("get by id student: %v", err)
	}
	if s.Email != student.Email {
		t.Error("error email")
	}
	if s.Name != student.Name {
		t.Error("error name")
	}

	if s.GPA != student.GPA {
		t.Error("error GPA")
	}

	if len(s.Courses) == 0 {
		t.Error("error courses")
	}

}

func TestStudentUpdate(t *testing.T) {
	student := domain.Student{
		Email:   "test@test.com",
		Name:    "test",
		GPA:     3.5,
		Courses: []string{"1", "2", "3"},
		ID:      1,
	}
	student, err := createStudent(context.Background(), student)
	if err != nil {
		t.Errorf("error testing %v", err)
	}

	updateInput := domain.UpdateStudentInput{
		Email:   "new@new.com",
		Name:    "new@new.com",
		GPA:     2.0,
		Courses: []string{"1", "2"},
	}

	if err := repo.Students.Update(context.Background(), domain.Student{
		Email:   updateInput.Email,
		Name:    updateInput.Name,
		GPA:     updateInput.GPA,
		Courses: updateInput.Courses,
		ID:      student.ID,
	}); err != nil {
		t.Error("error update", err)
	}

}

func TestStudentDelete(t *testing.T) {
	student := domain.Student{
		Email:   "test@test.com",
		Name:    "test",
		GPA:     3.5,
		Courses: []string{"1", "2", "3"},
		ID:      1,
	}

	student, err := createStudent(context.Background(), student)
	if err != nil {
		t.Errorf("error testing %v", err)
	}
	if err := repo.Students.Delete(context.Background(), student.ID); err != nil {
		t.Error("error delete", err)
	}
}

func TestStudentGetStudentsByCoursesID(t *testing.T) {
	student := domain.Student{
		Email:   "test@test.com",
		Name:    "test",
		GPA:     3.5,
		Courses: []string{"1", "2", "3"},
		ID:      1,
	}
	_, err := createStudent(context.Background(), student)

	if err != nil {
		t.Error("error testing", err)
	}
	students, err := repo.Students.GetStudentsByCoursesID(context.Background(), "2")
	if err != nil {
		t.Error("error get student courses id", err)
	}
	if len(students) == 0 {
		t.Error("empty students", students)
	}
}

func createStudent(ctx context.Context, student domain.Student) (domain.Student, error) {
	if err := repo.Students.Create(context.Background(), student); err != nil {
		return domain.Student{}, nil
	}
	return student, nil
}
