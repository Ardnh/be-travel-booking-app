package services

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
)

type UserRolesService interface {
	GetUserRoleByID(ctx context.Context, userRoleID string) (*dto.UserRoleDTO, error)
	GetUserRolesByUserID(ctx context.Context, userID string) ([]dto.UserRoleDTO, error)
	GetUserRolesByVendorID(ctx context.Context, vendorID string) ([]dto.UserRoleDTO, error)
	CreateUserRole(ctx context.Context, req dto.CreateUserRoleDTO) (*dto.UserRoleDTO, error)
	UpdateUserRole(ctx context.Context, userRoleID string, req dto.UpdateUserRoleDTO) (*dto.UserRoleDTO, error)
	DeleteUserRole(ctx context.Context, userRoleID string) error
}