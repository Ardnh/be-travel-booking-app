package entities

import (
	"github.com/google/uuid"
)

// ============================================================
// Pool
// ============================================================

type Pools struct {
	PoolID      uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"pool_id"`
	VendorID    uuid.UUID `gorm:"type:uuid;not null;index" json:"vendor_id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Slug        string    `gorm:"type:varchar(255);uniqueIndex" json:"slug"`
	Address     string    `gorm:"type:text" json:"address"`
	City        string    `gorm:"type:varchar(100)" json:"city"`
	Province    string    `gorm:"type:varchar(100)" json:"province"`
	Latitude    float64   `gorm:"type:decimal(10,7)" json:"latitude"`
	Longitude   float64   `gorm:"type:decimal(10,7)" json:"longitude"`
	OpenTime    string    `gorm:"type:time" json:"open_time"`
	CloseTime   string    `gorm:"type:time" json:"close_time"`
	Status      string    `gorm:"type:varchar(50);default:'active'" json:"status"`
	Description string    `gorm:"type:text" json:"description"`

	// Relations
	Vendor Vendors

	BaseModel
}

func (Pools) TableName() string { return "pools" }
