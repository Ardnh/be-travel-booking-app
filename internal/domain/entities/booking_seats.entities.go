package entities

import (
	"time"

	"github.com/google/uuid"
)

// ============================================================
// BookingSeat
// ============================================================
type BookingSeats struct {
	SeatID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"seat_id"`
	BookingID        uuid.UUID `gorm:"type:uuid;not null;index" json:"booking_id"`
	ScheduleID       uuid.UUID `gorm:"type:uuid;not null;index" json:"schedule_id"`
	LayoutPositionID uuid.UUID `gorm:"type:uuid;not null;index" json:"layout_position_id"`
	SeatLabel        string    `gorm:"type:varchar(50);not null" json:"seat_label"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relations
	Booking        Bookings
	Schedule       Schedules
	LayoutPosition LayoutPositions
}

func (BookingSeats) TableName() string { return "booking_seats" }
