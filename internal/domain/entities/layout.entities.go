package entities

import (
	"time"

	"github.com/google/uuid"
)

type Layout struct {
	LayoutID  uuid.UUID `gorm:"column:layout_id;type:uuid;primaryKey"`
	Name      string    `gorm:"column:name;type:varchar;not null"`
	GridSizeX int       `gorm:"column:grid_size_x;type:int;not null"`
	GridSizeY int       `gorm:"column:grid_size_y;type:int;not null"`
	SeatCount int       `gorm:"column:seat_count;type:int;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime:milli"`
	CreatedBy uuid.UUID `gorm:"column:created_by;type:uuid;not null"`

	// Relasi → ini yang bikin FK constraint
	Creator Users `gorm:"foreignKey:CreatedBy;references:UserID"`

	// Has-many → FK dibuat di tabel layout_positions
	LayoutPositions []LayoutPosition `gorm:"foreignKey:LayoutID;references:LayoutID;constraint:OnDelete:CASCADE;"`
}

func (Layout) TableName() string {
	return "layouts"
}
