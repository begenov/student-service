package handlers

import (
	"github.com/begenov/student-servcie/internal/config"
	"github.com/begenov/student-servcie/internal/handlers/api"
	"github.com/begenov/student-servcie/internal/services"
	"github.com/begenov/student-servcie/pkg/auth"
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
