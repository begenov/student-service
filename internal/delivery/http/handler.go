package http

import (
	"github.com/begenov/student-service/internal/config"
	v1 "github.com/begenov/student-service/internal/delivery/http/v1"
	"github.com/begenov/student-service/internal/service"
	"github.com/begenov/student-service/pkg/auth"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/begenov/student-service/docs"
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

func (h *Handler) Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	h.initAPI(router)
	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.tokenManager)
	api := router.Group("/api")
	{

		handlerV1.Init(api)
	}
}
