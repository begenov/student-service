package services

import (
	"context"

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
