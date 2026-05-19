package entities

import (
	"time"

	"gorm.io/gorm"
)

// ============================================================
// Base
// ============================================================

type BaseModel struct {
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
