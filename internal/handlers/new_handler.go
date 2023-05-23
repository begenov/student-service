package handlers

import (
	"github.com/begenov/test-task-backend/pkg/auth"
	"github.com/begenov/test-task-backend/students-app/internal/config"
	"github.com/begenov/test-task-backend/students-app/internal/handlers/api"
	"github.com/begenov/test-task-backend/students-app/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Services
	manager  auth.TokenManager
}

func NewHandler(services *services.Services, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services: services,
		manager:  tokenManager,
	}
}

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	handlerAPI := api.NewHandler(h.services)

	api := router.Group("/api")

	{
		handlerAPI.Init(api)
	}

	return router
}
