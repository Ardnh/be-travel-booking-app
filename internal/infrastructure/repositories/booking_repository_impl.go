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

type bookingRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewBookingRepository(db *gorm.DB, redis *redis.Client) repositories.BookingRepository {
	return &bookingRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *bookingRepositoryImpl) GetBookingByID(ctx context.Context, bookingID uuid.UUID) (*entities.Bookings, error) {
	var booking entities.Bookings
	err := r.db.WithContext(ctx).Where("booking_id = ?", bookingID).First(&booking).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepositoryImpl) GetBookingsByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Bookings, error) {
	var bookings []entities.Bookings
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *bookingRepositoryImpl) GetBookingsByScheduleID(ctx context.Context, scheduleID uuid.UUID) ([]entities.Bookings, error) {
	var bookings []entities.Bookings
	err := r.db.WithContext(ctx).Where("schedule_id = ?", scheduleID).Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *bookingRepositoryImpl) CreateBooking(ctx context.Context, booking entities.Bookings) error {
	err := r.db.WithContext(ctx).Create(&booking).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *bookingRepositoryImpl) UpdateBooking(ctx context.Context, booking entities.Bookings) error {
	err := r.db.WithContext(ctx).Save(&booking).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *bookingRepositoryImpl) DeleteBooking(ctx context.Context, bookingID uuid.UUID) error {
	err := r.db.WithContext(ctx).Where("booking_id = ?", bookingID).Delete(&entities.Bookings{}).Error
	if err != nil {
		return err
	}
	return nil
}