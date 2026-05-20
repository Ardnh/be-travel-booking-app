package services

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type BookingServiceImpl struct {
	bookingRepository repositories.BookingRepository
	log               *logrus.Logger
}

func NewBookingServiceImpl(bookingRepository repositories.BookingRepository, log *logrus.Logger) *BookingServiceImpl {
	return &BookingServiceImpl{
		bookingRepository: bookingRepository,
		log:               log,
	}
}

func (s *BookingServiceImpl) GetBookingByID(ctx context.Context, bookingID string) (*dto.BookingDTO, error) {
	bookingUUID, err := uuid.Parse(bookingID)
	if err != nil {
		s.log.WithField("booking_id", bookingID).Error("invalid booking ID")
		return nil, errorConst.ErrBadRequest
	}

	booking, err := s.bookingRepository.GetBookingByID(ctx, bookingUUID)
	if err != nil {
		return nil, err
	}

	return s.toDTO(booking), nil
}

func (s *BookingServiceImpl) GetBookingsByUserID(ctx context.Context, userID string) ([]dto.BookingDTO, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		s.log.WithField("user_id", userID).Error("invalid user ID")
		return nil, errorConst.ErrBadRequest
	}

	bookings, err := s.bookingRepository.GetBookingsByUserID(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	return s.toDTOList(bookings), nil
}

func (s *BookingServiceImpl) GetBookingsByScheduleID(ctx context.Context, scheduleID string) ([]dto.BookingDTO, error) {
	scheduleUUID, err := uuid.Parse(scheduleID)
	if err != nil {
		s.log.WithField("schedule_id", scheduleID).Error("invalid schedule ID")
		return nil, errorConst.ErrBadRequest
	}

	bookings, err := s.bookingRepository.GetBookingsByScheduleID(ctx, scheduleUUID)
	if err != nil {
		return nil, err
	}

	return s.toDTOList(bookings), nil
}

func (s *BookingServiceImpl) CreateBooking(ctx context.Context, req dto.CreateBookingDTO) (*dto.BookingDTO, error) {
	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		s.log.WithField("user_id", req.UserID).Error("invalid user ID")
		return nil, errorConst.ErrBadRequest
	}

	scheduleUUID, err := uuid.Parse(req.ScheduleID)
	if err != nil {
		s.log.WithField("schedule_id", req.ScheduleID).Error("invalid schedule ID")
		return nil, errorConst.ErrBadRequest
	}

	booking := entities.Bookings{
		BookingID:     uuid.New(),
		BookingCode:   generateBookingCode(),
		UserID:        userUUID,
		ScheduleID:    scheduleUUID,
		PassengerName: req.PassengerName,
		PassengerPhone: req.PassengerPhone,
		PickupAddress: req.PickupAddress,
		DropoffAddress: req.DropoffAddress,
		PricePerSeat:  req.PricePerSeat,
		SeatCount:     req.SeatCount,
		TotalPrice:    req.PricePerSeat * float64(req.SeatCount),
	}

	err = s.bookingRepository.CreateBooking(ctx, booking)
	if err != nil {
		return nil, err
	}

	return s.toDTO(&booking), nil
}

func (s *BookingServiceImpl) UpdateBooking(ctx context.Context, bookingID string, req dto.UpdateBookingDTO) (*dto.BookingDTO, error) {
	bookingUUID, err := uuid.Parse(bookingID)
	if err != nil {
		s.log.WithField("booking_id", bookingID).Error("invalid booking ID")
		return nil, errorConst.ErrBadRequest
	}

	existingBooking, err := s.bookingRepository.GetBookingByID(ctx, bookingUUID)
	if err != nil {
		return nil, err
	}

	if req.PassengerName != nil {
		existingBooking.PassengerName = *req.PassengerName
	}
	if req.PassengerPhone != nil {
		existingBooking.PassengerPhone = *req.PassengerPhone
	}
	if req.PickupAddress != nil {
		existingBooking.PickupAddress = *req.PickupAddress
	}
	if req.DropoffAddress != nil {
		existingBooking.DropoffAddress = *req.DropoffAddress
	}
	if req.SeatCount != nil {
		existingBooking.SeatCount = *req.SeatCount
		existingBooking.TotalPrice = existingBooking.PricePerSeat * float64(*req.SeatCount)
	}
	if req.PricePerSeat != nil {
		existingBooking.PricePerSeat = *req.PricePerSeat
		existingBooking.TotalPrice = *req.PricePerSeat * float64(existingBooking.SeatCount)
	}
	if req.PaymentMethod != nil {
		existingBooking.PaymentMethod = *req.PaymentMethod
	}
	if req.Notes != nil {
		existingBooking.Notes = *req.Notes
	}

	err = s.bookingRepository.UpdateBooking(ctx, *existingBooking)
	if err != nil {
		return nil, err
	}

	return s.toDTO(existingBooking), nil
}

func (s *BookingServiceImpl) DeleteBooking(ctx context.Context, bookingID string) error {
	bookingUUID, err := uuid.Parse(bookingID)
	if err != nil {
		s.log.WithField("booking_id", bookingID).Error("invalid booking ID")
		return errorConst.ErrBadRequest
	}

	err = s.bookingRepository.DeleteBooking(ctx, bookingUUID)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookingServiceImpl) toDTO(booking *entities.Bookings) *dto.BookingDTO {
	return &dto.BookingDTO{
		BookingID:        booking.BookingID.String(),
		BookingCode:      booking.BookingCode,
		UserID:           booking.UserID.String(),
		ScheduleID:       booking.ScheduleID.String(),
		PassengerName:    booking.PassengerName,
		PassengerPhone:   booking.PassengerPhone,
		PickupAddress:    booking.PickupAddress,
		DropoffAddress:   booking.DropoffAddress,
		PricePerSeat:     booking.PricePerSeat,
		SeatCount:        booking.SeatCount,
		TotalPrice:       booking.TotalPrice,
		PaymentStatus:    booking.PaymentStatus,
		PaymentReference: booking.PaymentReference,
		PaymentMethod:    booking.PaymentMethod,
		BookingStatus:    booking.BookingStatus,
	}
}

func (s *BookingServiceImpl) toDTOList(bookings []entities.Bookings) []dto.BookingDTO {
	var result []dto.BookingDTO
	for _, b := range bookings {
		result = append(result, *s.toDTO(&b))
	}
	return result
}

func generateBookingCode() string {
	return "BOOK" + uuid.New().String()[:8]
}