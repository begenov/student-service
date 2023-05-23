package services

import (
	"context"

	"github.com/begenov/student-servcie/internal/config"
	"github.com/begenov/student-servcie/internal/models"
	"github.com/begenov/student-servcie/internal/storage"
	"github.com/begenov/student-servcie/pkg/auth"
	m "github.com/begenov/student-servcie/pkg/models"
)

type Students interface {
	CreateStudent(ctx context.Context, student models.Student) error
	GetStudentByID(ctx context.Context, id int) (*models.Student, error)
	Update(ctx context.Context, student models.Student) error
	Delete(ctx context.Context, studentID int) error
	ByIDCourses(ctx context.Context, studentID int) ([]m.Course, error)
	ByIDStudents(ctx context.Context, studentID int) ([]models.Student, error)
}

type Admin interface {
	SignUpAdmin(ctx context.Context, admin models.Admin) error
	SignInAdmin(ctx context.Context, admin models.Admin) (*models.Token, error)
	RefreshToken(ctx context.Context, access models.Token) (*models.Token, error)
}

type Services struct {
	Students
	Admin
}

func NewService(storage *storage.Storage, tokenManager auth.TokenManager, cfg *config.Config) *Services {
	return &Services{
		Students: NewStudentsService(storage.Students),
		Admin:    NewAdminService(storage.Admins, tokenManager, cfg),
	}
}
