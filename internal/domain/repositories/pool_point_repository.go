package repositories

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type PoolPointRepository interface {
	GetPoolPointByID(ctx context.Context, poolID uuid.UUID) (*entities.Pools, error)
	GetAllPoolPoints(ctx context.Context) ([]entities.Pools, error)
	GetPoolPointsByVendorID(ctx context.Context, vendorID uuid.UUID) ([]entities.Pools, error)
	CreatePoolPoint(ctx context.Context, poolPoint entities.Pools) error
	UpdatePoolPoint(ctx context.Context, poolPoint entities.Pools) error
	DeletePoolPoint(ctx context.Context, poolID uuid.UUID) error
}
