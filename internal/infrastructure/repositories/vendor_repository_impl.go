package repositories

import (
	"context"
	"errors"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type vendorRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewVendorRepository(db *gorm.DB, redis *redis.Client) repositories.VendorRepository {
	return &vendorRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *vendorRepositoryImpl) GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*entities.Vendors, error) {
	var vendor entities.Vendors
	err := r.db.WithContext(ctx).Where("vendor_id = ?", vendorID).First(&vendor).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return &vendor, nil
}

func (r *vendorRepositoryImpl) GetVendorByOwnerUserID(ctx context.Context, ownerUserID uuid.UUID) (*entities.Vendors, error) {
	var vendor entities.Vendors
	err := r.db.WithContext(ctx).Where("owner_user_id = ?", ownerUserID).First(&vendor).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return &vendor, nil
}

func (r *vendorRepositoryImpl) GetAllVendors(ctx context.Context) ([]entities.Vendors, error) {
	var vendors []entities.Vendors
	err := r.db.WithContext(ctx).Find(&vendors).Error
	if err != nil {
		return nil, err
	}
	return vendors, nil
}

// Repository
func (r *vendorRepositoryImpl) CreateVendor(ctx context.Context, vendor entities.Vendors) error {
	if err := r.db.WithContext(ctx).Create(&vendor).Error; err != nil {
		return err
	}
	return nil
}

func (r *vendorRepositoryImpl) UpdateVendor(ctx context.Context, vendor entities.Vendors) error {
	err := r.db.WithContext(ctx).Save(&vendor).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *vendorRepositoryImpl) DeleteVendor(ctx context.Context, vendorID uuid.UUID) error {
	err := r.db.WithContext(ctx).Where("vendor_id = ?", vendorID).Delete(&entities.Vendors{}).Error
	if err != nil {
		return err
	}
	return nil
}
