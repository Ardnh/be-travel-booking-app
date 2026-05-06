package services

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type PoolPointService interface {
	GetPoolPointByID(ctx context.Context, poolID uuid.UUID) (*entities.PoolPoint, error)
	GetAllPoolPoints(ctx context.Context) ([]entities.PoolPoint, error)
	GetPoolPointsByVendorID(ctx context.Context, vendorID uuid.UUID) ([]entities.PoolPoint, error)
	CreatePoolPoint(ctx context.Context, req dto.CreatePoolPointDTO) error
	UpdatePoolPoint(ctx context.Context, poolID uuid.UUID, req dto.UpdatePoolPointDTO) error
	DeletePoolPoint(ctx context.Context, poolID uuid.UUID) error
}