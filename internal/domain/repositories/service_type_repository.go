package repositories

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type ServiceTypeRepository interface {
	GetServiceTypeByID(ctx context.Context, serviceTypeID uuid.UUID) (*entities.ServiceTypes, error)
	GetAllServiceTypes(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.ServiceTypes, int64, error)
	CreateServiceType(ctx context.Context, serviceType entities.ServiceTypes) (*entities.ServiceTypes, error)
	UpdateServiceType(ctx context.Context, serviceType entities.ServiceTypes) (*entities.ServiceTypes, error)
	DeleteServiceType(ctx context.Context, serviceTypeID uuid.UUID) error
}
