package entities

import (
	"time"

	"github.com/google/uuid"
)

type Schedule struct {
	ScheduleID           uuid.UUID  `gorm:"column:schedule_id;type:uuid;primaryKey"`
	VendorID             uuid.UUID  `gorm:"column:vendor_id;type:uuid;not null;foreignKey:VendorID;references:VendorID"`
	ServiceTypeID        uuid.UUID  `gorm:"column:service_type_id;type:uuid;not null;foreignKey:ServiceTypeID;references:ServiceTypeID"`
	OriginPoolID         uuid.UUID  `gorm:"column:origin_pool_id;type:uuid;not null;foreignKey:OriginPoolID;references:PoolID"`
	DestinationPoolID    uuid.UUID  `gorm:"column:destination_pool_id;type:uuid;not null;foreignKey:DestinationPoolID;references:PoolID"`
	LayoutID             uuid.UUID  `gorm:"column:layout_id;type:uuid;not null;foreignKey:LayoutID;references:LayoutID"`
	VehicleType          string     `gorm:"column:vehicle_type;type:varchar;not null"`
	DepartureDate        time.Time  `gorm:"column:departure_date;type:date;not null"`
	DepartureTime        string     `gorm:"column:departure_time;type:time;not null"`
	EstimatedArrivalTime *string    `gorm:"column:estimated_arrival_time;type:time;null"`
	PricePerSeat         float64    `gorm:"column:price_per_seat;type:decimal(12,2);not null"`
	TotalSeat            int        `gorm:"column:total_seat;type:int;not null"`
	AvailableSeat        int        `gorm:"column:available_seat;type:int;not null"`
	Status               string     `gorm:"column:status;type:varchar;default:'open'"`
	ActualDepartureTime  *time.Time `gorm:"column:actual_departure_time;type:timestamp;null"`
	DepartedBy           *uuid.UUID `gorm:"column:departed_by;type:uuid;null"`
	CreatedBy            uuid.UUID  `gorm:"column:created_by;type:uuid;not null"`
	CreatedAt            time.Time  `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;autoCreateTime:milli"`
	UpdatedAt            time.Time  `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime:milli"`
}

func (Schedule) TableName() string {
	return "schedules"
}
