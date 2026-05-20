package dto

type UserProfileDTO struct {
	UserID      string              `json:"user_id"`
	Name        string              `json:"name"`
	Email       string              `json:"email"`
	Phone       string              `json:"phone"`
	AvatarURL   string              `json:"avatar_url"`
	IsActive    bool                `json:"is_active"`
	Roles       []UserRoleDTO       `json:"roles"`
	Permissions []string            `json:"permissions"`
	CreatedAt   string              `json:"created_at"`
	UpdatedAt   string              `json:"updated_at"`
}
