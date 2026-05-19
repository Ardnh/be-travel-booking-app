package entities

import (
	"time"

	"github.com/google/uuid"
)

// ============================================================
// Booking
// ============================================================
type Bookings struct {
	BookingID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"booking_id"`
	BookingCode      string     `gorm:"type:varchar(50);uniqueIndex;not null" json:"booking_code"`
	UserID           uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	ScheduleID       uuid.UUID  `gorm:"type:uuid;not null;index" json:"schedule_id"`
	PassengerName    string     `gorm:"type:varchar(255);not null" json:"passenger_name"`
	PassengerPhone   string     `gorm:"type:varchar(50)" json:"passenger_phone"`
	PickupAddress    string     `gorm:"type:text" json:"pickup_address"`
	DropoffAddress   string     `gorm:"type:text" json:"dropoff_address"`
	PricePerSeat     float64    `gorm:"type:decimal(12,2);not null" json:"price_per_seat"`
	SeatCount        int        `gorm:"type:int;not null" json:"seat_count"`
	TotalPrice       float64    `gorm:"type:decimal(12,2);not null" json:"total_price"`
	PaymentStatus    string     `gorm:"type:varchar(50);default:'unpaid'" json:"payment_status"`
	PaymentReference string     `gorm:"type:varchar(255)" json:"payment_reference"`
	PaymentMethod    string     `gorm:"type:varchar(50)" json:"payment_method"`
	PaidAt           *time.Time `gorm:"type:timestamp" json:"paid_at,omitempty"`
	BookingStatus    string     `gorm:"type:varchar(50);default:'pending'" json:"booking_status"`
	CheckinStatus    string     `gorm:"type:varchar(50)" json:"checkin_status"`
	CheckinAt        *time.Time `gorm:"type:timestamp" json:"checkin_at,omitempty"`
	CheckinBy        *uuid.UUID `gorm:"type:uuid" json:"checkin_by,omitempty"`
	Source           string     `gorm:"type:varchar(50)" json:"source"`
	CreatedBy        *uuid.UUID `gorm:"type:uuid" json:"created_by,omitempty"`
	Notes            string     `gorm:"type:text" json:"notes"`
	CancelledAt      *time.Time `gorm:"type:timestamp" json:"cancelled_at,omitempty"`
	CancelReason     string     `gorm:"type:text" json:"cancel_reason"`

	// Relations — belongs to tetap
	User     Users
	Schedule Schedules
	// CheckinByUser *Users `gorm:"foreignKey:CheckinBy;references:UserID" json:"checkin_by_user,omitempty"`
	// CreatedByUser *Users `gorm:"foreignKey:CreatedByUserID;references:UserID" json:"created_by_user,omitempty"`
	// // Relations — has many → gorm:"-"
	// BookingSeats []BookingSeats `gorm:"-" json:"booking_seats,omitempty"`
	// Payments     []Payments     `gorm:"-" json:"payments,omitempty"`

	BaseModel
}

func (Bookings) TableName() string { return "bookings" }
