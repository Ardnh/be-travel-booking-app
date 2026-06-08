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

func (r *serviceTypeRepositoryImpl) GetServiceTypeByID(ctx context.Context, serviceTypeID uuid.UUID) (*entities.ServiceTypes, error) {
	var serviceType entities.ServiceTypes
	err := r.db.WithContext(ctx).Where("service_type_id = ?", serviceTypeID).First(&serviceType).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return &serviceType, nil
}

func (r *serviceTypeRepositoryImpl) GetAllServiceTypes(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.ServiceTypes, int64, error) {
	var (
		serviceTypes []entities.ServiceTypes
		total        int64
	)

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize >= 1000 {
		pageSize = 30
	}

	offset := (page - 1) * pageSize

	baseQuery := r.db.WithContext(ctx).Model(&entities.ServiceTypes{})

	if search != "" {
		baseQuery = baseQuery.Where("name ILIKE ?", "%"+search+"%")
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	allowedSort := map[string]bool{
		"name":          true,
		"created_at":    true,
		"display_order": true,
	}

	if !allowedSort[sortBy] {
		sortBy = "created_at"
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	err := baseQuery.
		Order(sortBy + " " + sortOrder).
		Limit(pageSize).
		Offset(offset).
		Find(&serviceTypes).Error

	if err != nil {
		return nil, 0, err
	}

	return serviceTypes, total, nil
}

func (r *serviceTypeRepositoryImpl) CreateServiceType(ctx context.Context, serviceType entities.ServiceTypes) (*entities.ServiceTypes, error) {
	err := r.db.WithContext(ctx).Create(&serviceType).Error
	if err != nil {
		return nil, err
	}
	return &serviceType, nil
}

func (r *serviceTypeRepositoryImpl) UpdateServiceType(ctx context.Context, serviceType entities.ServiceTypes) (*entities.ServiceTypes, error) {
	err := r.db.WithContext(ctx).Save(&serviceType).Error
	if err != nil {
		return nil, err
	}
	return &serviceType, nil
}

func (r *serviceTypeRepositoryImpl) DeleteServiceType(ctx context.Context, serviceTypeID uuid.UUID) error {
	err := r.db.WithContext(ctx).Where("service_type_id = ?", serviceTypeID).Delete(&entities.ServiceTypes{}).Error
	if err != nil {
		return err
	}
	return nil
}
