package entities

import (
	"github.com/google/uuid"
)

// ============================================================
// Schedule
// ============================================================
type Schedules struct {
	ScheduleID           uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"schedule_id"`
	VendorID             uuid.UUID  `gorm:"type:uuid;not null;index" json:"vendor_id"`
	ServiceTypeID        uuid.UUID  `gorm:"type:uuid;not null;index" json:"service_type_id"`
	OriginPoolID         uuid.UUID  `gorm:"type:uuid;not null;index" json:"origin_pool_id"`
	DestinationPoolID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"destination_pool_id"`
	LayoutID             uuid.UUID  `gorm:"type:uuid;not null;index" json:"layout_id"`
	VehicleType          string     `gorm:"type:varchar(100)" json:"vehicle_type"`
	DepartureDate        string     `gorm:"type:date;not null" json:"departure_date"`
	DepartureTime        string     `gorm:"type:time;not null" json:"departure_time"`
	EstimatedArrivalTime string     `gorm:"type:time" json:"estimated_arrival_time"`
	PricePerSeat         float64    `gorm:"type:decimal(12,2);not null" json:"price_per_seat"`
	TotalSeat            int        `gorm:"type:int;not null" json:"total_seat"`
	AvailableSeat        int        `gorm:"type:int;not null" json:"available_seat"`
	Status               string     `gorm:"type:varchar(50);default:'scheduled'" json:"status"`
	ActualDepartureTime  *string    `gorm:"type:time" json:"actual_departure_time,omitempty"`
	DepartedBy           *uuid.UUID `gorm:"type:uuid" json:"departed_by,omitempty"`
	CreatedBy            *uuid.UUID `gorm:"type:uuid" json:"created_by,omitempty"`

	// Relations
	// Vendor          Vendors
	Layout          Layouts
	ServiceType     ServiceTypes
	OriginPool      Pools `gorm:"foreignKey:OriginPoolID;references:PoolID" json:"origin_pool"`
	DestinationPool Pools `gorm:"foreignKey:DestinationPoolID;references:PoolID" json:"destination_pool"`
	// DepartedByUser  *Users `gorm:"foreignKey:DepartedBy;references:UserID" json:"departed_by_user,omitempty"`
	// CreatedByUser   *Users `gorm:"foreignKey:CreatedBy;references:UserID" json:"created_by_user,omitempty"`

	// // Relations — has many → gorm:"-"
	// Bookings     []Bookings     `gorm:"-" json:"bookings,omitempty"`
	// BookingSeats []BookingSeats `gorm:"-" json:"booking_seats,omitempty"`
	// SeatHolds    []SeatHolds    `gorm:"-" json:"seat_holds,omitempty"`

	BaseModel
}

func (Schedules) TableName() string { return "schedules" }
