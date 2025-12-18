package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleSuperAdmin   UserRole = "super_admin"
	RoleOrgAdmin     UserRole = "org_admin"
	RoleTeamMember   UserRole = "team_member"
	RolePublic       UserRole = "public"
)

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
)

type User struct {
	ID             uuid.UUID  `gorm:"type:char(36);primary_key" json:"id"`
	Email          string     `gorm:"uniqueIndex;not null" json:"email"`
	Password       string     `gorm:"not null" json:"-"` // Don't return password in JSON
	Role           UserRole   `gorm:"type:varchar(20);not null;default:'public'" json:"role"`
	OrganizationID *uuid.UUID `gorm:"type:char(36);index" json:"organization_id,omitempty"`
	FullName       string     `gorm:"type:varchar(255)" json:"full_name"`
	Phone          string     `gorm:"type:varchar(20)" json:"phone"`
	ProfileImageURL string    `gorm:"type:varchar(500)" json:"profile_image_url,omitempty"`
	Status         UserStatus `gorm:"type:varchar(20);not null;default:'active'" json:"status"`
	LastLoginAt    *time.Time `json:"last_login_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Organization *Organization `gorm:"foreignKey:OrganizationID" json:"organization,omitempty"`
}

// BeforeCreate hook to generate UUID
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

// TableName specifies the table name
func (User) TableName() string {
	return "users"
}

// IsAdmin checks if user is super admin
func (u *User) IsAdmin() bool {
	return u.Role == RoleSuperAdmin
}

// IsOrgAdmin checks if user is organization admin
func (u *User) IsOrgAdmin() bool {
	return u.Role == RoleOrgAdmin
}

// CanManageOrganization checks if user can manage an organization
func (u *User) CanManageOrganization(orgID uuid.UUID) bool {
	if u.IsAdmin() {
		return true
	}
	if u.IsOrgAdmin() && u.OrganizationID != nil && *u.OrganizationID == orgID {
		return true
	}
	return false
}

// IsActive checks if user is active
func (u *User) IsActive() bool {
	return u.Status == UserStatusActive
}

