package entities

import (
	"github.com/google/uuid"
)

// ============================================================
// LayoutPosition
// ============================================================

type LayoutPositions struct {
	LayoutPositionID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"layout_position_id"`
	LayoutID         uuid.UUID `gorm:"type:uuid;not null;index" json:"layout_id"`
	Label            string    `gorm:"type:varchar(50);not null" json:"label"`
	Row              int       `gorm:"type:int;not null" json:"row"`
	Col              int       `gorm:"type:int;not null" json:"col"`
	PositionType     string    `gorm:"type:varchar(50);not null" json:"position_type"`
	IsUsed           bool      `gorm:"default:true" json:"is_used"`

	// // Relations
	Layout Layouts
	// BookingSeats []BookingSeats `gorm:"-" json:"booking_seats,omitempty"`
	// SeatHolds    []SeatHolds    `gorm:"-" json:"seat_holds,omitempty"`

	BaseModel
}

func (LayoutPositions) TableName() string { return "layout_positions" }
