package server

import (
	"context"
	"net/http"

	"github.com/begenov/test-task-backend/students-app/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, hadnler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
