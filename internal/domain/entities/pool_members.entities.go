package entities

import (
	"time"

	"github.com/google/uuid"
)

type PoolMembers struct {
	PoolMemberID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID       uuid.UUID `gorm:"type:uuid;not null;index"`
	PoolID       uuid.UUID `gorm:"type:uuid;not null;index"`
	VendorID     uuid.UUID `gorm:"type:uuid;not null;index"`  // denormalisasi
	Role         string    `gorm:"type:varchar(50);not null"` // admin, operator, cashier
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	User   Users
	Pool   Pools
	Vendor Vendors
}

func (PoolMembers) TableName() string { return "pool_members" }
