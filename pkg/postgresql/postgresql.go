package postgresql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Client interface {
}

func NewPostgreSQLDB(driver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
