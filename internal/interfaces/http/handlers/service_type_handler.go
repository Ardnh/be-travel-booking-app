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

type ServiceTypeHandler struct {
	serviceTypeService services.ServiceTypeService
	validator          *validator.Validate
	log                *logrus.Logger
}

func NewServiceTypeHandler(serviceTypeService services.ServiceTypeService, validator *validator.Validate, log *logrus.Logger) *ServiceTypeHandler {
	return &ServiceTypeHandler{
		serviceTypeService: serviceTypeService,
		validator:          validator,
		log:                log,
	}
}

func (h *ServiceTypeHandler) GetServiceTypeByID(c fiber.Ctx) error {
	id := c.Params("id")
	serviceTypeID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid service type ID", err)
	}

	serviceType, err := h.serviceTypeService.GetServiceTypeByID(c.Context(), serviceTypeID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Service type retrieved successfully", serviceType)
}

func (h *ServiceTypeHandler) GetAllServiceTypes(c fiber.Ctx) error {
	page := c.Query("page", "1")
	pageSize := c.Query("page_size", "30")
	search := c.Query("search")
	sortBy := c.Query("sort_by", "created_at")
	sortOrder := c.Query("sort_order", "desc")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	serviceTypes, total, err := h.serviceTypeService.GetAllServiceTypes(c.Context(), pageInt, pageSizeInt, search, sortBy, sortOrder)
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

	return httpResponses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Service types retrieved successfully", serviceTypes, pagination)
}

func (h *ServiceTypeHandler) CreateServiceType(c fiber.Ctx) error {
	var req dto.CreateServiceTypeDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	// Get user ID from context (assuming set by auth middleware)
	userIDStr, ok := c.Locals("user_id").(string)
	h.log.Errorf("Parse owner user ID from c.local: %v", userIDStr)
	if !ok || userIDStr == "" {
		h.log.Errorf("User ID not found in context")
		return httpResponses.NewErrorResponse(c, fiber.ErrUnauthorized.Code, fiber.ErrUnauthorized.Message, "User ID not found")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		h.log.Errorf("Invalid owner user ID: %v", userID)
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, "Invalid owner user ID")
	}

	serviceType, errCreate := h.serviceTypeService.CreateServiceType(c.Context(), req, userID)
	if errCreate != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, errCreate)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusCreated, "Service type created successfully", serviceType)
}

func (h *ServiceTypeHandler) UpdateServiceType(c fiber.Ctx) error {
	id := c.Params("id")
	serviceTypeID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid service type ID", err)
	}

	var req dto.UpdateServiceTypeDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	serviceType, err := h.serviceTypeService.UpdateServiceType(c.Context(), serviceTypeID, req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Service type updated successfully", serviceType)
}

func (h *ServiceTypeHandler) DeleteServiceType(c fiber.Ctx) error {
	id := c.Params("id")
	serviceTypeID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid service type ID", err)
	}

	err = h.serviceTypeService.DeleteServiceType(c.Context(), serviceTypeID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Service type deleted successfully", nil)
}
