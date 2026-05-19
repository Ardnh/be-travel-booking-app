package routes

import (
	// middleware "github.com/ardnh/be-travel-booking-app/internal/interfaces/http/middleware"
	"github.com/ardnh/be-travel-booking-app/internal/interfaces/http/handlers"
	"github.com/casbin/casbin/v3"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

func SetupAPIRoutes(app *fiber.App, log *logrus.Logger, enforcer *casbin.Enforcer, serviceTypeHandler *handlers.ServiceTypeHandler, poolPointHandler *handlers.PoolPointHandler, vendorHandler *handlers.VendorHandler, layoutHandler *handlers.LayoutHandler, layoutPositionHandler *handlers.LayoutPositionHandler, scheduleHandler *handlers.ScheduleHandler, validator *validator.Validate) {

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

	// Vendor routes
	api.Get("/vendors", vendorHandler.GetAllVendors)
	api.Get("/vendors/:id", vendorHandler.GetVendorByID)
	api.Post("/vendors", vendorHandler.CreateVendor)
	api.Put("/vendors/:id", vendorHandler.UpdateVendor)
	api.Delete("/vendors/:id", vendorHandler.DeleteVendor)

	// Layout routes
	api.Get("/layouts", layoutHandler.GetLayout)
	api.Get("/layouts/:id", layoutHandler.GetLayoutByID)
	api.Post("/layouts", layoutHandler.CreateLayout)
	api.Put("/layouts/:id", layoutHandler.UpdateLayout)
	api.Delete("/layouts/:id", layoutHandler.DeleteLayout)

	// Layout Position routes
	api.Get("/layout-positions", layoutPositionHandler.GetAllLayoutPositions)
	api.Get("/layout-positions/:id", layoutPositionHandler.GetLayoutPositionByID)
	api.Get("/layouts/:layoutId/layout-positions", layoutPositionHandler.GetLayoutPositionsByLayoutID)
	api.Post("/layouts/:layoutId/layout-positions", layoutPositionHandler.CreateLayoutPosition)
	api.Put("/layout-positions/:id", layoutPositionHandler.UpdateLayoutPosition)
	api.Delete("/layout-positions/:id", layoutPositionHandler.DeleteLayoutPosition)

	// Schedule routes
	api.Get("/schedules", scheduleHandler.GetAllSchedules)
	api.Get("/schedules/:id", scheduleHandler.GetScheduleByID)
	api.Get("/vendors/:vendorId/schedules", scheduleHandler.GetSchedulesByVendorID)
	api.Post("/schedules", scheduleHandler.CreateSchedule)
	api.Put("/schedules/:id", scheduleHandler.UpdateSchedule)
	api.Delete("/schedules/:id", scheduleHandler.DeleteSchedule)
}
