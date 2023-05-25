package service

import (
	"context"
	"strconv"
	"time"

	"github.com/begenov/student-service/internal/domain"
	"github.com/begenov/student-service/internal/repository"
	"github.com/begenov/student-service/pkg/auth"
	"github.com/begenov/student-service/pkg/hash"
)

type AdminService struct {
	hash            hash.PasswordHasher
	repo            repository.Admins
	manager         auth.TokenManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAdminService(repo repository.Admins, hash hash.PasswordHasher, manager auth.TokenManager, accessTokenTTL time.Duration, refreshTokenTTL time.Duration) *AdminService {
	return &AdminService{
		repo:            repo,
		hash:            hash,
		manager:         manager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (s *AdminService) SignUp(ctx context.Context, admin domain.Admin) error {
	var err error
	admin.Password, err = s.hash.GenerateFromPassword(admin.Password)
	if err != nil {
		return err
	}
	return s.repo.Create(ctx, admin)
}

func (s *AdminService) SignIn(ctx context.Context, email string, password string) (domain.Token, error) {
	var err error
	admin, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return domain.Token{}, domain.ErrNotFound
	}

	if err = s.hash.CompareHashAndPassword(admin.Password, password); err != nil {
		return domain.Token{}, domain.ErrNotFound
	}

	return s.createSession(ctx, admin.ID)
}

func (s *AdminService) GetByRefreshToken(ctx context.Context, refreshToken string) (domain.Token, error) {
	student, err := s.repo.GetByRefresh(ctx, refreshToken)
	if err != nil {
		return domain.Token{}, err
	}

	return s.createSession(ctx, student.ID)
}

func (s *AdminService) createSession(ctx context.Context, adminID int) (domain.Token, error) {
	var (
		res domain.Token
		err error
	)
	res.AccessToken, err = s.manager.NewJWT(strconv.Itoa(adminID), s.accessTokenTTL)
	if err != nil {
		return res, err
	}
	res.RefreshToken, err = s.manager.NewRefreshToken()
	if err != nil {
		return res, err
	}
	session := domain.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	}
	err = s.repo.SetSession(ctx, session, adminID)

	return res, err
}
