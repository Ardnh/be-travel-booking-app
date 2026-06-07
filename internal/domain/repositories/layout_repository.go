package repositories

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type LayoutRepository interface {
	GetLayoutById(ctx context.Context, layoutID uuid.UUID) (*entities.Layouts, error)
	GetLayout(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]*entities.Layouts, int64, error)
	CreateLayout(ctx context.Context, layout entities.Layouts) (*entities.Layouts, error)
	UpdateLayout(ctx context.Context, layout entities.Layouts) (*entities.Layouts, error)
	DeleteLayout(ctx context.Context, layoutID uuid.UUID) error
}
