package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/begenov/student-service/internal/config"
	delivery "github.com/begenov/student-service/internal/delivery/http"
	"github.com/begenov/student-service/internal/repository"
	"github.com/begenov/student-service/internal/server"
	"github.com/begenov/student-service/internal/service"
	"github.com/begenov/student-service/pkg/auth"
	"github.com/begenov/student-service/pkg/database"
	"github.com/begenov/student-service/pkg/hash"
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

	hasher := hash.NewHash(cfg.Hash.Cost)

	tokenManager, err := auth.NewManager(cfg.JWT.SigningKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	repos := repository.NewRepository(db)

	service := service.NewService(repos, hasher, tokenManager, *cfg)

	handler := delivery.NewHandler(service, tokenManager)

	srv := server.NewServer(cfg, handler.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	log.Println("Server started")

	quit := make(chan os.Signal, 1)

	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Fatalf("failed to stop server: %v", err)
	}
}
