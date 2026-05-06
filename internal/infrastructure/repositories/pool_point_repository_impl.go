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

type poolPointRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPoolPointRepository(db *gorm.DB, redis *redis.Client) repositories.PoolPointRepository {
	return &poolPointRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *poolPointRepositoryImpl) GetPoolPointByID(ctx context.Context, poolID uuid.UUID) (*entities.PoolPoint, error) {
	var poolPoint entities.PoolPoint
	err := r.db.WithContext(ctx).Where("pool_id = ?", poolID).First(&poolPoint).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return &poolPoint, nil
}

func (r *poolPointRepositoryImpl) GetAllPoolPoints(ctx context.Context) ([]entities.PoolPoint, error) {
	var poolPoints []entities.PoolPoint
	err := r.db.WithContext(ctx).Find(&poolPoints).Error
	if err != nil {
		return nil, err
	}
	return poolPoints, nil
}

func (r *poolPointRepositoryImpl) GetPoolPointsByVendorID(ctx context.Context, vendorID uuid.UUID) ([]entities.PoolPoint, error) {
	var poolPoints []entities.PoolPoint
	err := r.db.WithContext(ctx).Where("vendor_id = ?", vendorID).Find(&poolPoints).Error
	if err != nil {
		return nil, err
	}
	return poolPoints, nil
}

func (r *poolPointRepositoryImpl) CreatePoolPoint(ctx context.Context, poolPoint entities.PoolPoint) error {
	err := r.db.WithContext(ctx).Create(&poolPoint).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *poolPointRepositoryImpl) UpdatePoolPoint(ctx context.Context, poolPoint entities.PoolPoint) error {
	err := r.db.WithContext(ctx).Save(&poolPoint).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *poolPointRepositoryImpl) DeletePoolPoint(ctx context.Context, poolID uuid.UUID) error {
	err := r.db.WithContext(ctx).Where("pool_id = ?", poolID).Delete(&entities.PoolPoint{}).Error
	if err != nil {
		return err
	}
	return nil
}