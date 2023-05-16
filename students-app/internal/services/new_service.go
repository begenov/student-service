package services

import "github.com/begenov/test-task-backend/students-app/internal/storage"

type Students interface {
}

type Services struct {
	Students Students
}

func NewService(storage *storage.Storage) *Services {
	return &Services{
		Students: NewStudentsService(storage.Students),
	}
}
