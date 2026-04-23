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
}

func (Layout) TableName() string {
	return "layouts"
}
