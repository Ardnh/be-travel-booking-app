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
	h.log.Infof("Starting GetVendorByID")
	id := c.Params("id")
	vendorID, err := uuid.Parse(id)
	if err != nil {
		h.log.Errorf("Failed to parse vendor ID: %v", err)
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid vendor ID", err)
	}

	vendor, err := h.vendorService.GetVendorByID(c.Context(), vendorID)
	if err != nil {
		h.log.Errorf("Failed to get vendor by ID: %v", err)
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	h.log.Infof("Successfully retrieved vendor by ID: %s", vendorID)
	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Vendor retrieved successfully", vendor)
}

func (h *VendorHandler) GetVendorByOwnerUserID(c fiber.Ctx) error {
	h.log.Infof("Starting GetVendorByOwnerUserID")
	ownerUserID, ok := c.Locals("user_id").(string)
	h.log.Errorf("Parse owner user ID from c.local: %v", ownerUserID)
	if !ok || ownerUserID == "" {
		h.log.Errorf("User ID not found in context")
		return httpResponses.NewErrorResponse(c, fiber.ErrUnauthorized.Code, fiber.ErrUnauthorized.Message, "User ID not found")
	}

	ownerUserIDUuid, err := uuid.Parse(ownerUserID)
	h.log.Errorf("Parsed owner user ID: %v", ownerUserIDUuid)
	if err != nil {
		h.log.Errorf("Failed to parse owner user ID: %v", err)
		return httpResponses.NewErrorResponse(c, fiber.ErrUnauthorized.Code, fiber.ErrUnauthorized.Message, "Invalid owner user id")
	}

	vendor, err := h.vendorService.GetVendorByOwnerUserID(c.Context(), ownerUserIDUuid)
	if err != nil {
		h.log.Errorf("Failed to get vendor by owner user ID: %v", err)
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	h.log.Infof("Successfully retrieved vendor by owner user ID: %s", ownerUserID)
	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Vendor retrieved successfully", vendor)
}

func (h *VendorHandler) GetAllVendors(c fiber.Ctx) error {
	h.log.Infof("Starting GetAllVendors")

	page := c.Query("page", "1")
	pageSize := c.Query("page_size", "10")
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

	h.log.Infof("Fetching vendors with page=%d, page_size=%d, search=%s, sort_by=%s, sort_order=%s", pageInt, pageSizeInt, search, sortBy, sortOrder)

	vendors, total, err := h.vendorService.GetAllVendors(c.Context(), pageInt, pageSizeInt, search, sortBy, sortOrder)
	if err != nil {
		h.log.Errorf("Failed to get all vendors: %v", err)
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

	h.log.Infof("Successfully retrieved vendors. Total: %d, Page: %d, Total Pages: %d", total, pageInt, pagination.TotalPages)
	return httpResponses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Vendors retrieved successfully", vendors, pagination)
}

func (h *VendorHandler) CreateVendor(c fiber.Ctx) error {
	h.log.Infof("Starting CreateVendor")
	var req dto.CreateVendorDTO
	if err := c.Bind().Body(&req); err != nil {
		h.log.Errorf("Failed to bind request body: %v", err)
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		h.log.Errorf("Validation failed for CreateVendor: %v", err)
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	err := h.vendorService.CreateVendor(c.Context(), req)
	if err != nil {
		h.log.Errorf("Failed to create vendor: %v", err)
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	h.log.Infof("Successfully created vendor")
	return httpResponses.NewSuccessResponse(c, fiber.StatusCreated, "Vendor created successfully", nil)
}

func (h *VendorHandler) UpdateVendor(c fiber.Ctx) error {
	h.log.Infof("Starting UpdateVendor")
	id := c.Params("id")
	vendorID, err := uuid.Parse(id)
	if err != nil {
		h.log.Errorf("Failed to parse vendor ID: %v", err)
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid vendor ID", err)
	}

	var req dto.UpdateVendorDTO
	if err := c.Bind().Body(&req); err != nil {
		h.log.Errorf("Failed to bind request body: %v", err)
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	err = h.vendorService.UpdateVendor(c.Context(), vendorID, req)
	if err != nil {
		h.log.Errorf("Failed to update vendor: %v", err)
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	h.log.Infof("Successfully updated vendor ID: %s", vendorID)
	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Vendor updated successfully", nil)
}

func (h *VendorHandler) DeleteVendor(c fiber.Ctx) error {
	h.log.Infof("Starting DeleteVendor")
	id := c.Params("id")
	vendorID, err := uuid.Parse(id)
	if err != nil {
		h.log.Errorf("Failed to parse vendor ID: %v", err)
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid vendor ID", err)
	}

	err = h.vendorService.DeleteVendor(c.Context(), vendorID)
	if err != nil {
		h.log.Errorf("Failed to delete vendor: %v", err)
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	h.log.Infof("Successfully deleted vendor ID: %s", vendorID)
	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Vendor deleted successfully", nil)
}
