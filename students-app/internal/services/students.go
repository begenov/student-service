package services

import (
	"context"

	"github.com/begenov/test-task-backend/pkg/auth"
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

func (s *StudentsService) SignUp(ctx context.Context) {
}

func (s *StudentsService) SingIn(ctx context.Context) {
}

func (s *StudentsService) Update(ctx context.Context, studentID int) {
}

func (s *StudentsService) Delete(ctx context.Context, studentID int) {
}

func (s *StudentsService) ByIDCourses(ctx context.Context, studentID int) {
}
