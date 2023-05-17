package api

import "github.com/gin-gonic/gin"

func (h *Handler) initStudentsRoutes(api *gin.RouterGroup) {
	students := api.Group("/students")
	{
		students.POST("/sign-up", h.studentSignUp)
		students.POST("/sig-in", h.studentSingIn)
		auth := students.Group("/") // TODO middleaware
		{
			auth.PUT("/update/:id", h.studentUpdate)
			auth.DELETE("/delete/:id", h.studentDelete)
			auth.GET("/:id/courses", h.studentByIDCourses)
		}
	}
}

func (h *Handler) studentSignUp(ctx *gin.Context) {

}

func (h *Handler) studentSingIn(ctx *gin.Context) {
}

func (h *Handler) studentUpdate(ctx *gin.Context) {
}

func (h *Handler) studentDelete(ctx *gin.Context) {
}

func (h *Handler) studentByIDCourses(ctx *gin.Context) {

}
