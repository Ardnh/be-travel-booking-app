package entities

import (
	"github.com/google/uuid"
)

// ============================================================
// ServiceTypes
// ============================================================
type ServiceTypes struct {
	ServiceTypeID      uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"service_type_id"`
	Name               string    `gorm:"type:varchar(255);not null" json:"name"`
	UniqueCode         string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"unique_code"`
	Description        string    `gorm:"type:text" json:"description"`
	NeedChair          bool      `gorm:"default:false" json:"need_chair"`
	NeedPickupAddress  bool      `gorm:"default:false" json:"need_pickup_address"`
	NeedDropoffAddress bool      `gorm:"default:false" json:"need_dropoff_address"`
	DisplayOrder       int       `gorm:"type:int;default:0" json:"display_order"`
	Status             bool      `gorm:"default:true" json:"status"`
	CreatedBy          uuid.UUID `gorm:"type:uuid" json:"created_by,omitempty"`

	// Relations

	BaseModel
}

func (ServiceTypes) TableName() string { return "service_types" }
