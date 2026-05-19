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

type LayoutHandler struct {
	layoutService services.LayoutService
	validator     *validator.Validate
	log           *logrus.Logger
}

func NewLayoutHandler(layoutService services.LayoutService, validator *validator.Validate, log *logrus.Logger) *LayoutHandler {
	return &LayoutHandler{
		layoutService: layoutService,
		validator:     validator,
		log:           log,
	}
}

func (h *LayoutHandler) GetLayoutByID(c fiber.Ctx) error {
	id := c.Params("id")
	layoutID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid layout ID", err)
	}

	layoutDTO, err := h.layoutService.GetLayoutById(c.Context(), layoutID.String())
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Layout retrieved successfully", layoutDTO)
}

func (h *LayoutHandler) GetLayout(c fiber.Ctx) error {
	page := c.Query("page", "1")
	pageSize := c.Query("page_size", "30")
	search := c.Query("search")
	sortBy := c.Query("sort_by", "created_at")
	sortOrder := c.Query("sort_order", "desc")

	layouts, total, err := h.layoutService.GetLayout(c.Context(), page, pageSize, search, sortBy, sortOrder)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	pagination := dto.Pagination{
		CurrentPage: page,
		PageSize:    pageSize,
		TotalItems:  int(total),
		TotalPages:  (int(total) + pageSize - 1) / pageSize,
		HasNext:     page*pageSize < int(total),
		HasPrevious: page > 1,
	}

	return httpResponses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Layouts retrieved successfully", layouts, pagination)
}

func (h *LayoutHandler) CreateLayout(c fiber.Ctx) error {
	var req dto.CreateLayoutDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	err := h.layoutService.CreateLayout(c.Context(), req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusCreated, "Layout created successfully", nil)
}

func (h *LayoutHandler) UpdateLayout(c fiber.Ctx) error {
	id := c.Params("id")
	layoutID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid layout ID", err)
	}

	var req dto.CreateLayoutDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	err = h.layoutService.UpdateLayout(c.Context(), layoutID.String(), req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Layout updated successfully", nil)
}

func (h *LayoutHandler) DeleteLayout(c fiber.Ctx) error {
	id := c.Params("id")
	layoutID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid layout ID", err)
	}

	err = h.layoutService.DeleteLayout(c.Context(), layoutID.String())
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Layout deleted successfully", nil)
}
