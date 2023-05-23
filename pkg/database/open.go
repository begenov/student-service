package database

import (
	"context"
	"database/sql"
)

func OpenDB(driver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(context.Background()); err != nil {
		return nil, err
	}
	return db, nil
}
