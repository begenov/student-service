package v1

import (
	"github.com/begenov/student-service/internal/service"
	"github.com/begenov/student-service/pkg/auth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services     *service.Service
	tokenManager auth.TokenManager
	responseCh   chan []byte
}

func NewHandler(service *service.Service, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     service,
		tokenManager: tokenManager,
		responseCh:   make(chan []byte),
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		go h.consumeResponseMessages()
		h.initAdminStudentsRouter(v1)
		h.initStudentsRoutes(v1)
	}
}
