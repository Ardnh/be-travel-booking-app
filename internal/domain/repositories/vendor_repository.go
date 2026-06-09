package repositories

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type VendorRepository interface {
	GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*entities.Vendors, error)
	GetVendorByOwnerUserID(ctx context.Context, ownerUserID uuid.UUID) (*entities.Vendors, error)
	GetAllVendors(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.Vendors, int64, error)
	CreateVendor(ctx context.Context, vendor entities.Vendors) error
	UpdateVendor(ctx context.Context, vendor entities.Vendors) error
	DeleteVendor(ctx context.Context, vendorID uuid.UUID) error
}
