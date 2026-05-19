package entities

import (
	"github.com/google/uuid"
)

// ============================================================
// User
// ============================================================

type Users struct {
	UserID       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"user_id"`
	Name         string    `gorm:"type:varchar(255);not null" json:"name"`
	Email        string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"-"`
	Phone        string    `gorm:"type:varchar(50)" json:"phone"`
	AvatarURL    string    `gorm:"type:varchar(500)" json:"avatar_url"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`

	// Relations
	// Roles []UserRoles
	// Bookings []Bookings
	// SeatHolds []SeatHolds

	BaseModel
}

func (Users) TableName() string { return "users" }
