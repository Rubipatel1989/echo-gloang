package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to release mode in production
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Health check endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Basketball API Server",
			"status":  "running",
		})
	})

	// API v1 routes will be added here
	api := r.Group("/api/v1")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "healthy",
			})
		})
	}

	// Start server
	r.Run(":8080")
}
