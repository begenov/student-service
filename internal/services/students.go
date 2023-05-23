package services

import (
	"context"
	"log"

	m "github.com/begenov/test-task-backend/pkg/models"
	"github.com/begenov/test-task-backend/students-app/internal/models"
	"github.com/begenov/test-task-backend/students-app/internal/storage"
)

type StudentsService struct {
	storage storage.Students
}

func NewStudentsService(storage storage.Students) *StudentsService {
	return &StudentsService{
		storage: storage,
	}
}

func (s *StudentsService) CreateStudent(ctx context.Context, student models.Student) error {
	// нужно реолизовать validate
	if err := s.storage.CreateStudent(ctx, student); err != nil {
		log.Printf("Error creating student: %s", err.Error())
		return err
	}
	return nil
}

func (s *StudentsService) GetStudentByID(ctx context.Context, id int) (*models.Student, error) {
	student, err := s.storage.GetStudentByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentsService) Update(ctx context.Context, studentUpdate models.Student) error {

	student, err := s.storage.GetStudentByID(ctx, studentUpdate.ID)
	if err != nil {
		return err
	}
	// нужно реолизовать validate
	if studentUpdate.Email != "" {
		student.Email = studentUpdate.Email
	}

	if studentUpdate.Name != "" {
		student.Name = studentUpdate.Name
	}

	if studentUpdate.GPA != 0 {
		student.GPA = studentUpdate.GPA
	}

	if len(studentUpdate.Courses) > 0 {
		student.Courses = studentUpdate.Courses
	}
	if err = s.storage.Update(ctx, *student); err != nil {
		return err
	}
	return nil
}

func (s *StudentsService) Delete(ctx context.Context, studentID int) error {

	if err := s.storage.Delete(ctx, studentID); err != nil {
		return err
	}

	return nil
}

func (s *StudentsService) ByIDStudents(ctx context.Context, studentID int) ([]models.Student, error) {

	students, err := s.storage.ByIDCourses(ctx, studentID)

	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s *StudentsService) ByIDCourses(ctx context.Context, studentID int) ([]m.Course, error) {
	return nil, nil
}
