package main

import (
	"log"

	"github.com/begenov/test-task-backend/pkg/postgresql"
	"github.com/begenov/test-task-backend/students-app/internal/config"
	"github.com/begenov/test-task-backend/students-app/internal/storage"
)

const (
	path_config = ""
)

func main() {
	cfg, err := config.NewConfig(path_config)
	if err != nil {
		log.Fatalf("can't load config: %v", err)
		return
	}

	db, err := postgresql.NewPostgreSQLDB(cfg.Database.Driver, cfg.Database.DSN)
	if err != nil {
		log.Fatalf("error creating database object: %v", err)
		return
	}

	storage := storage.NewStorage()

}
