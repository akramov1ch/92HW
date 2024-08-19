package main

import (
	"github.com/gin-gonic/gin"
	"92HW/config"
	"92HW/middleware"
	"92HW/routes"
)

func main() {
	r := gin.Default()

	config.ConnectRedis()
	config.SetupConfig()

	r.Use(middleware.CORSMiddleware())
	routes.SetupRoutes(r)

	r.Run(":8080")
}
