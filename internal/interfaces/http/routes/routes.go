package routes

import (
	// middleware "github.com/ardnh/be-travel-booking-app/internal/interfaces/http/middleware"
	"github.com/casbin/casbin/v3"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

func SetupAPIRoutes(app *fiber.App, log *logrus.Logger, enforcer *casbin.Enforcer) {

	// Middleware
	// casbinMiddleware := middleware.NewCasbinMiddleware(enforcer)
	// authMiddleware := middleware.NewAuthMiddleware()

	// API v1 group
	// api := app.Group("/api/v1")

}
