package api

import "github.com/gin-gonic/gin"

func (h *Handler) initAdminRoutes(api *gin.RouterGroup) {
	admins := api.Group("/admins")
	{
		admins.POST("/sign-up")
		admins.POST("/sign-in")
		admins.POST("/auth/refresh")
	}
}
