package repositories

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type PoolPointRepository interface {
	GetPoolPointByID(ctx context.Context, poolID uuid.UUID) (*entities.PoolPoint, error)
	GetAllPoolPoints(ctx context.Context) ([]entities.PoolPoint, error)
	GetPoolPointsByVendorID(ctx context.Context, vendorID uuid.UUID) ([]entities.PoolPoint, error)
	CreatePoolPoint(ctx context.Context, poolPoint entities.PoolPoint) error
	UpdatePoolPoint(ctx context.Context, poolPoint entities.PoolPoint) error
	DeletePoolPoint(ctx context.Context, poolID uuid.UUID) error
}