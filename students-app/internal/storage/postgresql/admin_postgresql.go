package postgresql

import (
	"context"

	"github.com/begenov/test-task-backend/students-app/internal/models"
)

type AdminsStorage struct {
}

func NewAdminsStorage() *AdminsStorage {
	return &AdminsStorage{}
}

func (s *AdminsStorage) CreateAdmin(ctx context.Context, admin models.Admin) error {
	return nil
}

func (s *AdminsStorage) AdminByEmail(ctx context.Context, admin models.Admin) (*models.Token, error) {
	return nil, nil
}
