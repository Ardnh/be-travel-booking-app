package entities

import (
	"time"

	"github.com/google/uuid"
)

// 1. Global role — tidak terikat vendor/pool
type UserRoles struct {
	UserRoleID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID     uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
	Role       string    `gorm:"type:varchar(50);not null"` // daily_user, admin_platform, platform_owner
	CreatedAt  time.Time `gorm:"autoCreateTime"`

	User Users `gorm:"foreignKey:UserID;references:UserID"`
}

// 2. Vendor membership — siapa pegang vendor apa
type VendorMembers struct {
	VendorMemberID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID         uuid.UUID `gorm:"type:uuid;not null;index"`
	VendorID       uuid.UUID `gorm:"type:uuid;not null;index"`
	Role           string    `gorm:"type:varchar(50);not null"` // owner, admin, staff
	CreatedAt      time.Time `gorm:"autoCreateTime"`

	User   Users   `gorm:"foreignKey:UserID;references:UserID"`
	Vendor Vendors `gorm:"foreignKey:VendorID;references:VendorID"`
}

// 3. Pool membership — siapa pegang pool apa
type PoolMembers struct {
	PoolMemberID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID       uuid.UUID `gorm:"type:uuid;not null;index"`
	PoolID       uuid.UUID `gorm:"type:uuid;not null;index"`
	VendorID     uuid.UUID `gorm:"type:uuid;not null;index"`  // denormalisasi
	Role         string    `gorm:"type:varchar(50);not null"` // admin, operator, cashier
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	User   Users   `gorm:"foreignKey:UserID;references:UserID"`
	Pool   Pools   `gorm:"foreignKey:PoolID;references:PoolID"`
	Vendor Vendors `gorm:"foreignKey:VendorID;references:VendorID"`
}
