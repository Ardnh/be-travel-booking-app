package services

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type PoolPointService interface {
	GetPoolPointByID(ctx context.Context, poolID uuid.UUID) (*entities.Pools, error)
	GetAllPoolPoints(ctx context.Context) ([]entities.Pools, error)
	GetPoolPointsByVendorID(ctx context.Context, vendorID uuid.UUID) ([]entities.Pools, error)
	CreatePoolPoint(ctx context.Context, req dto.CreatePoolsDTO) error
	UpdatePoolPoint(ctx context.Context, poolID uuid.UUID, req dto.UpdatePoolsDTO) error
	DeletePoolPoint(ctx context.Context, poolID uuid.UUID) error
}
