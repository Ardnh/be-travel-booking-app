package handlers

import (
	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/services"
	httpResponses "github.com/ardnh/be-travel-booking-app/internal/interfaces/http/responses"
	validator_utils "github.com/ardnh/be-travel-booking-app/internal/utils/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

type UserRolesHandler struct {
	userRolesService services.UserRolesService
	validator        *validator.Validate
	log              *logrus.Logger
}

func NewUserRolesHandler(userRolesService services.UserRolesService, validator *validator.Validate, log *logrus.Logger) *UserRolesHandler {
	return &UserRolesHandler{
		userRolesService: userRolesService,
		validator:        validator,
		log:              log,
	}
}

func (h *UserRolesHandler) GetUserRoleByID(c fiber.Ctx) error {
	id := c.Params("id")
	userRole, err := h.userRolesService.GetUserRoleByID(c.Context(), id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}
	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "User role retrieved successfully", userRole)
}

func (h *UserRolesHandler) GetUserRolesByUserID(c fiber.Ctx) error {
	userID := c.Params("userId")
	userRoles, err := h.userRolesService.GetUserRolesByUserID(c.Context(), userID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}
	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "User roles retrieved successfully", userRoles)
}

func (h *UserRolesHandler) GetUserRolesByVendorID(c fiber.Ctx) error {
	vendorID := c.Params("vendorId")
	userRoles, err := h.userRolesService.GetUserRolesByVendorID(c.Context(), vendorID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}
	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "User roles retrieved successfully", userRoles)
}

func (h *UserRolesHandler) CreateUserRole(c fiber.Ctx) error {
	var req dto.CreateUserRoleDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	userRole, err := h.userRolesService.CreateUserRole(c.Context(), req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusCreated, "User role created successfully", userRole)
}

func (h *UserRolesHandler) UpdateUserRole(c fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateUserRoleDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	userRole, err := h.userRolesService.UpdateUserRole(c.Context(), id, req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "User role updated successfully", userRole)
}

func (h *UserRolesHandler) DeleteUserRole(c fiber.Ctx) error {
	id := c.Params("id")
	err := h.userRolesService.DeleteUserRole(c.Context(), id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}
	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "User role deleted successfully", nil)
}