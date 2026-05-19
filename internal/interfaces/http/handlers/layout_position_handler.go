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

type LayoutPositionHandler struct {
	layoutPositionService services.LayoutPositionService
	validator             *validator.Validate
	log                   *logrus.Logger
}

func NewLayoutPositionHandler(layoutPositionService services.LayoutPositionService, validator *validator.Validate, log *logrus.Logger) *LayoutPositionHandler {
	return &LayoutPositionHandler{
		layoutPositionService: layoutPositionService,
		validator:             validator,
		log:                   log,
	}
}

func (h *LayoutPositionHandler) GetLayoutPositionByID(c fiber.Ctx) error {
	id := c.Params("id")
	layoutPosID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid layout position ID", err)
	}

	position, err := h.layoutPositionService.GetLayoutPositionByID(c.Context(), layoutPosID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Layout position retrieved successfully", position)
}

func (h *LayoutPositionHandler) GetLayoutPositionsByLayoutID(c fiber.Ctx) error {
	layoutIDStr := c.Params("layoutId")
	layoutID, err := uuid.Parse(layoutIDStr)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid layout ID", err)
	}

	positions, err := h.layoutPositionService.GetLayoutPositionsByLayoutID(c.Context(), layoutID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Layout positions retrieved successfully", positions)
}

func (h *LayoutPositionHandler) GetAllLayoutPositions(c fiber.Ctx) error {
	positions, err := h.layoutPositionService.GetAllLayoutPositions(c.Context())
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Layout positions retrieved successfully", positions)
}

func (h *LayoutPositionHandler) CreateLayoutPosition(c fiber.Ctx) error {
	layoutIDStr := c.Params("layoutId")
	layoutID, err := uuid.Parse(layoutIDStr)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid layout ID", err)
	}

	var req dto.CreateLayoutPositionDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	err = h.layoutPositionService.CreateLayoutPosition(c.Context(), req, layoutID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusCreated, "Layout position created successfully", nil)
}

func (h *LayoutPositionHandler) UpdateLayoutPosition(c fiber.Ctx) error {
	id := c.Params("id")
	layoutPosID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid layout position ID", err)
	}

	var req dto.UpdateLayoutPositionDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	err = h.layoutPositionService.UpdateLayoutPosition(c.Context(), layoutPosID, req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Layout position updated successfully", nil)
}

func (h *LayoutPositionHandler) DeleteLayoutPosition(c fiber.Ctx) error {
	id := c.Params("id")
	layoutPosID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid layout position ID", err)
	}

	err = h.layoutPositionService.DeleteLayoutPosition(c.Context(), layoutPosID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Layout position deleted successfully", nil)
}
