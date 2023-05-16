package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/begenov/test-task-backend/pkg/postgresql"
	"github.com/begenov/test-task-backend/students-app/internal/config"
	"github.com/begenov/test-task-backend/students-app/internal/handlers"
	"github.com/begenov/test-task-backend/students-app/internal/server"
	"github.com/begenov/test-task-backend/students-app/internal/services"
	"github.com/begenov/test-task-backend/students-app/internal/storage"
)

const (
	path_config = "../.env"
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

	storage := storage.NewStorage(db)

	services := services.NewService(storage)

	handlers := handlers.NewHandler(services)

	srv := server.NewServer(cfg, handlers.Init(cfg))

	go func() {
		if err := srv.Run(); errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server closed with error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Fatalf("error stopping HTTP server: %v", err)
	}

}
