package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/begenov/student-service/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initStudentsRoutes(api *gin.RouterGroup) {
	students := api.Group("/students")
	{
		students.POST("/sign-in", h.studentsSignIn)
		students.POST("/auth/refresh", h.studentsRefreshToken)
		authenticated := students.Group("/", h.studentIdentity)
		{
			students.GET("/:id/students", h.getStudentsByCourseID)
			authenticated.GET("/courses", h.studentsGetCourses)
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

func (h *Handler) studentsRefreshToken(ctx *gin.Context) {
	var inp domain.Session
	if err := ctx.BindJSON(&inp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect input data format",
		})
		return
	}
	token, err := h.services.Students.GetByRefreshToken(context.Background(), inp.RefreshToken)
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

func (h *Handler) getStudentsByCourseID(ctx *gin.Context) {
	// responseHandler := func(message string) {
	// 	// Добавьте здесь логику обработки ответа от Kafka
	// 	log.Println("Received message from Kafka:", message)
	// }

	// Потребляем сообщения из Kafka
	// err := h.services.Kafka.ConsumeMessages("students", responseHandler)
	// if err != nil {
	// 	log.Println("Failed to consume messages from Kafka:", err)
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Failed to get information about students",
	// 	})
	// 	return
	// }

}

/*
func (h *Handler) getStudentsByCourseID(ctx *gin.Context) {
	id := ctx.Param("id")

	students, err := h.services.Students.GetStudentsByCoursesID(context.Background(), id)
	log.Println(students, id)
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
*/

var (
	api = "http://localhost:8080/api/v1/courses/"
)

func (h *Handler) studentsGetCourses(ctx *gin.Context) {
	studentID, ok := ctx.Get(studentCtx)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Student ID not found in context"})
		return
	}
	url := api + studentID.(string) + "/courses"
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Mistake in receiving courses"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the body of the answer",
		})
		return
	}

	fmt.Println(string(body))
	var courses domain.Response

	if err := json.Unmarshal(body, &courses); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error when decoding courses"})
		return
	}
	if len(courses.Courses) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Courses not found.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"courses": courses})
}
