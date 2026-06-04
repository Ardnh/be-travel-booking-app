package dto

type LayoutDTO struct {
	LayoutID        string               `json:"layout_id"`
	Name            string               `json:"name"`
	GridSizeX       int                  `json:"grid_size_x"`
	GridSizeY       int                  `json:"grid_size_y"`
	SeatCount       int                  `json:"seat_count"`
	CreatedBy       string               `json:"created_by"`
	LayoutPositions []*LayoutPositionDTO `json:"layout_positions"`
}

type Pagination struct {
	CurrentPage int  `json:"current_page"`
	PageSize    int  `json:"page_size"`
	TotalItems  int  `json:"total_items"`
	TotalPages  int  `json:"total_pages"`
	HasNext     bool `json:"has_next"`
	HasPrevious bool `json:"has_previous"`
}

// type LayoutPositionDTO struct {
// 	LayoutPositionID string `json:"layout_position_id"`
// 	LayoutID         string `json:"layout_id"`
// 	Label            string `json:"label"`
// 	Row              int    `json:"row"`
// 	Col              int    `json:"col"`
// 	PositionType     string `json:"positionType"`
// 	IsUsed           bool   `json:"isUsed"`
// }

type CreateLayoutDTO struct {
	Name            string                    `json:"name" validate:"required"`
	GridSizeX       int                       `json:"grid_size_x" validate:"required,min=1"`
	GridSizeY       int                       `json:"grid_size_y" validate:"required,min=1"`
	SeatCount       int                       `json:"seat_count" validate:"required,min=1"`
	CreatedBy       string                    `json:"created_by"`
	LayoutPositions []CreateLayoutPositionDTO `json:"layout_positions"`
}

// type CreateLayoutPositionDTO struct {
// 	LayoutDTO    string `json:"layoutDTO" validate:"required"`
// 	Label        string `json:"label" validate:"required"`
// 	Row          int    `json:"row" validate:"required,min=1"`
// 	Col          int    `json:"column" validate:"required,min=1"`
// 	PositionType string `json:"position_type" validate:"required"`
// 	IsUsed       bool   `json:"isUsed"`
// }
