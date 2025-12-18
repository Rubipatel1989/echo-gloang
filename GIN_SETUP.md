# Gin Framework Setup Guide

## Overview

This project uses **Gin Framework** for the backend API. Gin is one of the fastest and most popular Go web frameworks.

## Key Features

- ✅ High performance (50,000-60,000 requests/second)
- ✅ Powerful routing and middleware
- ✅ Large ecosystem and community
- ✅ Excellent JSON binding
- ✅ WebSocket support via gorilla/websocket

## Dependencies

### Required Packages

```bash
# Gin framework
go get github.com/gin-gonic/gin

# WebSocket support
go get github.com/gorilla/websocket

# Database
go get gorm.io/gorm
go get gorm.io/driver/mysql

# JWT authentication
go get github.com/golang-jwt/jwt/v5

# Password hashing
go get golang.org/x/crypto/bcrypt

# Environment variables
go get github.com/joho/godotenv

# UUID generation
go get github.com/google/uuid

# Redis client
go get github.com/redis/go-redis/v9

# CORS middleware
go get github.com/gin-contrib/cors
```

## Gin Handler Pattern

### Basic Handler Structure

```go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetTeams(c *gin.Context) {
    // Get query parameters
    page := c.DefaultQuery("page", "1")
    limit := c.DefaultQuery("limit", "10")
    
    // Call service
    teams, err := teamService.GetTeams(page, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error": gin.H{
                "code": "INTERNAL_ERROR",
                "message": err.Error(),
            },
        })
        return
    }
    
    // Success response
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data": teams,
    })
}
```

## Middleware Pattern

### Authentication Middleware

```go
package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "success": false,
                "error": gin.H{
                    "code": "UNAUTHORIZED",
                    "message": "Authorization header required",
                },
            })
            c.Abort()
            return
        }
        
        // Validate token
        claims, err := validateToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "success": false,
                "error": gin.H{
                    "code": "INVALID_TOKEN",
                    "message": "Invalid or expired token",
                },
            })
            c.Abort()
            return
        }
        
        // Set user in context
        c.Set("user_id", claims.UserID)
        c.Set("role", claims.Role)
        c.Set("organization_id", claims.OrganizationID)
        
        c.Next()
    }
}
```

### CORS Middleware

```go
import "github.com/gin-contrib/cors"

func SetupCORS() gin.HandlerFunc {
    return cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    })
}
```

## Route Setup

### Basic Route Structure

```go
package main

import (
    "github.com/gin-gonic/gin"
    "your-project/internal/handlers"
    "your-project/internal/middleware"
)

func setupRoutes(r *gin.Engine) {
    // Public routes
    public := r.Group("/api/v1")
    {
        public.GET("/health", handlers.HealthCheck)
        public.POST("/auth/register", handlers.Register)
        public.POST("/auth/login", handlers.Login)
    }
    
    // Protected routes
    protected := r.Group("/api/v1")
    protected.Use(middleware.AuthMiddleware())
    {
        // User routes
        protected.GET("/auth/me", handlers.GetCurrentUser)
        
        // Team routes
        teams := protected.Group("/teams")
        {
            teams.GET("", handlers.GetTeams)
            teams.GET("/:id", handlers.GetTeam)
            teams.POST("", handlers.CreateTeam)
            teams.PUT("/:id", handlers.UpdateTeam)
            teams.DELETE("/:id", handlers.DeleteTeam)
        }
    }
    
    // Admin routes
    admin := r.Group("/api/v1/admin")
    admin.Use(middleware.AuthMiddleware())
    admin.Use(middleware.AdminOnly())
    {
        // Admin specific routes
    }
}
```

## WebSocket Setup

### WebSocket Handler with Gin

```go
package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow all origins in development
    },
}

func LiveMatchWebSocket(c *gin.Context) {
    matchID := c.Param("id")
    
    // Upgrade connection
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        return
    }
    defer conn.Close()
    
    // Subscribe to Redis pub/sub for this match
    // Send updates to client
    for {
        // Read from Redis channel
        // Send to WebSocket client
        err := conn.WriteJSON(update)
        if err != nil {
            break
        }
    }
}
```

### WebSocket Route

```go
r.GET("/ws/matches/:id/live", handlers.LiveMatchWebSocket)
```

## Error Handling

### Standardized Error Response

```go
package utils

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, statusCode int, code, message string) {
    c.JSON(statusCode, gin.H{
        "success": false,
        "error": gin.H{
            "code":    code,
            "message": message,
        },
    })
}

func SuccessResponse(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    data,
    })
}
```

## Request Binding

### JSON Binding

```go
type CreateTeamRequest struct {
    Name        string `json:"name" binding:"required"`
    CoachName   string `json:"coach_name" binding:"required"`
    CoachEmail  string `json:"coach_email" binding:"required,email"`
}

func CreateTeam(c *gin.Context) {
    var req CreateTeamRequest
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error": gin.H{
                "code": "VALIDATION_ERROR",
                "message": err.Error(),
            },
        })
        return
    }
    
    // Process request...
}
```

### Query Parameters

```go
func GetTeams(c *gin.Context) {
    page := c.DefaultQuery("page", "1")
    limit := c.DefaultQuery("limit", "10")
    status := c.Query("status")
    
    // Use parameters...
}
```

### Path Parameters

```go
func GetTeam(c *gin.Context) {
    teamID := c.Param("id")
    
    // Use teamID...
}
```

## Best Practices

1. **Use Route Groups** - Organize related routes
2. **Middleware Chain** - Apply middleware in order
3. **Error Handling** - Always return consistent error format
4. **Context Usage** - Use `c.Set()` and `c.Get()` for request-scoped data
5. **Binding Validation** - Use struct tags for validation
6. **Response Format** - Keep responses consistent
7. **Abort on Errors** - Use `c.Abort()` in middleware to stop chain

## Performance Tips

1. **Release Mode** - Use `gin.SetMode(gin.ReleaseMode)` in production
2. **Connection Pooling** - Configure database connection pool
3. **Caching** - Use Redis for frequently accessed data
4. **Async Operations** - Use goroutines for non-blocking operations
5. **Response Compression** - Enable gzip compression

## Testing

### Unit Testing Handlers

```go
func TestGetTeams(t *testing.T) {
    gin.SetMode(gin.TestMode)
    
    r := gin.New()
    r.GET("/teams", handlers.GetTeams)
    
    req, _ := http.NewRequest("GET", "/teams", nil)
    w := httptest.NewRecorder()
    
    r.ServeHTTP(w, req)
    
    assert.Equal(t, http.StatusOK, w.Code)
}
```

## Resources

- [Gin Documentation](https://gin-gonic.com/docs/)
- [Gin GitHub](https://github.com/gin-gonic/gin)
- [Gorilla WebSocket](https://github.com/gorilla/websocket)
- [GORM Documentation](https://gorm.io/docs/)

