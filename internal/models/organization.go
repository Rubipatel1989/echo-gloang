package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationStatus string

const (
	OrgStatusActive   OrganizationStatus = "active"
	OrgStatusInactive OrganizationStatus = "inactive"
)

type Organization struct {
	ID          uuid.UUID         `gorm:"type:char(36);primary_key" json:"id"`
	Name        string            `gorm:"type:varchar(255);not null" json:"name"`
	Email       string            `gorm:"type:varchar(255)" json:"email"`
	Phone       string            `gorm:"type:varchar(20)" json:"phone"`
	Address     string            `gorm:"type:text" json:"address,omitempty"`
	LogoURL     string            `gorm:"type:varchar(500)" json:"logo_url,omitempty"`
	AdminUserID *uuid.UUID        `gorm:"type:char(36);index" json:"admin_user_id,omitempty"`
	Status      OrganizationStatus `gorm:"type:varchar(20);not null;default:'active'" json:"status"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	DeletedAt   gorm.DeletedAt     `gorm:"index" json:"-"`

	// Relationships
	AdminUser *User `gorm:"foreignKey:AdminUserID" json:"admin_user,omitempty"`
	Teams     []Team `gorm:"foreignKey:OrganizationID" json:"teams,omitempty"`
}

// BeforeCreate hook to generate UUID
func (o *Organization) BeforeCreate(tx *gorm.DB) error {
	if o.ID == uuid.Nil {
		o.ID = uuid.New()
	}
	return nil
}

// TableName specifies the table name
func (Organization) TableName() string {
	return "organizations"
}

