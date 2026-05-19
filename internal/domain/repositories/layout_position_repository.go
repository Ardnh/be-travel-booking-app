package repositories

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type LayoutPositionRepository interface {
	GetLayoutPositionByID(ctx context.Context, layoutPositionID uuid.UUID) (*entities.LayoutPositions, error)
	GetLayoutPositionsByLayoutID(ctx context.Context, layoutID uuid.UUID) ([]entities.LayoutPositions, error)
	GetAllLayoutPositions(ctx context.Context) ([]entities.LayoutPositions, error)
	CreateLayoutPosition(ctx context.Context, layoutPosition entities.LayoutPositions) error
	UpdateLayoutPosition(ctx context.Context, layoutPosition entities.LayoutPositions) error
	DeleteLayoutPosition(ctx context.Context, layoutPositionID uuid.UUID) error
}
