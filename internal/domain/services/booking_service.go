package services

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
)

type BookingService interface {
	GetBookingByID(ctx context.Context, bookingID string) (*dto.BookingDTO, error)
	GetBookingsByUserID(ctx context.Context, userID string) ([]dto.BookingDTO, error)
	GetBookingsByScheduleID(ctx context.Context, scheduleID string) ([]dto.BookingDTO, error)
	CreateBooking(ctx context.Context, req dto.CreateBookingDTO) (*dto.BookingDTO, error)
	UpdateBooking(ctx context.Context, bookingID string, req dto.UpdateBookingDTO) (*dto.BookingDTO, error)
	DeleteBooking(ctx context.Context, bookingID string) error
}