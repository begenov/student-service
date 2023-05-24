package service

import (
	"context"
	"errors"

	"github.com/begenov/student-service/internal/domain"
	"github.com/begenov/student-service/internal/repository"
)

type StudentService struct {
	repo repository.Students
}

func NewStudentService(repo repository.Students) *StudentService {
	return &StudentService{
		repo: repo,
	}
}

func (s *StudentService) Create(ctx context.Context, student domain.Student) error {
	return s.repo.Create(ctx, student)

}

func (s *StudentService) GetStudentByID(ctx context.Context, id int) (domain.Student, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *StudentService) Update(ctx context.Context, student domain.Student) error {
	stud, err := s.repo.GetByID(ctx, student.ID)
	if err != nil {
		return err
	}

	if stud.ID != student.ID {
		return errors.New("ID mismatch: retrieved student ID does not match the expected ID")
	}

	if len(student.Courses) == 0 {
		student.Courses = stud.Courses
	}

	if student.Email == "" {
		student.Email = stud.Email
	}

	if student.Name == "" {
		student.Name = stud.Name
	}

	if student.GPA == 0 {
		student.GPA = stud.GPA
	}

	return s.repo.Update(ctx, student)

}

func (s *StudentService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *StudentService) GetStudentsByCoursesID(ctx context.Context, id string) ([]domain.Student, error) {
	return s.repo.GetStudentsByCoursesID(ctx, id)
}
