package services

import (
	"context"

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
}

func (s *StudentsService) GetStudentByID(ctx context.Context) (*models.Student, error) {
}

func (s *StudentsService) Update(ctx context.Context, studentID int) error {
}

func (s *StudentsService) Delete(ctx context.Context, studentID int) error {
}

func (s *StudentsService) ByIDCourses(ctx context.Context, studentID int) ([]Students, error) {
}
