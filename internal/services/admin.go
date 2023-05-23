package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/begenov/student-servcie/internal/config"
	"github.com/begenov/student-servcie/internal/models"
	"github.com/begenov/student-servcie/internal/storage"
	"github.com/begenov/student-servcie/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type AdminServcie struct {
	repos   storage.Admins
	manager auth.TokenManager
	cfg     *config.Config
}

func NewAdminService(repos storage.Admins, manager auth.TokenManager, cfg *config.Config) *AdminServcie {
	return &AdminServcie{
		repos:   repos,
		manager: manager,
		cfg:     cfg,
	}
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

	a, err := s.repos.AdminByEmail(ctx, admin.Email)

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(admin.Password)); err != nil {
		return nil, err
	}

	return s.createSession(ctx, a.ID)
}

func (s *AdminServcie) RefreshToken(ctx context.Context, access models.Token) (*models.Token, error) {
	if access.AccessToken == "" {
		return nil, errors.New("empty access token")
	}

	// adminID, err := s.manager.Parse(access.AccessToken)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (s *AdminServcie) createSession(ctx context.Context, adminID int) (*models.Token, error) {
	var (
		res *models.Token
		err error
	)

	res.AccessToken, err = s.manager.NewJWT(adminID, s.cfg.JWT.AccessTokenTTL)
	if err != nil {
		return nil, err
	}

	res.RefreshToken, err = s.manager.NewRefreshToken()

	if err != nil {
		return nil, err
	}

	session := models.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(s.cfg.JWT.RefreshTokenTTL),
	}

	if err := s.repos.RefreshTokenUpdate(ctx, session, adminID); err != nil {
		return nil, err
	}

	return res, nil
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
