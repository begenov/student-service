package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/begenov/student-service/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initStudentsRoutes(api *gin.RouterGroup) {
	students := api.Group("/students")
	{
		students.POST("/sign-in", h.studentsSignIn)
		authenticated := students.Group("/", h.studentIdentity)
		{
			authenticated.GET("/:id/courses", h.studentsGetCourses)
		}
	}
}

func (h *Handler) studentsSignIn(ctx *gin.Context) {
	var inp signInInput
	if err := ctx.BindJSON(&inp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect input data format"})
		return
	}

	token, err := h.services.Students.GetByEmail(context.Background(), inp.Email, inp.Password)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			log.Printf("Incorrect email address or password: %s", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect email address or password"})
			return
		}
		log.Printf("A login error occurred: %s", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "A login error occurred"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

}

func (h *Handler) studentsGetCourses(ctx *gin.Context) {
	studentID := ctx.Param("id")
	url := fmt.Sprintf("%s/courses", studentID)
	resp, err := http.Get(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Mistake in receiving courses"})
		return
	}
	defer resp.Body.Close()

	var courses []domain.Courses
	if err = json.NewDecoder(resp.Body).Decode(&courses); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error when decoding courses"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"courses": courses})
}
