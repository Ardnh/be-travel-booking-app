package entities

import "github.com/google/uuid"

type UsersRole struct {
	RoleID   uuid.UUID `gorm:"column:role_id;type:uuid;primaryKey"`
	UserID   uuid.UUID `gorm:"column:user_id;type:uuid;not null"`
	RoleName string    `gorm:"column:role_name;type:varchar(255);not null"`
}

func (UsersRole) TableName() string {
	return "users_roles"
}
