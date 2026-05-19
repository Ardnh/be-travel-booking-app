package entities

import (
	"time"

	"github.com/google/uuid"
)

// ============================================================
// Payment
// ============================================================

type Payments struct {
	PaymentID uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"payment_id"`
	BookingID uuid.UUID  `gorm:"type:uuid;not null;index" json:"booking_id"`
	Amount    float64    `gorm:"type:decimal(12,2);not null" json:"amount"`
	Method    string     `gorm:"type:varchar(50);not null" json:"method"`
	Status    string     `gorm:"type:varchar(50);default:'pending'" json:"status"`
	Reference string     `gorm:"type:varchar(255)" json:"reference"`
	PaidAt    *time.Time `gorm:"type:timestamp" json:"paid_at,omitempty"`
	Notes     string     `gorm:"type:text" json:"notes"`

	// // Relations
	Booking Bookings

	BaseModel
}

func (Payments) TableName() string { return "payments" }
