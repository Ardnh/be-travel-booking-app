package entities

import (
	"github.com/google/uuid"
)

// ============================================================
// Vendor
// ============================================================
type Vendors struct {
	VendorID            uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"vendor_id"`
	BusinessName        string    `gorm:"type:varchar(255);not null" json:"business_name"`
	OwnerName           string    `gorm:"type:varchar(255);not null" json:"owner_name"`
	Description         string    `gorm:"type:text" json:"description"`
	FoundedYear         int       `gorm:"type:int" json:"founded_year"`
	PhoneNumber         string    `gorm:"type:varchar(50)" json:"phone_number"`
	Email               string    `gorm:"type:varchar(255)" json:"email"`
	HeadOfficeAddress   string    `gorm:"type:text" json:"head_office_address"`
	LogoURL             string    `gorm:"type:text" json:"logo_url"`
	BannerURL           string    `gorm:"type:text" json:"banner_url"`
	LegalDocumentNumber string    `gorm:"type:varchar(255)" json:"legal_document_number"`
	IsVerified          bool      `gorm:"default:false" json:"is_verified"`
	Status              string    `gorm:"type:varchar(50);default:'active'" json:"status"`

	// // Relations
	// Pools []Pools
	// Schedules []Schedules `gorm:"-" json:"schedules,omitempty"`

	BaseModel
}

func (Vendors) TableName() string { return "vendors" }
