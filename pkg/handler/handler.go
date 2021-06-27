package handler

import (
	"github.com/DanYesmagulov/go-video-streaming/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	upload := router.Group("/upload")
	{
		upload.POST("/video", h.uploadVideo)
	}

	authenticatedGroup := router.Group("/api", h.userIdentity)
	{
		usersCourses := authenticatedGroup.Group("/course")
		{
			usersCourses.POST("/", h.createCourse)
			usersCourses.PATCH("/:id", h.updateCourse)
			usersCourses.DELETE("/:id", h.deleteCourse)
			usersCourses.POST("/:id/video-file", h.uploadVideo)
		}
	}

	api := router.Group("/api")
	{
		category := api.Group("/category")
		{
			category.POST("/", h.createCategory)
			category.GET("/", h.getAllCategory)
			category.GET("/:id", h.getCategoryById)
			category.DELETE("/:id", h.deleteCategory)
		}

		course := api.Group("/course")
		{
			course.GET("/", h.getAllCourse)
			course.GET("/:id", h.getCourseById)
		}
	}

	return router
}
