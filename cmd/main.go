package main

import (
	"log"

	"github.com/begenov/student-service/internal/config"
	"github.com/begenov/student-service/internal/repository"
	"github.com/begenov/student-service/pkg/database"
)

const (
	path = "./.env"
)

func main() {
	cfg, err := config.Init(path)
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.OpenDB(cfg.Database.Driver, cfg.Database.DSN)

	if err != nil {
		log.Fatalf("error creating database object: %v", err)
	}

	// tokenManager, err := auth.NewManager(cfg.JWT.SigningKey)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	repos := repository.NewRepository(db)

}
