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

type PoolPointHandler struct {
	poolPointService services.PoolPointService
	validator        *validator.Validate
	log              *logrus.Logger
}

func NewPoolPointHandler(poolPointService services.PoolPointService, validator *validator.Validate, log *logrus.Logger) *PoolPointHandler {
	return &PoolPointHandler{
		poolPointService: poolPointService,
		validator:        validator,
		log:              log,
	}
}

func (h *PoolPointHandler) GetPoolPointByID(c fiber.Ctx) error {
	id := c.Params("id")
	poolID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid pool point ID", err)
	}

	poolPoint, err := h.poolPointService.GetPoolPointByID(c.Context(), poolID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Pool point retrieved successfully", poolPoint)
}

func (h *PoolPointHandler) GetAllPoolPoints(c fiber.Ctx) error {
	poolPoints, err := h.poolPointService.GetAllPoolPoints(c.Context())
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Pool points retrieved successfully", poolPoints)
}

func (h *PoolPointHandler) GetPoolPointsByVendorID(c fiber.Ctx) error {
	vendorIDStr := c.Params("vendorId")
	vendorID, err := uuid.Parse(vendorIDStr)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid vendor ID", err)
	}

	poolPoints, err := h.poolPointService.GetPoolPointsByVendorID(c.Context(), vendorID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Pool points retrieved successfully", poolPoints)
}

func (h *PoolPointHandler) CreatePoolPoint(c fiber.Ctx) error {
	var req dto.CreatePoolPointDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	err := h.poolPointService.CreatePoolPoint(c.Context(), req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusCreated, "Pool point created successfully", nil)
}

func (h *PoolPointHandler) UpdatePoolPoint(c fiber.Ctx) error {
	id := c.Params("id")
	poolID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid pool point ID", err)
	}

	var req dto.UpdatePoolPointDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	err = h.poolPointService.UpdatePoolPoint(c.Context(), poolID, req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Pool point updated successfully", nil)
}

func (h *PoolPointHandler) DeletePoolPoint(c fiber.Ctx) error {
	id := c.Params("id")
	poolID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid pool point ID", err)
	}

	err = h.poolPointService.DeletePoolPoint(c.Context(), poolID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Pool point deleted successfully", nil)
}