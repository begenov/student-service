package v1

import (
	"github.com/begenov/student-service/internal/service"
	"github.com/begenov/student-service/pkg/auth"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services     *service.Service
	tokenManager auth.TokenManager
}

func NewHandler(service *service.Service, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     service,
		tokenManager: tokenManager,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAdminRoutes(v1)
		h.initStudentsRoutes(v1)
	}
}
