package entities

import (
	"time"

	"github.com/google/uuid"
)

type ServiceType struct {
	ServiceTypeID      uuid.UUID `gorm:"column:service_type_id;type:uuid;primaryKey"`
	Name               string    `gorm:"column:name;type:varchar;not null"`
	UniqueCode         string    `gorm:"column:unique_code;type:varchar;not null;uniqueIndex"`
	Description        string    `gorm:"column:description;type:text;null"`
	NeedChair          bool      `gorm:"column:need_chair;type:boolean;default:false"`
	NeedPickupAddress  bool      `gorm:"column:need_pickup_address;type:boolean;default:false"`
	NeedDropoffAddress bool      `gorm:"column:need_dropoff_address;type:boolean;default:false"`
	DisplayOrder       int       `gorm:"column:display_order;type:int;default:0"`
	Status             bool      `gorm:"column:status;type:boolean;default:true"`

	// Kolom di database
	CreatedBy uuid.UUID `gorm:"column:created_by;type:uuid;not null"`

	// Relasi → ini yang bikin FK constraint
	Creator Users `gorm:"foreignKey:CreatedBy;references:UserID"`

	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime:milli"`
}

func (ServiceType) TableName() string {
	return "service_types"
}
