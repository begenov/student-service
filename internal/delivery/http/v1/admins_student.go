package v1

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/begenov/student-service/internal/domain"
	"github.com/gin-gonic/gin"
)

type inputStudent struct {
	Email    string   `json:"email" binding:"required,email,max=64"`
	Name     string   `json:"name" binding:"required,min=3,max=64"`
	Password string   `json:"password" binding:"required,min=8,max=64"`
	GPA      float64  `json:"gpa" binding:"required"`
	Courses  []string `json:"courses"`
}

func (h *Handler) adminCreatestudent(ctx *gin.Context) {
	var inp inputStudent

	if err := ctx.BindJSON(&inp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect input data format",
		})
		return
	}
	if err := h.services.Students.Create(context.Background(), domain.Student{
		Email:    inp.Email,
		Name:     inp.Name,
		Password: inp.Password,
		GPA:      inp.GPA,
		Courses:  inp.Courses,
	}); err != nil {
		log.Printf("Error when creating a student: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error when creating a student",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "The student is successfully established",
	})

}

func (h *Handler) adminGetStudentByID(ctx *gin.Context) {
	param := ctx.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect student ID format"})
		return
	}

	student, err := h.services.Students.GetStudentByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Mistake in getting a student"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"student": student})
}

func (h *Handler) adminUpdateStudent(ctx *gin.Context) {
	var inp domain.UpdateStudentInput
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect student ID format"})
		return
	}
	if err := ctx.BindJSON(&inp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect input data format"})
		return
	}
	if err := h.services.Students.Update(context.Background(), domain.Student{
		Email:    inp.Email,
		Name:     inp.Name,
		Password: inp.Password,
		GPA:      inp.GPA,
		Courses:  inp.Courses,
		ID:       id,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error when updating a student"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Student has been successfully "})
}

func (h *Handler) adminDeleteStudent(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect student ID format"})
		return
	}
	if err := h.services.Students.Delete(context.Background(), id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Student not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error when deleting a student"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Student successfully deleted"})
}

func (h *Handler) adminGetByCoursesIDstudent(ctx *gin.Context) {
	id := ctx.Param("id")

	students, err := h.services.Students.GetStudentsByCoursesID(context.Background(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Students not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error in getting students by course ID"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"students": students,
	})
}
