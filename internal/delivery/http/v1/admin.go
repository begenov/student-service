package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/begenov/student-service/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initAdminRoutes(api *gin.RouterGroup) {
	admin := api.Group("/admin")
	{
		admin.POST("/sign-up", h.adminSignUp)
		admin.POST("/sign-in", h.adminSignIn)
		admin.POST("/auth/refresh", h.adminRefreshToken)
		authentocated := admin.Group("/", h.adminIdentity)
		{
			students := authentocated.Group("/students")
			{
				students.POST("/create", h.adminCreatestudent)
				students.GET("/:id", h.adminGetStudentByID)
				students.PUT("/update/:id", h.adminUpdateStudent)
				students.DELETE("/delete/:id", h.adminDeleteStudent)
				students.GET("/:id/courses", h.adminGetCoursesStudents)
			}
		}
	}
}

type inputAdmin struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Name     string `json:"name" binding:"required,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func (h *Handler) adminSignUp(ctx *gin.Context) {
	var inp inputAdmin

	if err := ctx.BindJSON(&inp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect input data format",
		})
		return
	}

	if err := h.services.Admins.SignUp(ctx, domain.Admin{
		Email:    inp.Email,
		Name:     inp.Name,
		Password: inp.Password,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error when registering as an administrator",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Administrator successfully registered",
	})

}

type signInInput struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func (h *Handler) adminSignIn(ctx *gin.Context) {
	var inp signInInput

	if err := ctx.BindJSON(&inp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect input data format",
		})
		return
	}

	token, err := h.services.Admins.SignIn(context.Background(), inp.Email, inp.Password)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email address or password"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "A login error occurred"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) adminRefreshToken(ctx *gin.Context) {
	var inp domain.Session
	if err := ctx.BindJSON(&inp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect input data format",
		})
		return
	}

	token, err := h.services.Admins.GetByRefreshToken(context.Background(), inp.RefreshToken)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

}
