package repositories

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type BookingRepository interface {
	GetBookingByID(ctx context.Context, bookingID uuid.UUID) (*entities.Bookings, error)
	GetBookingsByUserID(ctx context.Context, userID uuid.UUID) ([]entities.Bookings, error)
	GetBookingsByScheduleID(ctx context.Context, scheduleID uuid.UUID) ([]entities.Bookings, error)
	CreateBooking(ctx context.Context, booking entities.Bookings) error
	UpdateBooking(ctx context.Context, booking entities.Bookings) error
	DeleteBooking(ctx context.Context, bookingID uuid.UUID) error
}