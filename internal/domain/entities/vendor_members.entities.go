package entities

import (
	"time"

	"github.com/google/uuid"
)

type VendorMembers struct {
	VendorMemberID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID         uuid.UUID `gorm:"type:uuid;not null;index"`
	VendorID       uuid.UUID `gorm:"type:uuid;not null;index"`
	Role           string    `gorm:"type:varchar(50);not null"` // owner, admin, staff
	CreatedAt      time.Time `gorm:"autoCreateTime"`

	User   Users   `gorm:"foreignKey:UserID;references:UserID"`
	Vendor Vendors `gorm:"foreignKey:VendorID;references:VendorID"`
}

func (VendorMembers) TableName() string { return "vendor_members" }
