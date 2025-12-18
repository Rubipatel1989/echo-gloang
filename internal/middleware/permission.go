package middleware

import (
	"echo-golang/internal/models"
	"echo-golang/internal/utils"

	"github.com/gin-gonic/gin"
)

// RequireRole checks if user has required role
func RequireRole(roles ...models.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := GetUserRoleFromContext(c)
		if !exists {
			utils.Unauthorized(c, "User role not found in context")
			c.Abort()
			return
		}

		// Check if user role is in allowed roles
		allowed := false
		for _, role := range roles {
			if userRole == role {
				allowed = true
				break
			}
		}

		if !allowed {
			utils.Forbidden(c, "Insufficient permissions")
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireAdmin checks if user is super admin
func RequireAdmin() gin.HandlerFunc {
	return RequireRole(models.RoleSuperAdmin)
}

// RequireOrgAdmin checks if user is organization admin or super admin
func RequireOrgAdmin() gin.HandlerFunc {
	return RequireRole(models.RoleSuperAdmin, models.RoleOrgAdmin)
}

// RequireAdminOrOrgAdmin checks if user is admin or org admin
func RequireAdminOrOrgAdmin() gin.HandlerFunc {
	return RequireRole(models.RoleSuperAdmin, models.RoleOrgAdmin)
}

// CheckOrganizationAccess checks if user can access organization
func CheckOrganizationAccess(orgIDGetter func(*gin.Context) uuid.UUID) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, _ := GetUserRoleFromContext(c)
		userOrgID, _ := GetOrganizationIDFromContext(c)

		// Super admin can access any organization
		if userRole == models.RoleSuperAdmin {
			c.Next()
			return
		}

		// Get organization ID from request
		requestedOrgID := orgIDGetter(c)

		// Org admin can only access their own organization
		if userRole == models.RoleOrgAdmin {
			if userOrgID == nil || *userOrgID != requestedOrgID {
				utils.Forbidden(c, "Access denied to this organization")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

