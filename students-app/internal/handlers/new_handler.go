package handlers

import (
	"github.com/begenov/test-task-backend/students-app/internal/config"
	"github.com/begenov/test-task-backend/students-app/internal/handlers/api"
	"github.com/begenov/test-task-backend/students-app/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
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