package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ardnh/be-travel-booking-app/internal/application/services"
	"github.com/ardnh/be-travel-booking-app/internal/config"
	"github.com/ardnh/be-travel-booking-app/internal/infrastructure/database/postgresql"
	"github.com/ardnh/be-travel-booking-app/internal/infrastructure/database/redis"
	"github.com/ardnh/be-travel-booking-app/internal/infrastructure/database/seeder"
	"github.com/ardnh/be-travel-booking-app/internal/infrastructure/repositories"
	"github.com/ardnh/be-travel-booking-app/internal/interfaces/http/handlers"
	"github.com/ardnh/be-travel-booking-app/internal/interfaces/http/middleware"
	"github.com/ardnh/be-travel-booking-app/internal/interfaces/http/routes"
	casbin_utils "github.com/ardnh/be-travel-booking-app/internal/utils/casbin"
	logger "github.com/ardnh/be-travel-booking-app/internal/utils/logger" // tetap sama!
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func main() {

	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory:", err)
	}

	modelPath := filepath.Join(workDir, "internal/config", "casbin_model.conf")

	// TODO: Implement application initialization
	// 1. Load configuration
	logger := logger.New()
	cfg := config.LoadConfig()
	validator := validator.New()

	// 2. Initialize database
	db, err := postgresql.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Seed initial data
	if err := seeder.Seed(db); err != nil {
		log.Fatalf("❌ Failed to seed database: %v", err)
	}

	defer postgresql.CloseDB(db)

	redisDb := redis.NewRedisDB(cfg)
	defer redisDb.Close()

	// 4. Wire up dependencies
	// Casbin
	enforcer, err := casbin_utils.InitCasbin(modelPath, db)
	// Seed rules (jalankan sekali, atau cek dulu apakah sudah ada)
	casbin_utils.SeedCasbinRules(enforcer)

	// Repository
	serviceTypeRepo := repositories.NewServiceTypeRepository(db, redisDb)
	poolPointRepo := repositories.NewPoolPointRepository(db, redisDb)
	vendorRepo := repositories.NewVendorRepository(db, redisDb)
	layoutRepo := repositories.NewLayoutRepository(db, redisDb)
	layoutPositionRepo := repositories.NewLayoutPositionRepository(db, redisDb)
	scheduleRepo := repositories.NewScheduleRepository(db, redisDb)
	userRepo := repositories.NewUsersRepository(db, redisDb)
	userRolesRepo := repositories.NewUserRolesRepository(db, redisDb)
	bookingRepo := repositories.NewBookingRepository(db, redisDb)

	// Service
	serviceTypeService := services.NewServiceTypeServiceImpl(serviceTypeRepo, logger)
	poolPointService := services.NewPoolPointServiceImpl(poolPointRepo, logger)
	vendorService := services.NewVendorServiceImpl(vendorRepo, logger)
	layoutService := services.NewLayoutServiceImpl(layoutRepo, logger)
	layoutPositionService := services.NewLayoutPositionServiceImpl(layoutPositionRepo, logger)
	scheduleService := services.NewScheduleServiceImpl(scheduleRepo, logger)
	authService := services.NewAuthService(userRepo, userRolesRepo, logger, cfg, enforcer)
	userRolesService := services.NewUserRolesServiceImpl(userRolesRepo, logger)
	usersService := services.NewUsersServiceImpl(userRepo, userRolesRepo, enforcer, logger)
	bookingService := services.NewBookingServiceImpl(bookingRepo, logger)

	// Handler
	serviceTypeHandler := handlers.NewServiceTypeHandler(serviceTypeService, validator, logger)
	poolPointHandler := handlers.NewPoolPointHandler(poolPointService, validator, logger)
	vendorHandler := handlers.NewVendorHandler(vendorService, validator, logger)
	layoutHandler := handlers.NewLayoutHandler(layoutService, validator, logger)
	layoutPositionHandler := handlers.NewLayoutPositionHandler(layoutPositionService, validator, logger)
	scheduleHandler := handlers.NewScheduleHandler(scheduleService, validator, logger)
	authHandler := handlers.NewAuthHandler(authService, validator, logger)
	userRolesHandler := handlers.NewUserRolesHandler(userRolesService, validator, logger)
	usersHandler := handlers.NewUsersHandler(usersService, validator, logger)
	bookingHandler := handlers.NewBookingHandler(bookingService, validator, logger)

	// 5. Start HTTP server
	app := fiber.New()

	requestTimer := middleware.NewRequestTimerMiddleware(logger)
	app.Use(requestTimer.Track())

	routes.SetupAPIRoutes(
		app,
		logger,
		enforcer,
		serviceTypeHandler,
		poolPointHandler,
		vendorHandler,
		layoutHandler,
		layoutPositionHandler,
		scheduleHandler,
		authHandler,
		userRolesHandler,
		usersHandler,
		bookingHandler,
		validator,
	)

	portListen := fmt.Sprintf(":%s", cfg.App.Port)
	if err := app.Listen(portListen); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
	os.Exit(0)
}
