package entities

import (
	"time"

	"github.com/google/uuid"
)

// ============================================================
// UserRole
// ============================================================
type UserRoles struct {
	UserRoleID uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"user_role_id"`
	UserID     uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Role       string     `gorm:"type:varchar(50);not null" json:"role"`
	VendorID   *uuid.UUID `gorm:"type:uuid;index" json:"vendor_id,omitempty"`
	PoolID     *uuid.UUID `gorm:"type:uuid;index" json:"pool_id,omitempty"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`

	// Relations
	User Users
	// Vendor *Vendors
	// Pool   *Pools
}

func (UserRoles) TableName() string { return "user_roles" }
