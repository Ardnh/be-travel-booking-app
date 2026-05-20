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

type UsersHandler struct {
	usersService services.UsersService
	validator    *validator.Validate
	log          *logrus.Logger
}

func NewUsersHandler(usersService services.UsersService, validator *validator.Validate, log *logrus.Logger) *UsersHandler {
	return &UsersHandler{
		usersService: usersService,
		validator:    validator,
		log:          log,
	}
}

func (h *UsersHandler) CreateUser(c fiber.Ctx) error {
	var req dto.CreateUserDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	err := h.usersService.CreateUser(c.Context(), req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusCreated, "User created successfully", nil)
}

func (h *UsersHandler) UpdateUser(c fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateUserDTO
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, err)
	}

	err := h.usersService.UpdateUser(c.Context(), id, req)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "User updated successfully", nil)
}

func (h *UsersHandler) DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	err := h.usersService.DeleteUser(c.Context(), id)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}
	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "User deleted successfully", nil)
}

func (h *UsersHandler) GetUserProfile(c fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(string)
	if !ok || userID == "" {
		return httpResponses.NewErrorResponse(c, fiber.ErrUnauthorized.Code, fiber.ErrUnauthorized.Message, "User ID not found")
	}

	profile, err := h.usersService.GetUserProfile(c.Context(), userID)
	if err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message, err)
	}
	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "User profile retrieved successfully", profile)
}
