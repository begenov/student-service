package postgresql

import (
	"context"
	"database/sql"

	"github.com/begenov/test-task-backend/students-app/internal/models"
)

type AdminsStorage struct {
	db *sql.DB
}

func NewAdminsStorage(db *sql.DB) *AdminsStorage {
	return &AdminsStorage{
		db: db,
	}
}

func (s *AdminsStorage) CreateAdmin(ctx context.Context, admin models.Admin) error {
	stmt := `INSERT INTO email, name, password_hash FROM admin VALUES ($1, $2, password_hash)`
	if _, err := s.db.ExecContext(ctx, stmt, admin.Email, admin.Name, admin.Password); err != nil {
		return err
	}
	return nil
}

func (s *AdminsStorage) AdminByEmail(ctx context.Context, admin *models.Admin) (*models.Admin, error) {
	stmt := `SELECT * FROM admin WHERE email = $1`

	if err := s.db.QueryRowContext(ctx, stmt, admin.Email).Scan(admin.ID, admin.Email, admin.Name, admin.Password); err != nil {
		return nil, err
	}
	return admin, nil
}

func (s *AdminsStorage) RefreshTokenUpdate(ctx context.Context, refresh string, addminID int) error {
	stmt := `UPDATE admin SET refresh_token = $1 WHERE id = $2`

	if _, err := s.db.ExecContext(ctx, stmt, refresh, addminID); err != nil {
		return err
	}

	return nil
}
