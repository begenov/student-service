package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/begenov/test-task-backend/students-app/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initStudentsRoutes(api *gin.RouterGroup) {
	students := api.Group("/students")
	{
		students.POST("/create", h.studentCreate)
		students.GET("/:id", h.studentGetID)
		auth := students.Group("/")

		{
			auth.PUT("/update/:id", h.studentUpdate)
			auth.DELETE("/delete/:id", h.studentDelete)
			auth.GET("/:id/courses", h.studentByIDCourses)
		}
	}
}

func (h *Handler) studentCreate(ctx *gin.Context) {
	var newStudent models.Student

	if err := ctx.BindJSON(&newStudent); err != nil {
		log.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON data",
		})
		return
	}

	if err := h.services.Students.CreateStudent(ctx, newStudent); err != nil {
		log.Println("Error creating student:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create student",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Student created successfully",
	})

}

func (h *Handler) studentGetID(ctx *gin.Context) {

	studentID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		log.Printf("Invalid student ID: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid student ID",
		})
		return
	}

	student, err := h.services.Students.GetStudentByID(ctx, studentID)

	if err != nil {
		log.Printf("Failed to get student by ID: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get student by ID",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}

func (h *Handler) studentUpdate(ctx *gin.Context) {
	studentID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("Invalid student ID: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid student ID",
		})
		return
	}
	var student models.Student

	if err := ctx.BindJSON(student); err != nil {
		log.Println("Error binding JSON:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON data",
		})
		return
	}

	student.ID = studentID

	if err = h.services.Students.Update(ctx, student); err != nil {
		log.Printf("Failed to update student: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update student",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Student update successfully",
	})

}

func (h *Handler) studentDelete(ctx *gin.Context) {
	studentID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("Invalid student ID: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid student ID",
		})
		return
	}

	if err := h.services.Students.Delete(ctx, studentID); err != nil {
		log.Printf("Failed to delete student: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete student",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Student delete successfully",
	})
}

func (h *Handler) studentByIDCourses(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Printf("Invalid student ID: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid student ID",
		})
		return
	}
	students, err := h.services.Students.ByIDCourses(ctx, id)
	if err != nil {
		log.Printf("Invalid student ID: %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid student ID",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"students": students,
	})
}
