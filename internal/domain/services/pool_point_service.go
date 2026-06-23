package services

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type PoolPointService interface {
	GetPoolPointByID(ctx context.Context, poolID uuid.UUID) (*entities.Pools, error)
	GetAllPoolPoints(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.Pools, int64, error)
	GetPoolPointsByVendorID(ctx context.Context, vendorID uuid.UUID, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.Pools, int64, error)
	GetAvailableLocationsByVendorID(ctx context.Context, vendorID uuid.UUID, locationType string) ([]dto.AvailableLocationDTO, error)
	CreatePoolPoint(ctx context.Context, req dto.CreatePoolsDTO) (*entities.Pools, error)
	UpdatePoolPoint(ctx context.Context, poolID uuid.UUID, req dto.UpdatePoolsDTO) (*entities.Pools, error)
	DeletePoolPoint(ctx context.Context, poolID uuid.UUID) error
}
