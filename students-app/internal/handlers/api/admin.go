package api

import (
	"net/http"

	"github.com/begenov/test-task-backend/students-app/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initAdminRoutes(api *gin.RouterGroup) {
	admins := api.Group("/admins")
	{
		admins.POST("/sign-up", h.adminSignUp)
		admins.POST("/sign-in", h.adminsSignIn)
		admins.POST("/auth/refresh", h.adminRefresh)
	}
}

func (h *Handler) adminSignUp(ctx *gin.Context) {
	var input models.Admin

	if err := ctx.BindJSON(&input); err != nil {
		newResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	if err := h.services.Admin.SignUpAdmin(ctx, input); err != nil {
		newResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Admin created successfully",
	})
}

func (h *Handler) adminsSignIn(ctx *gin.Context) {
	var input models.Admin
	if err := ctx.BindJSON(&input); err != nil {
		newResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	tokens, err := h.services.Admin.SignInAdmin(ctx, input)
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": tokens,
	})
}

func (h *Handler) adminRefresh(ctx *gin.Context) {
	var input struct {
		AccessToken string `json:"access_token"`
	}

	if err := ctx.BindJSON(&input); err != nil {
		newResponse(ctx, http.StatusBadRequest, "invalid input body")
		return
	}

	tokens, err := h.services.Admin.RefreshToken(ctx, models.Token{
		AccessToken: input.AccessToken,
	})

	if err != nil {
		newResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": tokens,
	})

}
