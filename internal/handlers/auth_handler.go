package handlers

import (
	"echo-golang/internal/middleware"
	"echo-golang/internal/services"
	"echo-golang/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: services.NewAuthService(),
	}
}

// Login handles user login
// @Summary Login user
// @Description Authenticate user and return JWT tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body services.LoginRequest true "Login credentials"
// @Success 200 {object} services.LoginResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req services.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request data", err.Error())
		return
	}

	response, err := h.authService.Login(req)
	if err != nil {
		utils.Unauthorized(c, err.Error())
		return
	}

	utils.SuccessResponse(c, response, "Login successful")
}

// Register handles user registration
// @Summary Register new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param request body services.RegisterRequest true "Registration data"
// @Success 201 {object} models.User
// @Failure 400 {object} utils.APIResponse
// @Failure 409 {object} utils.APIResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req services.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request data", err.Error())
		return
	}

	user, err := h.authService.Register(req)
	if err != nil {
		utils.BadRequest(c, err.Error(), nil)
		return
	}

	c.JSON(201, utils.APIResponse{
		Success: true,
		Data:    user,
		Message: "User registered successfully",
	})
}

// GetCurrentUser gets the current authenticated user
// @Summary Get current user
// @Description Get information about the currently authenticated user
// @Tags auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {object} utils.APIResponse
// @Router /auth/me [get]
func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := middleware.GetUserIDFromContext(c)
	if !exists {
		utils.Unauthorized(c, "User not found in context")
		return
	}

	user, err := h.authService.GetCurrentUser(userID)
	if err != nil {
		utils.NotFound(c, "User not found")
		return
	}

	utils.SuccessResponse(c, user, "User retrieved successfully")
}

// RefreshToken handles token refresh
// @Summary Refresh access token
// @Description Refresh access token using refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param refresh_token body object{refresh_token=string} true "Refresh token"
// @Success 200 {object} object{access_token=string,refresh_token=string,expires_in=int64}
// @Failure 401 {object} utils.APIResponse
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Refresh token required", err.Error())
		return
	}

	// Validate refresh token
	claims, err := utils.ValidateToken(req.RefreshToken)
	if err != nil {
		utils.Unauthorized(c, "Invalid refresh token")
		return
	}

	// Get user
	user, err := h.authService.GetCurrentUser(claims.UserID)
	if err != nil {
		utils.Unauthorized(c, "User not found")
		return
	}

	// Generate new tokens
	accessToken, err := utils.GenerateToken(user)
	if err != nil {
		utils.InternalServerError(c, "Failed to generate token")
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user)
	if err != nil {
		utils.InternalServerError(c, "Failed to generate refresh token")
		return
	}

	utils.SuccessResponse(c, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"expires_in":    900, // 15 minutes
	}, "Token refreshed successfully")
}

