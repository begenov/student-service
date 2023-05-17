package api

import (
	"log"
	"net/http"

	"github.com/begenov/test-task-backend/students-app/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initStudentsRoutes(api *gin.RouterGroup) {
	students := api.Group("/students")
	{
		students.POST("/create", h.studentCreate)
		students.GET("/:id", h.studentGet)
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

	if err := h.services.


}

func (h *Handler) studentGet(ctx *gin.Context) {
}

func (h *Handler) studentUpdate(ctx *gin.Context) {
}

func (h *Handler) studentDelete(ctx *gin.Context) {
}

func (h *Handler) studentByIDCourses(ctx *gin.Context) {
}
