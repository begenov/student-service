package v1

import (
	"context"
	"errors"
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
		newResponse(ctx, http.StatusBadRequest, "Incorrect input data format")
		return
	}
	if err := h.services.Students.Create(context.Background(), domain.Student{
		Email:    inp.Email,
		Name:     inp.Name,
		Password: inp.Password,
		GPA:      inp.GPA,
		Courses:  inp.Courses,
	}); err != nil {
		newResponse(ctx, http.StatusInternalServerError, "Error when creating a student")
		return
	}

	ctx.JSON(http.StatusCreated, Resposne{"The student is successfully established"})
}

func (h *Handler) adminGetStudentByID(ctx *gin.Context) {
	param := ctx.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, "Incorrect student ID format")
		return
	}

	student, err := h.services.Students.GetStudentByID(ctx, id)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, "Mistake in getting a student")
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (h *Handler) adminUpdateStudent(ctx *gin.Context) {
	var inp domain.UpdateStudentInput
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, "Incorrect student ID format")
		return
	}
	if err := ctx.BindJSON(&inp); err != nil {
		newResponse(ctx, http.StatusBadRequest, "Incorrect input data format")
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
		newResponse(ctx, http.StatusBadRequest, "Error when updating a student")
		return
	}

	ctx.JSON(http.StatusOK, Resposne{"Student has been successfully updated"})
}

func (h *Handler) adminDeleteStudent(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, "Incorrect student ID format")
		return
	}
	if err := h.services.Students.Delete(context.Background(), id); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			newResponse(ctx, http.StatusBadRequest, "Student not found")
			return
		}
		newResponse(ctx, http.StatusInternalServerError, "Error when deleting a student")
		return
	}
	ctx.JSON(http.StatusOK, Resposne{"Student successfully deleted"})
}

func (h *Handler) adminGetCoursesStudents(ctx *gin.Context) {
	param := ctx.Param("id")
	err := h.services.Kafka.SendMessages("courses-request", param)
	if err != nil {
		newResponse(ctx, http.StatusInternalServerError, "Failed to get information about courses")
		return
	}
	responseData := <-h.responseCh

	ctx.Data(http.StatusOK, "application/json", responseData)
}

/*
func (h *Handler) adminGetCoursesStudents02(ctx *gin.Context) {
	id := ctx.Param("id")
	url := api + id + "/courses"
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
