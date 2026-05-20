package dto

type BookingDTO struct {
	BookingID        string  `json:"booking_id"`
	BookingCode      string  `json:"booking_code"`
	UserID           string  `json:"user_id"`
	ScheduleID       string  `json:"schedule_id"`
	PassengerName    string  `json:"passenger_name"`
	PassengerPhone   string  `json:"passenger_phone"`
	PickupAddress    string  `json:"pickup_address"`
	DropoffAddress   string  `json:"dropoff_address"`
	PricePerSeat     float64 `json:"price_per_seat"`
	SeatCount        int     `json:"seat_count"`
	TotalPrice       float64 `json:"total_price"`
	PaymentStatus    string  `json:"payment_status"`
	PaymentReference string  `json:"payment_reference"`
	PaymentMethod    string  `json:"payment_method"`
	BookingStatus    string  `json:"booking_status"`
}

type CreateBookingDTO struct {
	UserID         string  `json:"user_id" validate:"required"`
	ScheduleID     string  `json:"schedule_id" validate:"required"`
	PassengerName  string  `json:"passenger_name" validate:"required"`
	PassengerPhone string  `json:"passenger_phone"`
	PickupAddress  string  `json:"pickup_address"`
	DropoffAddress string  `json:"dropoff_address"`
	SeatCount      int     `json:"seat_count" validate:"required,min=1"`
	PricePerSeat   float64 `json:"price_per_seat" validate:"required,min=0"`
}

type UpdateBookingDTO struct {
	PassengerName  *string `json:"passenger_name,omitempty"`
	PassengerPhone *string `json:"passenger_phone,omitempty"`
	PickupAddress  *string `json:"pickup_address,omitempty"`
	DropoffAddress *string `json:"dropoff_address,omitempty"`
	SeatCount      *int    `json:"seat_count,omitempty"`
	PricePerSeat   *float64 `json:"price_per_seat,omitempty"`
	PaymentMethod  *string `json:"payment_method,omitempty"`
	Notes          *string `json:"notes,omitempty"`
}