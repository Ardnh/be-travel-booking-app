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

type serviceTypeRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewServiceTypeRepository(db *gorm.DB, redis *redis.Client) repositories.ServiceTypeRepository {
	return &serviceTypeRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *serviceTypeRepositoryImpl) GetServiceTypeByID(ctx context.Context, serviceTypeID uuid.UUID) (*entities.ServiceType, error) {
	var serviceType entities.ServiceType
	err := r.db.WithContext(ctx).Where("service_type_id = ?", serviceTypeID).First(&serviceType).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return &serviceType, nil
}

func (r *serviceTypeRepositoryImpl) GetAllServiceTypes(ctx context.Context) ([]entities.ServiceType, error) {
	var serviceTypes []entities.ServiceType
	err := r.db.WithContext(ctx).Find(&serviceTypes).Error
	if err != nil {
		return nil, err
	}
	return serviceTypes, nil
}

func (r *serviceTypeRepositoryImpl) CreateServiceType(ctx context.Context, serviceType entities.ServiceType) error {
	err := r.db.WithContext(ctx).Create(&serviceType).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *serviceTypeRepositoryImpl) UpdateServiceType(ctx context.Context, serviceType entities.ServiceType) error {
	err := r.db.WithContext(ctx).Save(&serviceType).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *serviceTypeRepositoryImpl) DeleteServiceType(ctx context.Context, serviceTypeID uuid.UUID) error {
	err := r.db.WithContext(ctx).Where("service_type_id = ?", serviceTypeID).Delete(&entities.ServiceType{}).Error
	if err != nil {
		return err
	}
	return nil
}