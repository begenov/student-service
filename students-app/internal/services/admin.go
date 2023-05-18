package services

import (
	"github.com/begenov/test-task-backend/pkg/auth"
	"github.com/begenov/test-task-backend/students-app/internal/storage"
)

type AdminServcie struct {
	repos   storage.Storage
	manager auth.TokenManager
}

func NewAdminService() *AdminServcie {
	return &AdminServcie{}
}
