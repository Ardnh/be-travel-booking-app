package entities

import (
	"time"

	"github.com/google/uuid"
)

type Vendor struct {
	VendorID             uuid.UUID  `gorm:"column:vendor_id;type:uuid;primaryKey"`
	BusinessName         string     `gorm:"column:business_name;type:varchar;not null"`
	OwnerName            string     `gorm:"column:owner_name;type:varchar;not null"`
	Description          *string    `gorm:"column:description;type:text;null"`
	FoundedYear          *int       `gorm:"column:founded_year;type:int;null"`
	PhoneNumber          string     `gorm:"column:phone_number;type:varchar;not null"`
	Email                string     `gorm:"column:email;type:varchar;not null;uniqueIndex"`
	HeadOfficeAddress    string     `gorm:"column:head_office_address;type:text;not null"`
	LogoURL              *string    `gorm:"column:logo_url;type:text;null"`
	BannerURL            *string    `gorm:"column:banner_url;type:text;null"`
	LegalDocumentNumber  *string    `gorm:"column:legal_document_number;type:varchar;null"`
	IsVerified           bool       `gorm:"column:is_verified;type:boolean;default:false"`
	Status               string     `gorm:"column:status;type:varchar;default:'pending'"`
	CreatedAt            time.Time  `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;autoCreateTime:milli"`
	UpdatedAt            time.Time  `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime:milli"`
}

func (Vendor) TableName() string {
	return "vendors"
}
