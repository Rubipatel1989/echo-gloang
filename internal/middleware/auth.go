package middleware

import (
	"echo-golang/internal/database"
	"echo-golang/internal/models"
	"echo-golang/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AuthMiddleware validates JWT token and sets user in context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Unauthorized(c, "Authorization header required")
			c.Abort()
			return
		}

		// Extract token
		token := utils.ExtractTokenFromHeader(authHeader)
		if token == "" {
			utils.Unauthorized(c, "Invalid authorization header format")
			c.Abort()
			return
		}

		// Validate token
		claims, err := utils.ValidateToken(token)
		if err != nil {
			utils.Unauthorized(c, "Invalid or expired token")
			c.Abort()
			return
		}

		// Get user from database to ensure they still exist and are active
		var user models.User
		if err := database.DB.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
			utils.Unauthorized(c, "User not found")
			c.Abort()
			return
		}

		// Check if user is active
		if !user.IsActive() {
			utils.Forbidden(c, "Account is inactive")
			c.Abort()
			return
		}

		// Set user data in context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)
		c.Set("organization_id", claims.OrganizationID)
		c.Set("user", &user)

		c.Next()
	}
}

// GetUserFromContext gets the user ID from context
func GetUserIDFromContext(c *gin.Context) (uuid.UUID, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return uuid.Nil, false
	}
	id, ok := userID.(uuid.UUID)
	return id, ok
}

// GetUserRoleFromContext gets the user role from context
func GetUserRoleFromContext(c *gin.Context) (models.UserRole, bool) {
	role, exists := c.Get("user_role")
	if !exists {
		return "", false
	}
	userRole, ok := role.(models.UserRole)
	return userRole, ok
}

// GetOrganizationIDFromContext gets the organization ID from context
func GetOrganizationIDFromContext(c *gin.Context) (*uuid.UUID, bool) {
	orgID, exists := c.Get("organization_id")
	if !exists {
		return nil, false
	}
	if orgID == nil {
		return nil, true
	}
	id, ok := orgID.(*uuid.UUID)
	return id, ok
}

