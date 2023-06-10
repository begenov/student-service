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

// adminSignUp godoc
// @Summary Admin Sign Up
// @Description Register as an administrator
// @Tags Admins
// @Accept json
// @Produce json
// @Param input body string true "Input data"
// @Success 201 {object} string
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /admin/signup [post]
func (h *Handler) adminSignUp(ctx *gin.Context) {
	var inp inputAdmin

	if err := ctx.BindJSON(&inp); err != nil {
		newResponse(ctx, http.StatusBadRequest, "Incorrect input data format")
		return
	}

	if err := h.services.Admins.SignUp(ctx, domain.Admin{
		Email:    inp.Email,
		Name:     inp.Name,
		Password: inp.Password,
	}); err != nil {
		newResponse(ctx, http.StatusInternalServerError, "Error when registering as an administrator")
		return
	}
	ctx.JSON(http.StatusCreated, "Administrator successfully registered")
}

type signInInput struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

// adminSignIn godoc
// @Summary Admin Sign In
// @Description Sign in as an administrator
// @Tags Admins
// @Accept json
// @Produce json
// @Param input body signInInput true "Input data"
// @Success 200 {object} domain.Token
// @Failure 400 {object} Response
// @Failure 401 {object} Response
// @Failure 500 {object} Response
// @Router /admin/signin [post]
func (h *Handler) adminSignIn(ctx *gin.Context) {
	var inp signInInput

	if err := ctx.BindJSON(&inp); err != nil {
		newResponse(ctx, http.StatusBadRequest, "Incorrect input data format")
		return
	}

	token, err := h.services.Admins.SignIn(context.Background(), inp.Email, inp.Password)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			newResponse(ctx, http.StatusUnauthorized, "Incorrect email address or password")
			return
		}

		newResponse(ctx, http.StatusInternalServerError, "A login error occurred")
		return
	}
	ctx.JSON(http.StatusOK, token)
}

// adminRefreshToken godoc
// @Summary Admin Refresh Token
// @Description Refresh the authentication token for an administrator
// @Tags Admins
// @Accept json
// @Produce json
// @Param input body domain.Session true "Refresh token data"
// @Success 200 {object} domain.Token
// @Failure 400 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /admin/refresh-token [post]
func (h *Handler) adminRefreshToken(ctx *gin.Context) {
	var inp domain.Session
	if err := ctx.BindJSON(&inp); err != nil {
		newResponse(ctx, http.StatusBadRequest, "Incorrect input data format")

		return
	}

	token, err := h.services.Admins.GetByRefreshToken(context.Background(), inp.RefreshToken)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			newResponse(ctx, http.StatusBadRequest, "Refresh token not found")
			return
		}
		newResponse(ctx, http.StatusInternalServerError, "Error retrieving token")

		return
	}

	ctx.JSON(http.StatusOK, token)
}
