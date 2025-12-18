package repositories

import (
	"time"

	"echo-golang/internal/database"
	"echo-golang/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.DB,
	}
}

// Create creates a new user
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// GetByID gets a user by ID
func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Organization").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail gets a user by email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Organization").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates a user
func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// UpdateLastLogin updates the last login time
func (r *UserRepository) UpdateLastLogin(userID uuid.UUID) error {
	now := time.Now()
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("last_login_at", now).Error
}

// List gets a list of users with pagination
func (r *UserRepository) List(offset, limit int, filters map[string]interface{}) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := r.db.Model(&models.User{})

	// Apply filters
	if role, ok := filters["role"]; ok {
		query = query.Where("role = ?", role)
	}
	if status, ok := filters["status"]; ok {
		query = query.Where("status = ?", status)
	}
	if orgID, ok := filters["organization_id"]; ok {
		query = query.Where("organization_id = ?", orgID)
	}

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get users
	err := query.Preload("Organization").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&users).Error

	return users, total, err
}

// Delete soft deletes a user
func (r *UserRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.User{}, id).Error
}

