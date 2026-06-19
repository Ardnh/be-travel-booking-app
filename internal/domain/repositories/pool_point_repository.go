package repositories

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type PoolPointRepository interface {
	GetPoolPointByID(ctx context.Context, poolID uuid.UUID) (*entities.Pools, error)
	GetAllPoolPoints(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.Pools, int64, error)
	GetPoolPointsByVendorID(ctx context.Context, vendorID uuid.UUID, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.Pools, int64, error)
	GetAvailableLocationsByVendorID(ctx context.Context, vendorID uuid.UUID, locationType string) ([]string, error)
	CreatePoolPoint(ctx context.Context, poolPoint entities.Pools) (*entities.Pools, error)
	UpdatePoolPoint(ctx context.Context, poolPoint entities.Pools) (*entities.Pools, error)
	DeletePoolPoint(ctx context.Context, poolID uuid.UUID) error
}
