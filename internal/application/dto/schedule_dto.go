package dto

import (
	"time"
)

type PoolSummaryDTO struct {
	PoolID    string  `json:"pool_id"`
	Name      string  `json:"name"`
	City      string  `json:"city"`
	Province  string  `json:"province"`
	Address   string  `json:"address"`
	OpenTime  string  `json:"open_time"`
	CloseTime string  `json:"close_time"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ScheduleDTO struct {
	ScheduleID           string          `json:"schedule_id"`
	VendorID             string          `json:"vendor_id"`
	ServiceTypeID        string          `json:"service_type_id"`
	LayoutID             string          `json:"layout_id"`
	VehicleType          string          `json:"vehicle_type"`
	DepartureDate        string          `json:"departure_date"`
	DepartureTime        string          `json:"departure_time"`
	EstimatedArrivalTime string          `json:"estimated_arrival_time"`
	ActualDepartureTime  string          `json:"actual_departure_time"`
	PricePerSeat         float64         `json:"price_per_seat"`
	TotalSeat            int             `json:"total_seat"`
	AvailableSeat        int             `json:"available_seat"`
	Status               string          `json:"status"`
	DepartedBy           string          `json:"departed_by"`
	CreatedBy            string          `json:"created_by"`
	OriginPool           *PoolSummaryDTO `json:"origin_pool"`
	DestinationPool      *PoolSummaryDTO `json:"destination_pool"`
	CreatedAt            time.Time       `json:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at"`
}

// type CreateScheduleDTO struct {
// 	VendorID             string   `json:"vendor_id" validate:"required"`
// 	ServiceTypeID        string   `json:"service_type_id" validate:"required"`
// 	OriginPoolID         string   `json:"origin_pool_id" validate:"required"`
// 	DestinationPoolID    string   `json:"destination_pool_id" validate:"required"`
// 	LayoutID             string   `json:"layout_id" validate:"required"`
// 	VehicleType          string   `json:"vehicle_type" validate:"required"`
// 	DepartureDate        string   `json:"departure_date" validate:"required"`
// 	DepartureTime        string   `json:"departure_time" validate:"required"`
// 	CreatedBy            string   `json:"created_by" validate:"required"`
// 	EstimatedArrivalTime *string  `json:"estimated_arrival_time"`
// 	PricePerSeat         *float64 `json:"price_per_seat" validate:"required"`
// 	TotalSeat            *int     `json:"total_seat" validate:"required"`
// 	AvailableSeat        *int     `json:"available_seat"`
// 	Status               *string  `json:"status"`
// }

type UpdateScheduleDTO struct {
	VehicleType          *string  `json:"vehicle_type"`
	DepartureDate        *string  `json:"departure_date"`
	DepartureTime        *string  `json:"departure_time"`
	EstimatedArrivalTime *string  `json:"estimated_arrival_time"`
	PricePerSeat         *float64 `json:"price_per_seat"`
	TotalSeat            *int     `json:"total_seat"`
	AvailableSeat        *int     `json:"available_seat"`
	Status               *string  `json:"status"`
	ActualDepartureTime  *string  `json:"actual_departure_time"`
	DepartedBy           *string  `json:"departed_by"`
}

// ── Bulk generate (dari FE wizard) ──────────────────────────
type PriceBandDTO struct {
	Label string  `json:"label" validate:"required"`
	From  string  `json:"from"  validate:"required,datetime=15:04"`
	To    string  `json:"to"    validate:"required,datetime=15:04"`
	Price float64 `json:"price" validate:"required,gt=0"`
}

type CreateScheduleDTO struct {
	ServiceTypeID      string         `json:"service_type_id"      validate:"required,uuid4"`
	OriginPoolID       string         `json:"origin_pool_id"       validate:"required,uuid4"`
	DestinationPoolIDs []string       `json:"destination_pool_ids" validate:"required,min=1,max=20,unique,dive,uuid4"`
	LayoutID           string         `json:"layout_id"            validate:"required,uuid4"`
	VehicleType        string         `json:"vehicle_type"         validate:"omitempty,max=100"`
	ValidFrom          string         `json:"valid_from"           validate:"required,datetime=2006-01-02"`
	ValidTo            string         `json:"valid_to"             validate:"required,datetime=2006-01-02"`
	DaysOfWeek         []int          `json:"days_of_week"         validate:"required,min=1,max=7,unique,dive,min=0,max=6"`
	DepartureTimes     []string       `json:"departure_times"      validate:"required,min=1,max=24,unique,dive,datetime=15:04"`
	DurationMinutes    *int           `json:"duration_minutes"     validate:"omitempty,min=1,max=2880"`
	PriceBands         []PriceBandDTO `json:"price_bands"          validate:"required,min=1,dive"`
	OverwriteExisting  bool           `json:"overwrite_existing"`
}

type CreateSchedule struct {
	CreateScheduleDTO
}

// func (d *GenerateScheduleDTO) Validate() error {
// 	from, _ := time.Parse("2006-01-02", d.ValidFrom)
// 	to, _ := time.Parse("2006-01-02", d.ValidTo)

// 	if to.Before(from) {
// 		return errors.New("valid_to must be >= valid_from")
// 	}
// 	if to.Sub(from) > 90*24*time.Hour {
// 		return errors.New("range must not exceed 90 days")
// 	}
// 	if from.Before(time.Now().Truncate(24 * time.Hour)) {
// 		return errors.New("valid_from must not be in the past")
// 	}
// 	if slices.Contains(d.DestinationPoolIDs, d.OriginPoolID) {
// 		return errors.New("destination must differ from origin")
// 	}
// 	// setiap departure_time harus match tepat satu band
// 	for _, t := range d.DepartureTimes {
// 		n := 0
// 		for _, b := range d.PriceBands {
// 			if matchBand(t, b) {
// 				n++
// 			}
// 		}
// 		if n != 1 {
// 			return fmt.Errorf("departure_time %s matches %d price bands, expected 1", t, n)
// 		}
// 	}
// 	return nil
// }
