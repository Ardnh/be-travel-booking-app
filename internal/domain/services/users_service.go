package services

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
)

type UsersService interface {
	CreateUser(ctx context.Context, user dto.CreateUserDTO) error
	UpdateUser(ctx context.Context, userId string, req dto.UpdateUserDTO) error
	DeleteUser(ctx context.Context, userID string) error
	GetUserProfile(ctx context.Context, userID string) (*dto.UserProfileDTO, error)
}
