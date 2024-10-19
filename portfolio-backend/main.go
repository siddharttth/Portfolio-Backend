package main

import (
	"os"
	"portfolio-backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set gin to production mode
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Enhanced CORS configuration
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	r.Use(cors.New(config))

	// Routes
	r.GET("/api/education", services.GetEducation)
	r.GET("/api/experience", services.GetExperience)
	r.GET("/api/projects", services.GetProjects)
	r.GET("/api/skills", services.GetSkills)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r.Run(":" + port)
}
