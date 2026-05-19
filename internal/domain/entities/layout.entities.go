package entities

import (
	"github.com/google/uuid"
)

// ============================================================
// Layout
// ============================================================
type Layouts struct {
	LayoutID  uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"layout_id"`
	Name      string     `gorm:"type:varchar(255);not null" json:"name"`
	GridSizeX int        `gorm:"type:int;not null" json:"grid_size_x"`
	GridSizeY int        `gorm:"type:int;not null" json:"grid_size_y"`
	SeatCount int        `gorm:"type:int;not null" json:"seat_count"`
	CreatedBy *uuid.UUID `gorm:"type:uuid" json:"created_by,omitempty"`

	// // Relations
	// Positions []LayoutPositions `gorm:"-" json:"positions,omitempty"`
	// Schedules []Schedules       `gorm:"-" json:"schedules,omitempty"`

	BaseModel
}

func (Layouts) TableName() string { return "layouts" }
