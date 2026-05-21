package services

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type VendorService interface {
	GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*entities.Vendors, error)
	GetVendorByOwnerUserID(ctx context.Context, ownerUserID uuid.UUID) (*entities.Vendors, error)
	GetAllVendors(ctx context.Context) ([]entities.Vendors, error)
	CreateVendor(ctx context.Context, req dto.CreateVendorDTO) error
	UpdateVendor(ctx context.Context, vendorID uuid.UUID, req dto.UpdateVendorDTO) error
	DeleteVendor(ctx context.Context, vendorID uuid.UUID) error
}
