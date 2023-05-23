package repository

import (
	"context"
	"database/sql"

	"github.com/begenov/student-service/internal/domain"
)

type AdminsRepo struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) *AdminsRepo {
	return &AdminsRepo{
		db: db,
	}
}

func (r *AdminsRepo) Create(ctx context.Context, admin domain.Admin) error {
	stmt := `INSERT INTO admin(email, name, password_hash) VALUES($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, stmt, admin.Email, admin.Name, admin.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *AdminsRepo) GetByEmail(ctx context.Context, email string) (domain.Admin, error) {
	stmt := `SELECT id, email, name, password_hash FROM admin WHERE email = $1`
	var admin domain.Admin
	if err := r.db.QueryRowContext(ctx, stmt, email).Scan(&admin.ID, &admin.Email, &admin.Name, &admin.Password); err != nil {
		return admin, err
	}
	return admin, nil
}

func (r *AdminsRepo) SetSession(ctx context.Context, session domain.Session, id int) error {
	stmt := `UPDATE admin SET refresh_token = $1, created_at = $2 WHERE id = $3`
	if _, err := r.db.ExecContext(ctx, stmt, session.RefreshToken, session.ExpiresAt, id); err != nil {
		return err
	}
	return nil
}

func (r *AdminsRepo) GetByRefresh(ctx context.Context, refreshToken string) (domain.Admin, error) {
	stmt := `SELECT id, email, name, refresh_token, created_at, password_hash FROM admin WHERE refresh_token = $1`

	var admin domain.Admin

	err := r.db.QueryRowContext(ctx, stmt, refreshToken).Scan(&admin.ID, &admin.Email, &admin.Name, &admin.RefreshToken, &admin.ExpiresAt, &admin.Password)
	if err != nil {
		return admin, err
	}
	return admin, nil
}
