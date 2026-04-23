package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `gorm:"column:user_id;type:uuid;primaryKey"`
	Name      string    `gorm:"column:name;type:varchar;not null"`
	Email     string    `gorm:"column:email;type:varchar;not null;uniqueIndex"`
	Password  string    `gorm:"column:password;type:varchar;not null"`
	Phone     string    `gorm:"column:phone;type:varchar;not null"`
	AvatarURL string    `gorm:"column:avatar_url;type:varchar;not null"`
	IsActive  bool      `gorm:"column:is_active;type:boolean;not null;default:true"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime:milli"`
}

func (User) TableName() string {
	return "users"
}
