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

type signInInput struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

// @Summary		Sign-in
// @Tags			Student
// @Description	Sign-in
// @Accept			json
// @Produce		json
// @Param			account	body		signInInput	true	"Student"
// @Success		200		{object}	Resposne
// @Failure		400		{object}	Resposne
// @Failure		500		{object}	Resposne
// @Router			/students/sign-in [post]
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

// @Summary		Refresh Token
// @Tags			Student
// @Description	Refresh Token
// @Accept			json
// @Produce		json
// @Param			account	body		domain.Session	true	"Student"
// @Success		200		{object}	domain.Token
// @Failure		400		{object}	Resposne
// @Failure		500		{object}	Resposne
// @Router			/admin/auth/refresh [post]
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

// @Summary		 Get Students By CoursesID
// @Tags			Students
// @Description	 Get Students By CoursesID
// @Accept			json
// @Produce		json
// @Param			id path string		true	"course id"
// @Success		200		{byte}	[]byte
// @Failure		500		{object}	Resposne
// @Router			/students/{id}/courses [get]
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

// @Summary		 Get Courses By Student
// @Security StudentAuth
// @Tags			Students
// @Description	 Get Courses By Student
// @Accept			json
// @Produce		json
// @Success		200		{byte}	[]byte
// @Failure		500		{object}	Resposne
// @Router			/students/courses [get]
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
