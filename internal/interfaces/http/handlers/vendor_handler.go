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

type VendorHandler struct {
	vendorService services.VendorService
	validator     *validator.Validate
	log           *logrus.Logger
}

func NewVendorHandler(vendorService services.VendorService, validator *validator.Validate, log *logrus.Logger) *VendorHandler {
	return &VendorHandler{
		vendorService: vendorService,
		validator:     validator,
		log:           log,
	}
}

func (h *VendorHandler) GetVendorByID(c fiber.Ctx) error {
	id := c.Params("id")
	vendorID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid vendor ID", err)
	}

	vendor, err := h.vendorService.GetVendorByID(c.Context(), vendorID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Vendor retrieved successfully", vendor)
}

func (h *VendorHandler) GetAllVendors(c fiber.Ctx) error {
	vendors, err := h.vendorService.GetAllVendors(c.Context())
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Vendors retrieved successfully", vendors)
}

func (h *VendorHandler) CreateVendor(c fiber.Ctx) error {
	var req dto.CreateVendorDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	err := h.vendorService.CreateVendor(c.Context(), req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusCreated, "Vendor created successfully", nil)
}

func (h *VendorHandler) UpdateVendor(c fiber.Ctx) error {
	id := c.Params("id")
	vendorID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid vendor ID", err)
	}

	var req dto.UpdateVendorDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	err = h.vendorService.UpdateVendor(c.Context(), vendorID, req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Vendor updated successfully", nil)
}

func (h *VendorHandler) DeleteVendor(c fiber.Ctx) error {
	id := c.Params("id")
	vendorID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid vendor ID", err)
	}

	err = h.vendorService.DeleteVendor(c.Context(), vendorID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Vendor deleted successfully", nil)
}
