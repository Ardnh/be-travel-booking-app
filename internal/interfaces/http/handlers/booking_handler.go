package handlers

import (
	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/services"
	httpResponses "github.com/ardnh/be-travel-booking-app/internal/interfaces/http/responses"
	validator_utils "github.com/ardnh/be-travel-booking-app/internal/utils/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type BookingHandler struct {
	bookingService services.BookingService
	validator      *validator.Validate
	log            *logrus.Logger
}

func NewBookingHandler(bookingService services.BookingService, validator *validator.Validate, log *logrus.Logger) *BookingHandler {
	return &BookingHandler{
		bookingService: bookingService,
		validator:      validator,
		log:            log,
	}
}

func (h *BookingHandler) GetBookingByID(c fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.bookingService.GetBookingByID(c.Context(), id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}
	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Booking retrieved successfully", booking)
}

func (h *BookingHandler) GetBookingsByUserID(c fiber.Ctx) error {
	userID := c.Params("userId")
	bookings, err := h.bookingService.GetBookingsByUserID(c.Context(), userID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}
	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Bookings retrieved successfully", bookings)
}

func (h *BookingHandler) GetBookingsByScheduleID(c fiber.Ctx) error {
	scheduleID := c.Params("scheduleId")
	bookings, err := h.bookingService.GetBookingsByScheduleID(c.Context(), scheduleID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}
	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Bookings retrieved successfully", bookings)
}

func (h *BookingHandler) CreateBooking(c fiber.Ctx) error {
	var req dto.CreateBookingDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	booking, err := h.bookingService.CreateBooking(c.Context(), req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusCreated, "Booking created successfully", booking)
}

func (h *BookingHandler) UpdateBooking(c fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateBookingDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	booking, err := h.bookingService.UpdateBooking(c.Context(), id, req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Booking updated successfully", booking)
}

func (h *BookingHandler) DeleteBooking(c fiber.Ctx) error {
	id := c.Params("id")
	err := h.bookingService.DeleteBooking(c.Context(), id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}
	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Booking deleted successfully", nil)
}