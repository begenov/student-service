package service

import (
	"context"

	"github.com/begenov/student-service/pkg/auth"
	"github.com/begenov/student-service/pkg/cache"
	"github.com/begenov/student-service/pkg/hash"
	"github.com/begenov/student-service/pkg/kafka"

	"github.com/begenov/student-service/internal/config"
	"github.com/begenov/student-service/internal/domain"
	"github.com/begenov/student-service/internal/repository"
)

//go:generate mockgen -source=services.go -destination=mocks/mock.go
type Students interface {
	Create(ctx context.Context, student domain.Student) error
	GetStudentByID(ctx context.Context, id int) (domain.Student, error)
	Update(ctx context.Context, student domain.Student) error
	Delete(ctx context.Context, id int) error
	GetStudentsByCoursesID(ctx context.Context, id string) ([]domain.Student, error)
	GetByEmail(ctx context.Context, email string, password string) (domain.Token, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (domain.Token, error)
}

type Admins interface {
	SignUp(ctx context.Context, admin domain.Admin) error
	SignIn(ctx context.Context, email string, password string) (domain.Token, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (domain.Token, error)
}

type Kafka interface {
	Read(ctx context.Context)
	Close()
}

type Service struct {
	Students Students
	Admins   Admins
	Kafka    Kafka
}

func NewService(repo *repository.Repository, hash hash.PasswordHasher, tokenManager auth.TokenManager, cache cache.Cache, cfg *config.Config, producer *kafka.Producer, consumer *kafka.Consumer) *Service {
	return &Service{
		Students: NewStudentService(repo.Students, hash, tokenManager, cache, cfg.JWT.AccessTokenTTL, cfg.JWT.RefreshTokenTTL),
		Admins:   NewAdminService(repo.Admins, hash, tokenManager, cache, cfg.JWT.AccessTokenTTL, cfg.JWT.RefreshTokenTTL),
		Kafka:    NewKafkaSerivce(repo.Students, producer, consumer),
	}
}
