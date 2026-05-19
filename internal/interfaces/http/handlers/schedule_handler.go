package handlers

import (
	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/services"
	httpResponses "github.com/ardnh/be-travel-booking-app/internal/interfaces/http/responses"
	validator_utils "github.com/ardnh/be-travel-booking-app/internal/utils/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ScheduleHandler struct {
	scheduleService services.ScheduleService
	validator       *validator.Validate
	log             *logrus.Logger
}

func NewScheduleHandler(scheduleService services.ScheduleService, validator *validator.Validate, log *logrus.Logger) *ScheduleHandler {
	return &ScheduleHandler{
		scheduleService: scheduleService,
		validator:       validator,
		log:             log,
	}
}

func (h *ScheduleHandler) GetScheduleByID(c fiber.Ctx) error {
	id := c.Params("id")
	scheduleID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid schedule ID", err)
	}

	schedule, err := h.scheduleService.GetScheduleByID(c.Context(), scheduleID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Schedule retrieved successfully", schedule)
}

func (h *ScheduleHandler) GetAllSchedules(c fiber.Ctx) error {
	schedules, err := h.scheduleService.GetAllSchedules(c.Context())
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Schedules retrieved successfully", schedules)
}

func (h *ScheduleHandler) GetSchedulesByVendorID(c fiber.Ctx) error {
	vendorIDStr := c.Params("vendorId")
	vendorID, err := uuid.Parse(vendorIDStr)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid vendor ID", err)
	}

	schedules, err := h.scheduleService.GetSchedulesByVendorID(c.Context(), vendorID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Schedules retrieved successfully", schedules)
}

func (h *ScheduleHandler) CreateSchedule(c fiber.Ctx) error {
	var req dto.CreateScheduleDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	err := h.scheduleService.CreateSchedule(c.Context(), req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusCreated, "Schedule created successfully", nil)
}

func (h *ScheduleHandler) UpdateSchedule(c fiber.Ctx) error {
	id := c.Params("id")
	scheduleID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid schedule ID", err)
	}

	var req dto.UpdateScheduleDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	err = h.scheduleService.UpdateSchedule(c.Context(), scheduleID, req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Schedule updated successfully", nil)
}

func (h *ScheduleHandler) DeleteSchedule(c fiber.Ctx) error {
	id := c.Params("id")
	scheduleID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid schedule ID", err)
	}

	err = h.scheduleService.DeleteSchedule(c.Context(), scheduleID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Schedule deleted successfully", nil)
}
