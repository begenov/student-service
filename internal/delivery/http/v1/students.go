package v1

import (
	"context"
	"errors"
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
		newResponse(ctx, http.StatusBadRequest, "Incorrect input data format")
		return
	}

	token, err := h.services.Students.GetByEmail(context.Background(), inp.Email, inp.Password)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			newResponse(ctx, http.StatusBadRequest, "Incorrect email address or password")
			return
		}

		newResponse(ctx, http.StatusInternalServerError, "A login error occurred")
		return
	}

	ctx.JSON(http.StatusOK, token)

}

func (h *Handler) studentsRefreshToken(ctx *gin.Context) {
	var inp domain.Session
	if err := ctx.BindJSON(&inp); err != nil {
		newResponse(ctx, http.StatusBadRequest, "Incorrect input data format")
		return
	}
	token, err := h.services.Students.GetByRefreshToken(context.Background(), inp.RefreshToken)
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

func (h *Handler) getStudentsByCourseID(ctx *gin.Context) {
	id := ctx.Param("id")

	students, err := h.services.Students.GetStudentsByCoursesID(context.Background(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			newResponse(ctx, http.StatusBadRequest, "Students not found")
			return
		}
		newResponse(ctx, http.StatusBadRequest, "Error in getting students by course ID")
		return
	}

	ctx.JSON(http.StatusOK, students)
}

func (h *Handler) studentsGetCourses(ctx *gin.Context) {
	studentID, ok := ctx.Get(studentCtx)
	if !ok {
		newResponse(ctx, http.StatusBadRequest, "Student ID not found in context")
		return
	}

	err := h.services.Kafka.SendMessages("courses-request", studentID.(string))
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, "Failed to get information about courses")
		return
	}

	responseData := <-h.responseCh
	ctx.Data(http.StatusOK, "application/json", responseData)
}
func (h *Handler) consumeResponseMessages() {
	err := h.services.Kafka.ConsumeMessages("courses-response", h.handleResponseMessage)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) handleResponseMessage(message string) {
	h.responseCh <- []byte(message)
}

/*
var (
	api = "http://localhost:8080/api/v1/courses/"
)

func (h *Handler) studentsGetCourses02(ctx *gin.Context) {
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
*/
