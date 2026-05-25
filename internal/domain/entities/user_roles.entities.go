package entities

// ============================================================
// UserRole
// ============================================================
// type UserRoles struct {
// 	UserRoleID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
// 	UserID     uuid.UUID `gorm:"type:uuid;not null;uniqueIndex"`
// 	Role       string    `gorm:"type:varchar(50);not null"` // daily_user, admin_platform, platform_owner
// 	CreatedAt  time.Time `gorm:"autoCreateTime"`

// 	User Users
// }

// func (UserRoles) TableName() string { return "user_roles" }
