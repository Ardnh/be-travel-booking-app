package repositories

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type LayoutRepository interface {
	GetLayoutById(ctx context.Context, layoutID uuid.UUID) (*entities.Layout, error)
	GetLayout(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]*entities.Layout, int64, error)
	CreateLayout(ctx context.Context, layout *entities.Layout, layoutPosition []*entities.LayoutPosition) error
	UpdateLayout(ctx context.Context, layout *entities.Layout, layoutPosition []*entities.LayoutPosition) error
	DeleteLayout(ctx context.Context, layoutID uuid.UUID) error
}
