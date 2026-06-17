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

func (r *scheduleRepositoryImpl) GetAllSchedules(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.Schedules, int64, error) {
	var schedules []entities.Schedules
	var total int64

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize >= 1000 {
		pageSize = 30
	}

	offset := (page - 1) * pageSize

	baseQuery := r.db.WithContext(ctx).
		Model(&entities.Schedules{}).
		Preload("Layout").
		Preload("ServiceType").
		Preload("OriginPool").
		Preload("DestinationPool")

	if search != "" {
		baseQuery = baseQuery.Where("vehicle_type ILIKE ? OR status ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	allowedSort := map[string]bool{
		"vehicle_type":   true,
		"departure_date": true,
		"departure_time": true,
		"price_per_seat": true,
		"total_seat":     true,
		"available_seat": true,
		"status":         true,
		"created_at":     true,
		"updated_at":     true,
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
		Find(&schedules).Error
	if err != nil {
		return nil, 0, err
	}
	return schedules, total, nil
}

func (r *scheduleRepositoryImpl) GetSchedulesByVendorID(ctx context.Context, vendorID uuid.UUID, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.Schedules, int64, error) {
	var schedules []entities.Schedules
	var total int64

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize >= 1000 {
		pageSize = 30
	}

	offset := (page - 1) * pageSize

	baseQuery := r.db.WithContext(ctx).
		Model(&entities.Schedules{}).
		Preload("Layout").
		Preload("ServiceType").
		Preload("OriginPool").
		Preload("DestinationPool").
		Where("vendor_id = ?", vendorID)

	if search != "" {
		baseQuery = baseQuery.Where("vehicle_type ILIKE ? OR status ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	allowedSort := map[string]bool{
		"vehicle_type":   true,
		"departure_date": true,
		"departure_time": true,
		"price_per_seat": true,
		"total_seat":     true,
		"available_seat": true,
		"status":         true,
		"created_at":     true,
		"updated_at":     true,
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
		Find(&schedules).Error
	if err != nil {
		return nil, 0, err
	}
	return schedules, total, nil
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
