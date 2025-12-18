package middleware

import (
	"echo-golang/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupCORS configures CORS middleware
func SetupCORS() gin.HandlerFunc {
	config := cors.Config{
		AllowOrigins:     config.AppConfig.CORSAllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
		ExposeHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           86400, // 24 hours
	}

	return cors.New(config)
}

