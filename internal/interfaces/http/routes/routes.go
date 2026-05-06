package routes

import (
	// middleware "github.com/ardnh/be-travel-booking-app/internal/interfaces/http/middleware"
	"github.com/ardnh/be-travel-booking-app/internal/interfaces/http/handlers"
	"github.com/casbin/casbin/v3"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

func SetupAPIRoutes(app *fiber.App, log *logrus.Logger, enforcer *casbin.Enforcer, serviceTypeHandler *handlers.ServiceTypeHandler, poolPointHandler *handlers.PoolPointHandler, validator *validator.Validate) {

	// Middleware
	// casbinMiddleware := middleware.NewCasbinMiddleware(enforcer)
	// authMiddleware := middleware.NewAuthMiddleware()

	// API v1 group
	api := app.Group("/api/v1")

	// Service Type routes
	api.Get("/service-types", serviceTypeHandler.GetAllServiceTypes)
	api.Get("/service-types/:id", serviceTypeHandler.GetServiceTypeByID)
	api.Post("/service-types", serviceTypeHandler.CreateServiceType)
	api.Put("/service-types/:id", serviceTypeHandler.UpdateServiceType)
	api.Delete("/service-types/:id", serviceTypeHandler.DeleteServiceType)

	// Pool Point routes
	api.Get("/pool-points", poolPointHandler.GetAllPoolPoints)
	api.Get("/pool-points/:id", poolPointHandler.GetPoolPointByID)
	api.Get("/vendors/:vendorId/pool-points", poolPointHandler.GetPoolPointsByVendorID)
	api.Post("/pool-points", poolPointHandler.CreatePoolPoint)
	api.Put("/pool-points/:id", poolPointHandler.UpdatePoolPoint)
	api.Delete("/pool-points/:id", poolPointHandler.DeletePoolPoint)

}
