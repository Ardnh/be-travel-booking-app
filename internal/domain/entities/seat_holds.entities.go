package entities

import (
	"time"

	"github.com/google/uuid"
)

// ============================================================
// SeatHold
// ============================================================
type SeatHolds struct {
	HoldID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"hold_id"`
	ScheduleID       uuid.UUID `gorm:"type:uuid;not null;index" json:"schedule_id"`
	LayoutPositionID uuid.UUID `gorm:"type:uuid;not null;index" json:"layout_position_id"`
	UserID           uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	ExpireAt         time.Time `gorm:"type:timestamp;not null" json:"expire_at"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`

	// // Relations
	Schedule       Schedules
	LayoutPosition LayoutPositions
	User           Users
}

func (SeatHolds) TableName() string { return "seat_holds" }
