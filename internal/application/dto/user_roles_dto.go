package dto

type UserRoleDTO struct {
	UserRoleID string  `json:"user_role_id"`
	UserID     string  `json:"user_id"`
	Role       string  `json:"role"`
	VendorID   *string `json:"vendor_id,omitempty"`
	PoolID     *string `json:"pool_id,omitempty"`
}

type CreateUserRoleDTO struct {
	UserID   string  `json:"user_id" validate:"required"`
	Role     string  `json:"role" validate:"required"`
	VendorID *string `json:"vendor_id,omitempty"`
	PoolID   *string `json:"pool_id,omitempty"`
}

type UpdateUserRoleDTO struct {
	Role     *string `json:"role,omitempty"`
	VendorID *string `json:"vendor_id,omitempty"`
	PoolID   *string `json:"pool_id,omitempty"`
}