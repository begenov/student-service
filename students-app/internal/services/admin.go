package services

import (
	"context"

	"github.com/begenov/test-task-backend/pkg/auth"
	"github.com/begenov/test-task-backend/students-app/internal/models"
	"github.com/begenov/test-task-backend/students-app/internal/storage"
)

type AdminServcie struct {
	repos   storage.Admins
	manager auth.TokenManager
}

func NewAdminService() *AdminServcie {
	return &AdminServcie{}
}

func (s *AdminServcie) SignUpAdmin(ctx context.Context, admin models.Admin) error {

	return nil
}

func (s *AdminServcie) SignInAdmin(ctx context.Context, admin models.Admin) (*models.Token, error) {
	return nil, nil
}

func (s *AdminServcie) RefreshToken(ctx context.Context, access models.Token) (*models.Token, error) {
	return nil, nil
}
