package dto

import "github.com/ardnh/be-travel-booking-app/internal/domain/entities"

type CellType string
type CellStatus string
type WindowPosition string

const (
	CellTypeSeat   CellType = "seat"
	CellTypeDriver CellType = "driver"
	CellTypeAisle  CellType = "aisle"
)

const (
	StatusAvailable CellStatus = "available"
	StatusBooked    CellStatus = "booked"
	StatusBlocked   CellStatus = "blocked"
	StatusSelected  CellStatus = "selected"
)

// type CellDto struct {
// 	Type           CellType       `json:"type"`
// 	Row            int            `json:"row"`
// 	Col            int            `json:"col"`
// 	ID             string         `json:"id"`
// 	Status         CellStatus     `json:"status"`
// 	IsWindow       bool           `json:"isWindow"`
// 	WindowPosition WindowPosition `json:"windowPosition"`
// }

// type LayoutConfigDto [][]CellDto
type LayoutDTO struct {
	LayoutID     string                `json:"layout_id"`
	Name         string                `json:"name"`
	GridSizeX    int                   `json:"grid_size_x"`
	GridSizeY    int                   `json:"grid_size_y"`
	SeatCount    int                   `json:"seat_count"`
	CreatedBy    string                `json:"created_by"`
	LayoutConfig entities.LayoutConfig `json:"layout_config"`
}

type Pagination struct {
	CurrentPage int  `json:"current_page"`
	PageSize    int  `json:"page_size"`
	TotalItems  int  `json:"total_items"`
	TotalPages  int  `json:"total_pages"`
	HasNext     bool `json:"has_next"`
	HasPrevious bool `json:"has_previous"`
}

type CreateLayoutDTO struct {
	Name         string                `json:"name" validate:"required"`
	GridSizeX    int                   `json:"grid_size_x" validate:"required,min=1"`
	GridSizeY    int                   `json:"grid_size_y" validate:"required,min=1"`
	SeatCount    int                   `json:"seat_count" validate:"required,min=1"`
	CreatedBy    string                `json:"created_by"`
	LayoutConfig entities.LayoutConfig `json:"layout_config"`
}
