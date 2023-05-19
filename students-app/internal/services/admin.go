package services

import (
	"context"
	"errors"
	"strings"

	"github.com/begenov/test-task-backend/pkg/auth"
	"github.com/begenov/test-task-backend/students-app/internal/models"
	"github.com/begenov/test-task-backend/students-app/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type AdminServcie struct {
	repos   storage.Admins
	manager auth.TokenManager
}

func NewAdminService() *AdminServcie {
	return &AdminServcie{}
}

func (s *AdminServcie) SignUpAdmin(ctx context.Context, admin models.Admin) error {
	if err := checkAdminInput(admin); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin.Password = string(hash)

	if err := s.repos.CreateAdmin(ctx, admin); err != nil {
		return err
	}

	return nil
}

func (s *AdminServcie) SignInAdmin(ctx context.Context, admin models.Admin) (*models.Token, error) {
	return nil, nil
}

func (s *AdminServcie) RefreshToken(ctx context.Context, access models.Token) (*models.Token, error) {
	return nil, nil
}

func checkAdminInput(admin models.Admin) error {
	if strings.TrimSpace(admin.Email) == "" {
		return errors.New("inccorect empty email")
	}

	if strings.TrimSpace(admin.Name) == "" {
		return errors.New("inccorect empty name")
	}

	if strings.TrimSpace(admin.Password) == "" {
		return errors.New("inccorect empty password")
	}

	return nil
}
