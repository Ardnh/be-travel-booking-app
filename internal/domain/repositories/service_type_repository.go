package repositories

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type ServiceTypeRepository interface {
	GetServiceTypeByID(ctx context.Context, serviceTypeID uuid.UUID) (*entities.ServiceType, error)
	GetAllServiceTypes(ctx context.Context) ([]entities.ServiceType, error)
	CreateServiceType(ctx context.Context, serviceType entities.ServiceType) error
	UpdateServiceType(ctx context.Context, serviceType entities.ServiceType) error
	DeleteServiceType(ctx context.Context, serviceTypeID uuid.UUID) error
}
