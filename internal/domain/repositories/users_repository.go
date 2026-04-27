package repositories

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.Users, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*entities.Users, error)
	CreateUser(ctx context.Context, user *entities.Users) error
	UpdateUser(ctx context.Context, user *entities.Users) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}
