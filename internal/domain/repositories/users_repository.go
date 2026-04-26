package repositories

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.Users, error)
}
