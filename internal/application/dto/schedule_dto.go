package dto

import "time"

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

type CreateScheduleDTO struct {
	VendorID             string   `json:"vendor_id" validate:"required"`
	ServiceTypeID        string   `json:"service_type_id" validate:"required"`
	OriginPoolID         string   `json:"origin_pool_id" validate:"required"`
	DestinationPoolID    string   `json:"destination_pool_id" validate:"required"`
	LayoutID             string   `json:"layout_id" validate:"required"`
	VehicleType          string   `json:"vehicle_type" validate:"required"`
	DepartureDate        string   `json:"departure_date" validate:"required"`
	DepartureTime        string   `json:"departure_time" validate:"required"`
	EstimatedArrivalTime *string  `json:"estimated_arrival_time"`
	PricePerSeat         *float64 `json:"price_per_seat" validate:"required"`
	TotalSeat            *int     `json:"total_seat" validate:"required"`
	AvailableSeat        *int     `json:"available_seat"`
	Status               *string  `json:"status"`
}

type UpdateScheduleDTO struct {
	VehicleType          *string  `json:"vehicle_type"`
	DepartureDate        *string  `json:"departure_date"`
	DepartureTime        *string  `json:"departure_time"`
	EstimatedArrivalTime *string  `json:"estimated_arrival_time"`
	PricePerSeat         *float64 `json:"price_per_seat"`
	TotalSeat            *int     `json:"total_seat"`
	AvailableSeat        *int     `json:"available_seat"`
	Status               *string  `json:"status"`
	ActualDepartureTime *string  `json:"actual_departure_time"`
	DepartedBy           *string  `json:"departed_by"`
}
