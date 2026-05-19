package mapper

import (
	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
	"strings"
)

func ScheduleToDTO(schedule *entities.Schedules) *dto.ScheduleDTO {
	estimatedArrival := ""
	if schedule.EstimatedArrivalTime != nil {
		estimatedArrival = *schedule.EstimatedArrivalTime
	}
	actualDeparture := ""
	if schedule.ActualDepartureTime != nil {
		actualDeparture = *schedule.ActualDepartureTime
	}
	departedBy := ""
	if schedule.DepartedBy != nil {
		departedBy = schedule.DepartedBy.String()
	}
	createdBy := ""
	if schedule.CreatedBy != nil {
		createdBy = schedule.CreatedBy.String()
	}

	originPool := dto.PoolSummaryDTO{
		PoolID:    schedule.OriginPool.PoolID.String(),
		Name:      schedule.OriginPool.Name,
		City:      schedule.OriginPool.City,
		Province:  schedule.OriginPool.Province,
		Address:   schedule.OriginPool.Address,
		OpenTime:  schedule.OriginPool.OpenTime,
		CloseTime: schedule.OriginPool.CloseTime,
	}

	destPool := dto.PoolSummaryDTO{
		PoolID:    schedule.DestinationPool.PoolID.String(),
		Name:      schedule.DestinationPool.Name,
		City:      schedule.DestinationPool.City,
		Province:  schedule.DestinationPool.Province,
		Address:   schedule.DestinationPool.Address,
		OpenTime:  schedule.DestinationPool.OpenTime,
		CloseTime: schedule.DestinationPool.CloseTime,
	}

	return &dto.ScheduleDTO{
		ScheduleID:           schedule.ScheduleID.String(),
		VendorID:             schedule.VendorID.String(),
		ServiceTypeID:        schedule.ServiceTypeID.String(),
		LayoutID:             schedule.LayoutID.String(),
		VehicleType:          schedule.VehicleType,
		DepartureDate:        schedule.DepartureDate,
		DepartureTime:        schedule.DepartureTime,
		EstimatedArrivalTime: estimatedArrival,
		ActualDepartureTime:  actualDeparture,
		PricePerSeat:         schedule.PricePerSeat,
		TotalSeat:            schedule.TotalSeat,
		AvailableSeat:        schedule.AvailableSeat,
		Status:               schedule.Status,
		DepartedBy:           departedBy,
		CreatedBy:            createdBy,
		OriginPool:           &originPool,
		DestinationPool:      &destPool,
		CreatedAt:            schedule.CreatedAt,
		UpdatedAt:            schedule.UpdatedAt,
	}
}

func SchedulesToDTO(schedules []entities.Schedules) []*dto.ScheduleDTO {
	result := make([]*dto.ScheduleDTO, 0, len(schedules))
	for _, s := range schedules {
		result = append(result, ScheduleToDTO(&s))
	}
	return result
}

func CreateScheduleDTOToEntity(req dto.CreateScheduleDTO, createdBy uuid.UUID) (entities.Schedules, error) {
	layoutID, err := uuid.Parse(req.LayoutID)
	if err != nil {
		return entities.Schedules{}, err
	}
	vendorID, err := uuid.Parse(req.VendorID)
	if err != nil {
		return entities.Schedules{}, err
	}
	serviceTypeID, err := uuid.Parse(req.ServiceTypeID)
	if err != nil {
		return entities.Schedules{}, err
	}
	originPoolID, err := uuid.Parse(req.OriginPoolID)
	if err != nil {
		return entities.Schedules{}, err
	}
	destPoolID, err := uuid.Parse(req.DestinationPoolID)
	if err != nil {
		return entities.Schedules{}, err
	}

	price := float64(0)
	if req.PricePerSeat != nil {
		price = *req.PricePerSeat
	}
	totalSeat := 0
	if req.TotalSeat != nil {
		totalSeat = *req.TotalSeat
	}

	availableSeat := totalSeat
	if req.AvailableSeat != nil {
		availableSeat = *req.AvailableSeat
	}

	status := "scheduled"
	if req.Status != nil && *req.Status != "" {
		status = *req.Status
		status = strings.ToLower(strings.TrimSpace(status))
	}

	return entities.Schedules{
		LayoutID:          layoutID,
		VendorID:          vendorID,
		ServiceTypeID:     serviceTypeID,
		OriginPoolID:      originPoolID,
		DestinationPoolID: destPoolID,
		DepartureDate:     req.DepartureDate,
		DepartureTime:     req.DepartureTime,
		PricePerSeat:      price,
		TotalSeat:         totalSeat,
		AvailableSeat:     availableSeat,
		Status:            status,
		CreatedBy:         &createdBy,
	}, nil
}
