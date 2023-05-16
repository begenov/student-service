package services

import "github.com/begenov/test-task-backend/students-app/internal/storage"

type StudentsService struct {
	storage storage.Students
}

func NewStudentsService(storage storage.Students) *StudentsService {
	return &StudentsService{
		storage: storage,
	}
}
