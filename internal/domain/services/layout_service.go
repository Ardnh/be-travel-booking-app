package services

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
)

type LayoutService interface {
	GetLayoutById(ctx context.Context, layoutID string) (*dto.LayoutDTO, error)
	GetLayout(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]*dto.LayoutDTO, int64, error)
	CreateLayout(ctx context.Context, layout dto.CreateLayoutDTO) error
	UpdateLayout(ctx context.Context, layoutID string, layout dto.CreateLayoutDTO) error
	DeleteLayout(ctx context.Context, layoutID string) error
}
