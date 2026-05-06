package services

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type ServiceTypeService interface {
	GetServiceTypeByID(ctx context.Context, serviceTypeID uuid.UUID) (*entities.ServiceType, error)
	GetAllServiceTypes(ctx context.Context) ([]entities.ServiceType, error)
	CreateServiceType(ctx context.Context, req dto.CreateServiceTypeDTO, createdBy uuid.UUID) error
	UpdateServiceType(ctx context.Context, serviceTypeID uuid.UUID, req dto.UpdateServiceTypeDTO) error
	DeleteServiceType(ctx context.Context, serviceTypeID uuid.UUID) error
}