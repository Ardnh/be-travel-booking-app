package routes

import (
	// middleware "github.com/ardnh/be-travel-booking-app/internal/interfaces/http/middleware"
	"github.com/ardnh/be-travel-booking-app/internal/interfaces/http/handlers"
	"github.com/ardnh/be-travel-booking-app/internal/interfaces/http/middleware"
	"github.com/ardnh/be-travel-booking-app/pkg/constants"
	"github.com/casbin/casbin/v3"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
)

func SetupAPIRoutes(app *fiber.App, log *logrus.Logger, enforcer *casbin.Enforcer, serviceTypeHandler *handlers.ServiceTypeHandler, poolPointHandler *handlers.PoolPointHandler, vendorHandler *handlers.VendorHandler, layoutHandler *handlers.LayoutHandler, layoutPositionHandler *handlers.LayoutPositionHandler, scheduleHandler *handlers.ScheduleHandler, authHandler *handlers.AuthHandler, userRolesHandler *handlers.UserRolesHandler, usersHandler *handlers.UsersHandler, bookingHandler *handlers.BookingHandler, validator *validator.Validate) {

	// Middleware
	casbinMw := middleware.NewCasbinMiddleware(enforcer, log)
	authMiddleware := middleware.NewAuthMiddleware()

	// API v1 group
	api := app.Group("/api/v1")

	// PUBLIC API
	// Auth routes
	api.Post("/auth/login", authHandler.Login)
	api.Post("/auth/register", authHandler.Register)

	// PRIVATE API
	// Service Type routes
	serviceTypes := api.Group("/service-types", authMiddleware.Authenticate())
	serviceTypes.Get("/", casbinMw.Authorize(constants.ResourceServiceTypes, constants.ActionRead), serviceTypeHandler.GetAllServiceTypes)
	serviceTypes.Get("/:id", casbinMw.Authorize(constants.ResourceServiceTypes, constants.ActionRead), serviceTypeHandler.GetServiceTypeByID)
	serviceTypes.Post("/", casbinMw.Authorize(constants.ResourceServiceTypes, constants.ActionCreate), serviceTypeHandler.CreateServiceType)
	serviceTypes.Put("/:id", casbinMw.Authorize(constants.ResourceServiceTypes, constants.ActionUpdate), serviceTypeHandler.UpdateServiceType)
	serviceTypes.Delete("/:id", casbinMw.Authorize(constants.ResourceServiceTypes, constants.ActionDelete), serviceTypeHandler.DeleteServiceType)

	// Pool Point routes
	poolPoints := api.Group("/pool-points", authMiddleware.Authenticate())
	poolPoints.Get("/", casbinMw.Authorize(constants.ResourcePoolPoints, constants.ActionRead), poolPointHandler.GetAllPoolPoints)
	poolPoints.Get("/:id", casbinMw.Authorize(constants.ResourcePoolPoints, constants.ActionRead), poolPointHandler.GetPoolPointByID)
	poolPoints.Post("/", casbinMw.Authorize(constants.ResourcePoolPoints, constants.ActionCreate), poolPointHandler.CreatePoolPoint)
	poolPoints.Put("/:id", casbinMw.Authorize(constants.ResourcePoolPoints, constants.ActionUpdate), poolPointHandler.UpdatePoolPoint)
	poolPoints.Delete("/:id", casbinMw.Authorize(constants.ResourcePoolPoints, constants.ActionDelete), poolPointHandler.DeletePoolPoint)

	poolPoints.Get("/vendors/:vendorId/pool-points", authMiddleware.Authenticate(), casbinMw.Authorize(constants.ResourcePoolPoints, constants.ActionRead), poolPointHandler.GetPoolPointsByVendorID)

	// Vendor routes
	vendors := api.Group("/vendors", authMiddleware.Authenticate())
	vendors.Get("/", casbinMw.Authorize(constants.ResourceVendors, constants.ActionRead), vendorHandler.GetAllVendors)
	vendors.Get("/:id", casbinMw.Authorize(constants.ResourceVendors, constants.ActionRead), vendorHandler.GetVendorByID)
	vendors.Get("/owner/:userId", casbinMw.Authorize(constants.ResourceVendors, constants.ActionRead), vendorHandler.GetVendorByOwnerUserID)
	vendors.Post("/", casbinMw.Authorize(constants.ResourceVendors, constants.ActionCreate), vendorHandler.CreateVendor)
	vendors.Put("/:id", casbinMw.Authorize(constants.ResourceVendors, constants.ActionUpdate), vendorHandler.UpdateVendor)
	vendors.Delete("/:id", casbinMw.Authorize(constants.ResourceVendors, constants.ActionDelete), vendorHandler.DeleteVendor)

	// Layout routes
	layouts := api.Group("/layouts", authMiddleware.Authenticate())
	layouts.Get("/", casbinMw.Authorize(constants.ResourceLayouts, constants.ActionRead), layoutHandler.GetLayout)
	layouts.Get("/:id", casbinMw.Authorize(constants.ResourceLayouts, constants.ActionRead), layoutHandler.GetLayoutByID)
	layouts.Post("/", casbinMw.Authorize(constants.ResourceLayouts, constants.ActionCreate), layoutHandler.CreateLayout)
	layouts.Put("/:id", casbinMw.Authorize(constants.ResourceLayouts, constants.ActionUpdate), layoutHandler.UpdateLayout)
	layouts.Delete("/:id", casbinMw.Authorize(constants.ResourceLayouts, constants.ActionDelete), layoutHandler.DeleteLayout)

	// Layout Position routes
	layoutPositions := api.Group("/layout-positions", authMiddleware.Authenticate())
	layoutPositions.Get("/", casbinMw.Authorize(constants.ResourceLayoutPositions, constants.ActionRead), layoutPositionHandler.GetAllLayoutPositions)
	layoutPositions.Get("/:id", casbinMw.Authorize(constants.ResourceLayoutPositions, constants.ActionRead), layoutPositionHandler.GetLayoutPositionByID)
	layoutPositions.Put("/:id", casbinMw.Authorize(constants.ResourceLayoutPositions, constants.ActionUpdate), layoutPositionHandler.UpdateLayoutPosition)
	layoutPositions.Delete("/:id", casbinMw.Authorize(constants.ResourceLayoutPositions, constants.ActionDelete), layoutPositionHandler.DeleteLayoutPosition)

	layouts.Get("/:layoutId/layout-positions", authMiddleware.Authenticate(), casbinMw.Authorize(constants.ResourceLayoutPositions, constants.ActionRead), layoutPositionHandler.GetLayoutPositionsByLayoutID)
	layouts.Post("/:layoutId/layout-positions", authMiddleware.Authenticate(), casbinMw.Authorize(constants.ResourceLayoutPositions, constants.ActionCreate), layoutPositionHandler.CreateLayoutPosition)

	// Schedule routes
	schedules := api.Group("/schedules", authMiddleware.Authenticate())
	schedules.Get("/", casbinMw.Authorize(constants.ResourceSchedules, constants.ActionRead), scheduleHandler.GetAllSchedules)
	schedules.Get("/:id", casbinMw.Authorize(constants.ResourceSchedules, constants.ActionRead), scheduleHandler.GetScheduleByID)
	schedules.Post("/", casbinMw.Authorize(constants.ResourceSchedules, constants.ActionCreate), scheduleHandler.CreateSchedule)
	schedules.Put("/:id", casbinMw.Authorize(constants.ResourceSchedules, constants.ActionUpdate), scheduleHandler.UpdateSchedule)
	schedules.Delete("/:id", casbinMw.Authorize(constants.ResourceSchedules, constants.ActionDelete), scheduleHandler.DeleteSchedule)

	vendors.Get("/:vendorId/schedules", authMiddleware.Authenticate(), casbinMw.Authorize(constants.ResourceSchedules, constants.ActionRead), scheduleHandler.GetSchedulesByVendorID)

	// User Roles routes
	userRoles := api.Group("/user-roles", authMiddleware.Authenticate())
	userRoles.Get("/:id", casbinMw.Authorize(constants.ResourceUserRoles, constants.ActionRead), userRolesHandler.GetUserRoleByID)
	userRoles.Post("/", casbinMw.Authorize(constants.ResourceUserRoles, constants.ActionCreate), userRolesHandler.CreateUserRole)
	userRoles.Put("/:id", casbinMw.Authorize(constants.ResourceUserRoles, constants.ActionUpdate), userRolesHandler.UpdateUserRole)
	userRoles.Delete("/:id", casbinMw.Authorize(constants.ResourceUserRoles, constants.ActionDelete), userRolesHandler.DeleteUserRole)

	api.Get("/users/:userId/user-roles", authMiddleware.Authenticate(), casbinMw.Authorize(constants.ResourceUserRoles, constants.ActionRead), userRolesHandler.GetUserRolesByUserID)
	vendors.Get("/:vendorId/user-roles", authMiddleware.Authenticate(), casbinMw.Authorize(constants.ResourceUserRoles, constants.ActionRead), userRolesHandler.GetUserRolesByVendorID)

	// Users routes
	users := api.Group("/users", authMiddleware.Authenticate())
	users.Post("/", usersHandler.CreateUser)
	users.Put("/:id", casbinMw.Authorize(constants.ResourceUsers, constants.ActionUpdate), usersHandler.UpdateUser)
	users.Delete("/:id", casbinMw.Authorize(constants.ResourceUsers, constants.ActionDelete), usersHandler.DeleteUser)
	users.Get("/profile", casbinMw.Authorize(constants.ResourceProfile, constants.ActionRead), usersHandler.GetUserProfile)

	// Booking routes
	bookings := api.Group("/bookings", authMiddleware.Authenticate())
	bookings.Get("/:id", casbinMw.Authorize(constants.ResourceBookings, constants.ActionRead), bookingHandler.GetBookingByID)
	bookings.Post("/", casbinMw.Authorize(constants.ResourceBookings, constants.ActionCreate), bookingHandler.CreateBooking)
	bookings.Put("/:id", casbinMw.Authorize(constants.ResourceBookings, constants.ActionUpdate), bookingHandler.UpdateBooking)
	bookings.Delete("/:id", casbinMw.Authorize(constants.ResourceBookings, constants.ActionDelete), bookingHandler.DeleteBooking)

	api.Get("/users/:userId/bookings", authMiddleware.Authenticate(), casbinMw.Authorize(constants.ResourceBookings, constants.ActionRead), bookingHandler.GetBookingsByUserID)
	api.Get("/schedules/:scheduleId/bookings", authMiddleware.Authenticate(), casbinMw.Authorize(constants.ResourceBookings, constants.ActionRead), bookingHandler.GetBookingsByScheduleID)
}
