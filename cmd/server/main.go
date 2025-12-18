package main

import (
	"log"

	"echo-golang/internal/admin"
	"echo-golang/internal/config"
	"echo-golang/internal/database"
	"echo-golang/internal/handlers"
	"echo-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Set Gin mode
	if config.AppConfig.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	// Initialize Gin router
	r := gin.Default()

	// Middleware
	r.Use(middleware.SetupCORS())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "basketball-api",
		})
	})

	// Initialize handlers
	authHandler := handlers.NewAuthHandler()

	// Create default admin user
	if err := admin.CreateDefaultAdmin(); err != nil {
		log.Printf("Warning: Failed to create default admin: %v", err)
	}

	// API routes
	api := r.Group("/api/v1")
	{
		// Public auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
			auth.POST("/refresh", authHandler.RefreshToken)
		}

		// Protected routes
		protected := api.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// Auth routes (protected)
			protected.GET("/auth/me", authHandler.GetCurrentUser)

			// Admin routes
			admin.SetupAdminRoutes(protected)
		}
	}

	// Start server
	port := ":" + config.AppConfig.Port
	log.Printf("Server starting on port %s", config.AppConfig.Port)
	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

