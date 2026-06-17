package handlers

import (
	"strconv"

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
	page := c.Query("page", "1")
	pageSize := c.Query("page_size", "30")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if pageInt <= 0 {
		pageInt = 1
	}
	if pageSizeInt <= 0 || pageSizeInt >= 1000 {
		pageSizeInt = 30
	}

	schedules, total, err := h.scheduleService.GetAllSchedules(c.Context(), pageInt, pageSizeInt, "", "created_at", "desc")
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	pagination := dto.Pagination{
		CurrentPage: pageInt,
		PageSize:    pageSizeInt,
		TotalItems:  int(total),
		TotalPages:  (int(total) + pageSizeInt - 1) / pageSizeInt,
		HasNext:     pageInt*pageSizeInt < int(total),
		HasPrevious: pageInt > 1,
	}

	return httpResponses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Schedules retrieved successfully", schedules, pagination)
}

func (h *ScheduleHandler) GetSchedulesByVendorID(c fiber.Ctx) error {
	vendorIDStr := c.Params("vendorId")
	vendorID, err := uuid.Parse(vendorIDStr)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid vendor ID", err)
	}

	page := c.Query("page", "1")
	pageSize := c.Query("page_size", "30")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if pageInt <= 0 {
		pageInt = 1
	}
	if pageSizeInt <= 0 || pageSizeInt >= 1000 {
		pageSizeInt = 30
	}

	schedules, total, err := h.scheduleService.GetSchedulesByVendorID(c.Context(), vendorID, pageInt, pageSizeInt, "", "created_at", "desc")
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	pagination := dto.Pagination{
		CurrentPage: pageInt,
		PageSize:    pageSizeInt,
		TotalItems:  int(total),
		TotalPages:  (int(total) + pageSizeInt - 1) / pageSizeInt,
		HasNext:     pageInt*pageSizeInt < int(total),
		HasPrevious: pageInt > 1,
	}

	return httpResponses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Schedules retrieved successfully", schedules, pagination)
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
