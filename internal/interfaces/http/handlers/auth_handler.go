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

type AuthHandler struct {
	authService services.AuthService
	validator   *validator.Validate
	log         *logrus.Logger
}

func NewAuthHandler(authService services.AuthService, validator *validator.Validate, log *logrus.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator,
		log:         log,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req dto.LoginRequestDto
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.HandleError(c, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	result, err := h.authService.Login(c.Context(), req)
	if err != nil {
		return httpResponses.HandleError(c, err)
	}

	// set cookie setelah login berhasil

	return httpResponses.NewSuccessResponse(c, fiber.StatusOK, "Login successful", result)
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req dto.RegisterRequestDto
	if err := c.Bind().Body(&req); err != nil {
		return httpResponses.HandleError(c, err)
	}

	if err := h.validator.Struct(&req); err != nil {
		return httpResponses.NewErrorResponse(c, fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message, validator_utils.FormatValidationErrors(err))
	}

	_, err := h.authService.Register(c.Context(), req)
	if err != nil {
		return httpResponses.HandleError(c, err)
	}

	return httpResponses.NewSuccessResponse(c, fiber.StatusCreated, "Register successful", nil)
}
