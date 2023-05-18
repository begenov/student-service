package api

import (
	"net/http"

	"github.com/begenov/test-task-backend/students-app/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initAdminRoutes(api *gin.RouterGroup) {
	admins := api.Group("/admins")
	{
		admins.POST("/sign-up")
		admins.POST("/sign-in")
		admins.POST("/auth/refresh")
	}
}

func (h *Handler) adminSignIn(ctx *gin.Context) {
	var input models.Admin

	if err := ctx.BindJSON(&input); err != nil {
		newResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	ctx.JSON(http.StatusOK, models.Token{})

}
