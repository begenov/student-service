package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/begenov/student-service/internal/domain"
	"github.com/begenov/student-service/pkg/auth"
)

func TestAdminsCreate(t *testing.T) {
	if _, err := createAdmin(context.Background()); err != nil {
		t.Error("error create admin", err)
	}
}

func TestAdminGetByEmail(t *testing.T) {
	admin, err := createAdmin(context.Background())
	if err != nil {
		t.Error("error get by email", err)
	}
	a, err := repo.Admins.GetByEmail(context.Background(), admin.Email)
	if err != nil {
		t.Error("error get by email", err)
	}

	if a.Email != admin.Email {
		t.Error("empty email")
	}

	if a.Name != admin.Name {
		t.Error("empty name")
	}

	if a.Password != admin.Password {
		t.Error("empty password")
	}
}

func TestAdminSetSession(t *testing.T) {
	_, err := updateSetSession()
	if err != nil {
		t.Error("set session error", err)
	}
}

func TestAdminGetByRefresh(t *testing.T) {
	refresh, err := updateSetSession()
	if err != nil {
		t.Error("get by refresh token", err)
	}
	admin, err := repo.Admins.GetByRefresh(context.Background(), refresh)
	if err != nil {
		t.Error("get by refresh token", err)
	}
	if admin.Email == "" {
		t.Error("empty email")
	}
	if admin.Name == "" {
		t.Error("empty name")
	}

	if admin.Password == "" {
		t.Error("empty password")
	}

	if admin.RefreshToken == "" {
		t.Error("empty refresh")
	}

}

func createAdmin(ctx context.Context) (domain.Admin, error) {
	admin := domain.Admin{
		ID:       1,
		Email:    "test@test.com",
		Name:     "test",
		Password: "test",
	}
	err := repo.Admins.Create(ctx, admin)
	if err != nil {
		return admin, err
	}
	return admin, nil
}
func updateSetSession() (string, error) {
	manager, err := auth.NewManager(cfg.JWT.SigningKey)
	if err != nil {
		return "", err
	}
	refreshtoken, err := manager.NewRefreshToken()
	if err != nil {
		return "", err
	}
	err = repo.Admins.SetSession(context.Background(), domain.Session{
		RefreshToken: refreshtoken,
		ExpiresAt:    time.Now().Add(cfg.JWT.RefreshTokenTTL),
	}, 1)
	if err != nil {
		return "", err
	}
	return refreshtoken, nil
}
