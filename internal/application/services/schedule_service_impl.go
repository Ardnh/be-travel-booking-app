package services

import (
	"context"
	"errors"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/application/mapper"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ScheduleServiceImpl struct {
	scheduleRepository repositories.ScheduleRepository
	log                *logrus.Logger
}

func NewScheduleServiceImpl(scheduleRepository repositories.ScheduleRepository, log *logrus.Logger) *ScheduleServiceImpl {
	return &ScheduleServiceImpl{
		scheduleRepository: scheduleRepository,
		log:                log,
	}
}

func (s *ScheduleServiceImpl) GetScheduleByID(ctx context.Context, scheduleID uuid.UUID) (*entities.Schedules, error) {
	schedule, err := s.scheduleRepository.GetScheduleByID(ctx, scheduleID)
	if err != nil {
		if errors.Is(err, errorConst.ErrNotFound) {
			s.log.WithFields(logrus.Fields{
				"schedule_id": scheduleID,
				"error":       err,
			}).Error("schedule not found")
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return schedule, nil
}

func (s *ScheduleServiceImpl) GetAllSchedules(ctx context.Context) ([]entities.Schedules, error) {
	schedules, err := s.scheduleRepository.GetAllSchedules(ctx)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func (s *ScheduleServiceImpl) GetSchedulesByVendorID(ctx context.Context, vendorID uuid.UUID) ([]entities.Schedules, error) {
	schedules, err := s.scheduleRepository.GetSchedulesByVendorID(ctx, vendorID)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func (s *ScheduleServiceImpl) CreateSchedule(ctx context.Context, req dto.CreateScheduleDTO) error {
	schedule, err := mapper.CreateScheduleDTOToEntity(req)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"vendor_id":       req.VendorID,
			"service_type_id": req.ServiceTypeID,
			"error":           err,
		}).Error("failed to map schedule dto to entity")
		return errorConst.ErrInternalServer
	}

	err = s.scheduleRepository.CreateSchedule(ctx, schedule)
	if err != nil {
		return err
	}
	return nil
}

func (s *ScheduleServiceImpl) UpdateSchedule(ctx context.Context, scheduleID uuid.UUID, req dto.UpdateScheduleDTO) error {
	schedule, err := s.scheduleRepository.GetScheduleByID(ctx, scheduleID)
	if err != nil {
		if errors.Is(err, errorConst.ErrNotFound) {
			s.log.WithFields(logrus.Fields{
				"schedule_id": scheduleID,
				"error":       err,
			}).Error("schedule not found")
			return errorConst.ErrNotFound
		}
		return err
	}

	if req.VehicleType != nil {
		schedule.VehicleType = *req.VehicleType
	}
	if req.DepartureDate != nil {
		schedule.DepartureDate = *req.DepartureDate
	}
	if req.DepartureTime != nil {
		schedule.DepartureTime = *req.DepartureTime
	}
	if req.EstimatedArrivalTime != nil {
		schedule.EstimatedArrivalTime = *req.EstimatedArrivalTime
	}
	if req.PricePerSeat != nil {
		schedule.PricePerSeat = *req.PricePerSeat
	}
	if req.TotalSeat != nil {
		schedule.TotalSeat = *req.TotalSeat
	}
	if req.AvailableSeat != nil {
		schedule.AvailableSeat = *req.AvailableSeat
	}
	if req.Status != nil {
		schedule.Status = *req.Status
	}
	if req.ActualDepartureTime != nil {
		schedule.ActualDepartureTime = req.ActualDepartureTime
	}
	if req.DepartedBy != nil {
		departedByID, err := uuid.Parse(*req.DepartedBy)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"departed_by": *req.DepartedBy,
				"error":       err,
			}).Error("failed to parse departed_by uuid")
			return errorConst.ErrInternalServer
		}
		schedule.DepartedBy = &departedByID
	}

	err = s.scheduleRepository.UpdateSchedule(ctx, *schedule)
	if err != nil {
		return err
	}
	return nil
}

func (s *ScheduleServiceImpl) DeleteSchedule(ctx context.Context, scheduleID uuid.UUID) error {
	err := s.scheduleRepository.DeleteSchedule(ctx, scheduleID)
	if err != nil {
		return err
	}
	return nil
}
