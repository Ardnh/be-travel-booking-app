package entities

import (
	"time"

	"github.com/google/uuid"
)

type LayoutPosition struct {
	LayoutPositionID uuid.UUID `gorm:"column:layout_position_id;type:uuid;primaryKey"`
	LayoutID         uuid.UUID `gorm:"column:layout_id;type:uuid;not null;foreignKey:LayoutID;references:LayoutID"`
	Label            string    `gorm:"column:label;type:varchar;not null"`
	Row              int       `gorm:"column:row;type:int;not null"`
	Col              int       `gorm:"column:col;type:int;not null"`
	PositionType     string    `gorm:"column:position_type;type:varchar;default:'seat'"`
	IsUsed           bool      `gorm:"column:is_used;type:boolean;default:true"`
	CreatedAt        time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;autoCreateTime:milli"`
	UpdatedAt        time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime:milli"`
}

func (LayoutPosition) TableName() string {
	return "layout_positions"
}
