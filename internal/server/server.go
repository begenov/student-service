package server

import (
	"context"
	"net/http"

	"github.com/begenov/student-servcie/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, hadnler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.Server.Port,
			Handler:        hadnler,
			ReadTimeout:    cfg.Server.ReadTimeout,
			WriteTimeout:   cfg.Server.WriteTimeout,
			MaxHeaderBytes: cfg.Server.MaxHeaderMegabytes << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
