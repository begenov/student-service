package api

import (
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

func (h *Handler) Init(api *gin.RouterGroup) {
	{
		h.initStudentsRoutes(api)
	}
}
