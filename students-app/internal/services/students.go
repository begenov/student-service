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

func (s *StudentsService) studentSignUp(ctx context.Context) {
}

func (s *StudentsService) studentSingIn(ctx context.Context) {
}

func (s *StudentsService) studentUpdate(ctx context.Context) {
}

func (s *StudentsService) studentDelete(ctx context.Context) {
}

func (s *StudentsService) studentByIDCourses(ctx context.Context) {

}
