package admin

import (
	"log"

	"echo-golang/internal/config"
	"echo-golang/internal/database"
	"echo-golang/internal/middleware"
	"echo-golang/internal/models"
	"echo-golang/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SetupAdminRoutes sets up admin panel routes
func SetupAdminRoutes(r *gin.RouterGroup) {
	// Admin routes require authentication and admin role
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.RequireAdmin())
	{
		// Admin dashboard
		admin.GET("/dashboard", AdminDashboard)
		
		// User management
		admin.GET("/users", GetUsers)
		admin.GET("/users/:id", GetUser)
		admin.POST("/users", CreateUser)
		admin.PUT("/users/:id", UpdateUser)
		admin.DELETE("/users/:id", DeleteUser)
		
		// Organization management
		admin.GET("/organizations", GetOrganizations)
		admin.GET("/organizations/:id", GetOrganization)
		admin.POST("/organizations", CreateOrganization)
		admin.PUT("/organizations/:id", UpdateOrganization)
		admin.DELETE("/organizations/:id", DeleteOrganization)
	}
}

// AdminDashboard returns admin dashboard data
func AdminDashboard(c *gin.Context) {
	// Get statistics for dashboard
	// This is a placeholder - implement actual statistics later
	
	utils.SuccessResponse(c, gin.H{
		"total_users":         0,
		"total_organizations":  0,
		"total_teams":          0,
		"total_matches":        0,
		"active_matches":      0,
	}, "Dashboard data retrieved")
}

// GetUsers gets list of users (admin only)
func GetUsers(c *gin.Context) {
	// Placeholder - implement user listing
	utils.SuccessResponse(c, gin.H{
		"users": []models.User{},
		"total": 0,
	}, "Users retrieved")
}

// GetUser gets a single user
func GetUser(c *gin.Context) {
	userID := c.Param("id")
	id, err := uuid.Parse(userID)
	if err != nil {
		utils.BadRequest(c, "Invalid user ID", nil)
		return
	}

	var user models.User
	if err := database.DB.Preload("Organization").Where("id = ?", id).First(&user).Error; err != nil {
		utils.NotFound(c, "User not found")
		return
	}

	utils.SuccessResponse(c, user, "User retrieved")
}

// CreateUser creates a new user (admin only)
func CreateUser(c *gin.Context) {
	// Placeholder - implement user creation
	utils.SuccessResponse(c, nil, "User created")
}

// UpdateUser updates a user
func UpdateUser(c *gin.Context) {
	// Placeholder - implement user update
	utils.SuccessResponse(c, nil, "User updated")
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	id, err := uuid.Parse(userID)
	if err != nil {
		utils.BadRequest(c, "Invalid user ID", nil)
		return
	}

	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		utils.InternalServerError(c, "Failed to delete user")
		return
	}

	utils.SuccessResponse(c, nil, "User deleted successfully")
}

// GetOrganizations gets list of organizations
func GetOrganizations(c *gin.Context) {
	var orgs []models.Organization
	if err := database.DB.Preload("AdminUser").Find(&orgs).Error; err != nil {
		utils.InternalServerError(c, "Failed to fetch organizations")
		return
	}

	utils.SuccessResponse(c, gin.H{
		"organizations": orgs,
		"total":        len(orgs),
	}, "Organizations retrieved")
}

// GetOrganization gets a single organization
func GetOrganization(c *gin.Context) {
	orgID := c.Param("id")
	id, err := uuid.Parse(orgID)
	if err != nil {
		utils.BadRequest(c, "Invalid organization ID", nil)
		return
	}

	var org models.Organization
	if err := database.DB.Preload("AdminUser").Where("id = ?", id).First(&org).Error; err != nil {
		utils.NotFound(c, "Organization not found")
		return
	}

	utils.SuccessResponse(c, org, "Organization retrieved")
}

// CreateOrganization creates a new organization
func CreateOrganization(c *gin.Context) {
	// Placeholder - implement organization creation
	utils.SuccessResponse(c, nil, "Organization created")
}

// UpdateOrganization updates an organization
func UpdateOrganization(c *gin.Context) {
	// Placeholder - implement organization update
	utils.SuccessResponse(c, nil, "Organization updated")
}

// DeleteOrganization deletes an organization
func DeleteOrganization(c *gin.Context) {
	orgID := c.Param("id")
	id, err := uuid.Parse(orgID)
	if err != nil {
		utils.BadRequest(c, "Invalid organization ID", nil)
		return
	}

	if err := database.DB.Delete(&models.Organization{}, id).Error; err != nil {
		utils.InternalServerError(c, "Failed to delete organization")
		return
	}

	utils.SuccessResponse(c, nil, "Organization deleted successfully")
}

// CreateDefaultAdmin creates a default admin user if it doesn't exist
func CreateDefaultAdmin() error {
	var count int64
	database.DB.Model(&models.User{}).Where("role = ?", models.RoleSuperAdmin).Count(&count)
	
	if count > 0 {
		log.Println("Admin user already exists")
		return nil
	}

	// Create default admin
	hashedPassword, err := utils.HashPassword("admin123")
	if err != nil {
		return err
	}

	admin := &models.User{
		Email:    "admin@basketball.com",
		Password: hashedPassword,
		Role:     models.RoleSuperAdmin,
		FullName: "System Administrator",
		Status:   models.UserStatusActive,
	}

	if err := database.DB.Create(admin).Error; err != nil {
		return err
	}

	log.Println("Default admin user created:")
	log.Println("Email: admin@basketball.com")
	log.Println("Password: admin123")
	log.Println("Please change the password after first login!")

	return nil
}

