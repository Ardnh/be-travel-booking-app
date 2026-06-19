package handlers

import (
	"errors"
	"strconv"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/services"
	httpResponses "github.com/ardnh/be-travel-booking-app/internal/interfaces/http/responses"
	validator_utils "github.com/ardnh/be-travel-booking-app/internal/utils/validator"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
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

	poolPoints, total, err := h.poolPointService.GetAllPoolPoints(c.Context(), pageInt, pageSizeInt, search, sortBy, sortOrder)
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

	return httpResponses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Pool points retrieved successfully", poolPoints, pagination)
}

func (h *PoolPointHandler) GetPoolPointsByVendorID(c fiber.Ctx) error {
	vendorIDStr := c.Params("vendorId")
	vendorID, err := uuid.Parse(vendorIDStr)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid vendor ID", err)
	}

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

	poolPoints, total, err := h.poolPointService.GetPoolPointsByVendorID(c.Context(), vendorID, pageInt, pageSizeInt, search, sortBy, sortOrder)
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

	return httpResponses.NewSuccessResponseWithPagination(c, fiber.StatusOK, "Pool points retrieved successfully", poolPoints, pagination)
}

func (h *PoolPointHandler) GetAvailableLocationsByVendorID(c fiber.Ctx) error {
	vendorIDStr := c.Params("vendorId")
	vendorID, err := uuid.Parse(vendorIDStr)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid vendor ID", err)
	}

	locationType := c.Query("location_type", "")
	locations, err := h.poolPointService.GetAvailableLocationsByVendorID(c.Context(), vendorID, locationType)
	if err != nil {
		if errors.Is(err, errorConst.ErrBadRequest) {
			return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid location type", err)
		}
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Available locations retrieved successfully", locations)
}

func (h *PoolPointHandler) CreatePoolPoint(c fiber.Ctx) error {
	var req dto.CreatePoolsDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	poolPoint, err := h.poolPointService.CreatePoolPoint(c.Context(), req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusCreated, "Pool point created successfully", poolPoint)
}

func (h *PoolPointHandler) UpdatePoolPoint(c fiber.Ctx) error {
	id := c.Params("id")
	poolID, err := uuid.Parse(id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid pool point ID", err)
	}

	var req dto.UpdatePoolsDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	poolPoint, err := h.poolPointService.UpdatePoolPoint(c.Context(), poolID, req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Pool point updated successfully", poolPoint)
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
