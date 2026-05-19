package services

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type LayoutPositionService interface {
	GetLayoutPositionByID(ctx context.Context, layoutPositionID uuid.UUID) (*entities.LayoutPositions, error)
	GetLayoutPositionsByLayoutID(ctx context.Context, layoutID uuid.UUID) ([]entities.LayoutPositions, error)
	GetAllLayoutPositions(ctx context.Context) ([]entities.LayoutPositions, error)
	CreateLayoutPosition(ctx context.Context, req dto.CreateLayoutPositionDTO, layoutID uuid.UUID) error
	UpdateLayoutPosition(ctx context.Context, layoutPositionID uuid.UUID, req dto.UpdateLayoutPositionDTO) error
	DeleteLayoutPosition(ctx context.Context, layoutPositionID uuid.UUID) error
}
