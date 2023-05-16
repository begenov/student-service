package api

import "github.com/gin-gonic/gin"

func (h *Handler) initStudentsRoutes(api *gin.RouterGroup) {
	students := api.Group("/students")
	{
		students.GET("")
	}
}
