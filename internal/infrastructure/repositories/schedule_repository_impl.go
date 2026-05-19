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

type scheduleRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewScheduleRepository(db *gorm.DB, redis *redis.Client) repositories.ScheduleRepository {
	return &scheduleRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *scheduleRepositoryImpl) GetScheduleByID(ctx context.Context, scheduleID uuid.UUID) (*entities.Schedules, error) {
	var schedule entities.Schedules
	err := r.db.WithContext(ctx).
		Preload("Layout").
		Preload("ServiceType").
		Preload("OriginPool").
		Preload("DestinationPool").
		Where("schedule_id = ?", scheduleID).
		First(&schedule).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return &schedule, nil
}

func (r *scheduleRepositoryImpl) GetAllSchedules(ctx context.Context) ([]entities.Schedules, error) {
	var schedules []entities.Schedules
	err := r.db.WithContext(ctx).
		Preload("Layout").
		Preload("ServiceType").
		Preload("OriginPool").
		Preload("DestinationPool").
		Find(&schedules).Error
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *scheduleRepositoryImpl) GetSchedulesByVendorID(ctx context.Context, vendorID uuid.UUID) ([]entities.Schedules, error) {
	var schedules []entities.Schedules
	err := r.db.WithContext(ctx).
		Preload("Layout").
		Preload("ServiceType").
		Preload("OriginPool").
		Preload("DestinationPool").
		Where("vendor_id = ?", vendorID).
		Find(&schedules).Error
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *scheduleRepositoryImpl) CreateSchedule(ctx context.Context, schedule entities.Schedules) error {
	err := r.db.WithContext(ctx).Create(&schedule).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *scheduleRepositoryImpl) UpdateSchedule(ctx context.Context, schedule entities.Schedules) error {
	err := r.db.WithContext(ctx).Save(&schedule).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *scheduleRepositoryImpl) DeleteSchedule(ctx context.Context, scheduleID uuid.UUID) error {
	err := r.db.WithContext(ctx).Where("schedule_id = ?", scheduleID).Delete(&entities.Schedules{}).Error
	if err != nil {
		return err
	}
	return nil
}
