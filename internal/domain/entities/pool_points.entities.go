package entities

import (
	"time"

	"github.com/google/uuid"
)

type PoolPoint struct {
	PoolID      uuid.UUID `gorm:"column:pool_id;type:uuid;primaryKey"`
	VendorID    uuid.UUID `gorm:"column:vendor_id;type:uuid;not null;foreignKey:VendorID;references:VendorID"`
	Name        string    `gorm:"column:name;type:varchar;not null"`
	Slug        string    `gorm:"column:slug;type:varchar;uniqueIndex"`
	Address     string    `gorm:"column:address;type:text;not null"`
	City        string    `gorm:"column:city;type:varchar;not null"`
	Province    string    `gorm:"column:province;type:varchar;not null"`
	Latitude    float64   `gorm:"column:latitude;type:decimal(10,8);not null"`
	Longitude   float64   `gorm:"column:longitude;type:decimal(10,8);not null"`
	OpenTime    string    `gorm:"column:open_time;type:time;not null"`
	CloseTime   string    `gorm:"column:close_time;type:time;not null"`
	Status      string    `gorm:"column:status;type:varchar;default:'active'"`
	Description *string   `gorm:"column:description;type:text;null"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime:milli"`
}

func (PoolPoint) TableName() string {
	return "pool_points"
}
