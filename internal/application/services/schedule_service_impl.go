package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	helpers "github.com/ardnh/be-travel-booking-app/internal/utils/helpers"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (s *ScheduleServiceImpl) GetAllSchedules(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.Schedules, int64, error) {
	schedules, total, err := s.scheduleRepository.GetAllSchedules(ctx, page, pageSize, search, sortBy, sortOrder)
	if err != nil {
		return nil, 0, err
	}
	return schedules, total, nil
}

func (s *ScheduleServiceImpl) GetSchedulesByVendorID(ctx context.Context, vendorID uuid.UUID, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.Schedules, int64, error) {
	schedules, total, err := s.scheduleRepository.GetSchedulesByVendorID(ctx, vendorID, page, pageSize, search, sortBy, sortOrder)
	if err != nil {
		return nil, 0, err
	}
	return schedules, total, nil
}

const maxGeneratedRows = 1000

func (s *ScheduleServiceImpl) CreateSchedule(ctx context.Context, vendorId uuid.UUID, req dto.CreateScheduleDTO) (*dto.GenerateScheduleResultDTO, error) {

	originID, err := uuid.Parse(req.OriginPoolID)
	if err != nil {
		return nil, errorConst.ErrBadRequest
	}
	serviceTypeID, _ := uuid.Parse(req.ServiceTypeID)
	layoutID, _ := uuid.Parse(req.LayoutID)

	destIDs := make([]uuid.UUID, 0, len(req.DestinationPoolIDs))
	for _, id := range req.DestinationPoolIDs {
		u, err := uuid.Parse(id)
		if err != nil {
			return nil, errorConst.ErrBadRequest
		}
		destIDs = append(destIDs, u)
	}

	dates := helpers.ExpandDates(req.ValidFrom, req.ValidTo, req.DaysOfWeek)
	if len(dates) == 0 {
		return nil, errorConst.ErrBadRequest
	}

	total := len(dates) * len(destIDs) * len(req.DepartureTimes)
	if total > maxGeneratedRows {
		return nil, errors.New(fmt.Sprintf(
			"would generate %d schedules, max is %d — narrow the date range", total, maxGeneratedRows))
	}

	batchID := uuid.New()
	result := &dto.GenerateScheduleResultDTO{BatchID: batchID.String()}
	err = s.repo.WithTx(func(tx *gorm.DB) error {
		// ── validasi kepemilikan, semua di dalam tx ──
		ok, err := s.repo.ValidateServiceType(tx, vendorID, serviceTypeID)
		if err != nil {
			return err
		}
		if !ok {
			return apperr.Forbidden("service_type does not belong to this vendor")
		}

		allPools := append([]uuid.UUID{originID}, destIDs...)
		found, err := s.repo.ValidatePoolOwnership(tx, vendorID, allPools)
		if err != nil {
			return err
		}
		if found != len(allPools) {
			return apperr.Forbidden("one or more pools do not belong to this vendor")
		}

		seatCount, err := s.repo.CountLayoutSeats(tx, layoutID, vendorID)
		if err != nil {
			return err
		}
		if seatCount == 0 {
			return apperr.BadRequest("layout not found, empty, or not owned by this vendor")
		}

		// ── build rows ──
		rows := make([]models.Schedules, 0, total)
		for _, date := range dates {
			for _, destID := range destIDs {
				for _, t := range d.DepartureTimes {
					price, ok := resolvePrice(t, d.PriceBands)
					if !ok {
						return apperr.BadRequest("no price band matches departure_time " + t)
					}

					var eta *string
					if d.DurationMinutes != nil {
						eta = ptr(addMinutes(t, *d.DurationMinutes))
					}

					rows = append(rows, models.Schedules{
						VendorID:             vendorID,
						ServiceTypeID:        serviceTypeID,
						OriginPoolID:         originID,
						DestinationPoolID:    destID,
						LayoutID:             layoutID,
						ScheduleBatchID:      &batchID,
						VehicleType:          d.VehicleType,
						DepartureDate:        date,
						DepartureTime:        t,
						EstimatedArrivalTime: eta,
						PricePerSeat:         price,
						TotalSeat:            seatCount,
						AvailableSeat:        seatCount,
						Status:               "scheduled",
						CreatedBy:            &userID,
					})
				}
			}
		}

		affected, err := s.repo.BulkInsert(tx, rows, d.OverwriteExisting)
		if err != nil {
			return err
		}

		result.Created = int(affected)
		result.Skipped = total - int(affected)
		if result.Skipped > 0 && !d.OverwriteExisting {
			result.Warnings = append(result.Warnings, fmt.Sprintf(
				"%d schedules already existed and were skipped", result.Skipped))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
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
