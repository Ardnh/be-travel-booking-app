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
	serviceTypes, err := h.serviceTypeService.GetAllServiceTypes(c.Context())
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Service types retrieved successfully", serviceTypes)
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
	userID := c.Locals("userID").(uuid.UUID)

	err := h.serviceTypeService.CreateServiceType(c.Context(), req, userID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusCreated, "Service type created successfully", nil)
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

	err = h.serviceTypeService.UpdateServiceType(c.Context(), serviceTypeID, req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Service type updated successfully", nil)
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