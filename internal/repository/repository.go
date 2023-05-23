package repository

import (
	"context"
	"database/sql"

	"github.com/begenov/student-service/internal/domain"
)

type Students interface {
	Create(ctx context.Context, student domain.Student) error
	GetByID(ctx context.Context, id int) (domain.Student, error)
	Update(ctx context.Context, student domain.Student) error
	Delete(ctx context.Context, id int) error
}

type Admins interface {
	Create(ctx context.Context, admin domain.Admin) error
	GetByEmail(ctx context.Context, email string) (domain.Admin, error)
	SetSession(ctx context.Context, session domain.Session, id int) error
	GetByRefresh(ctx context.Context, refreshToken string) (domain.Admin, error)
}

type Repository struct {
	Students Students
	Admins   Admins
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Students: NewStudentsRepo(db),
		Admins:   NewAdminRepo(db),
	}
}
