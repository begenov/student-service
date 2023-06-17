package repository_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/begenov/student-service/internal/config"
	"github.com/begenov/student-service/internal/repository"

	_ "github.com/lib/pq"
)

const path = ".env"

var db *sql.DB
var repo *repository.Repository
var cfg *config.Config

func TestMain(m *testing.M) {
	var err error
	cfg, err = config.Init(path)
	if err != nil {
		log.Fatalf("connot load config: %v", err)
	}
	db, err = sql.Open(cfg.Database.Driver, cfg.Database.DSN)
	if err != nil {
		log.Fatal("connot connect to db:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("connot ping to db:", err)
	}

	repo = repository.NewRepository(db)
	os.Exit(m.Run())
}
