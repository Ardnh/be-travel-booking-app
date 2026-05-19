package repositories

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

type ScheduleRepository interface {
	GetScheduleByID(ctx context.Context, scheduleID uuid.UUID) (*entities.Schedules, error)
	GetAllSchedules(ctx context.Context) ([]entities.Schedules, error)
	GetSchedulesByVendorID(ctx context.Context, vendorID uuid.UUID) ([]entities.Schedules, error)
	CreateSchedule(ctx context.Context, schedule entities.Schedules) error
	UpdateSchedule(ctx context.Context, schedule entities.Schedules) error
	DeleteSchedule(ctx context.Context, scheduleID uuid.UUID) error
}
