package services

import (
	"context"
	"log"

	"github.com/begenov/test-task-backend/pkg/auth"
	"github.com/begenov/test-task-backend/students-app/internal/models"
	"github.com/begenov/test-task-backend/students-app/internal/storage"
)

type StudentsService struct {
	storage storage.Students
	manager auth.TokenManager
}

func NewStudentsService(storage storage.Students, tokenManager auth.TokenManager) *StudentsService {
	return &StudentsService{
		storage: storage,
		manager: tokenManager,
	}
}

func (s *StudentsService) CreateStudent(ctx context.Context, student models.Student) error {
	if err := s.storage.CreateStudent(ctx, student); err != nil {
		log.Printf("Error creating student: %s", err.Error())
		return err
	}
	return nil
}

func (s *StudentsService) GetStudentByID(ctx context.Context) (*models.Student, error) {
	return nil, nil
}

func (s *StudentsService) Update(ctx context.Context, studentID int) error {
	return nil
}

func (s *StudentsService) Delete(ctx context.Context, studentID int) error {
	return nil
}

func (s *StudentsService) ByIDCourses(ctx context.Context, studentID int) ([]models.Student, error) {
	return nil, nil
}
