package routes

import (
	"github.com/gin-gonic/gin"
	"92HW/controllers"
	"92HW/middleware"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/cors/origins", controllers.AddTrustedOrigin)
		auth.DELETE("/cors/origins", controllers.RemoveTrustedOrigin)
		auth.GET("/cors/origins", controllers.GetTrustedOrigins)
	}
}
