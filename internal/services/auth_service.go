package services

import (
	"errors"
	"time"

	"echo-golang/internal/models"
	"echo-golang/internal/repositories"
	"echo-golang/internal/utils"

	"github.com/google/uuid"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repositories.NewUserRepository(),
	}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Email          string    `json:"email" binding:"required,email"`
	Password       string    `json:"password" binding:"required,min=6"`
	FullName       string    `json:"full_name" binding:"required"`
	Role           string    `json:"role" binding:"required,oneof=super_admin org_admin team_member public"`
	OrganizationID *uuid.UUID `json:"organization_id,omitempty"`
	Phone          string    `json:"phone,omitempty"`
}

type LoginResponse struct {
	User         *models.User `json:"user"`
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresIn    int64        `json:"expires_in"`
}

// Login authenticates a user and returns JWT tokens
func (s *AuthService) Login(req LoginRequest) (*LoginResponse, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check if user is active
	if !user.IsActive() {
		return nil, errors.New("account is inactive")
	}

	// Verify password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	// Generate tokens
	accessToken, err := utils.GenerateToken(user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	refreshToken, err := utils.GenerateRefreshToken(user)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	// Update last login
	_ = s.userRepo.UpdateLastLogin(user.ID)

	// Calculate expires in (seconds)
	expiresIn := int64(15 * 60) // 15 minutes

	return &LoginResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}, nil
}

// Register creates a new user
func (s *AuthService) Register(req RegisterRequest) (*models.User, error) {
	// Check if email already exists
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create user
	user := &models.User{
		Email:          req.Email,
		Password:       hashedPassword,
		Role:           models.UserRole(req.Role),
		OrganizationID: req.OrganizationID,
		FullName:       req.FullName,
		Phone:          req.Phone,
		Status:         models.UserStatusActive,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("failed to create user")
	}

	// Get created user with relationships
	createdUser, err := s.userRepo.GetByID(user.ID)
	if err != nil {
		return nil, errors.New("failed to fetch created user")
	}

	return createdUser, nil
}

// GetCurrentUser gets the current authenticated user
func (s *AuthService) GetCurrentUser(userID uuid.UUID) (*models.User, error) {
	return s.userRepo.GetByID(userID)
}

