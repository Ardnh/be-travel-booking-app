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

type LayoutPositionDTO struct {
	LayoutPositionID string `json:"layout_position_id"`
	LayoutID         string `json:"layout_id"`
	Label            string `json:"label"`
	Row              int    `json:"row"`
	Col              int    `json:"col"`
	PositionType     string `json:"positionType"`
	IsUsed           bool   `json:"isUsed"`
}

type CreateLayoutDTO struct {
	Name            string                    `json:"name" validate:"required"`
	GridSizeX       int                       `json:"gridSizeX" validate:"required,min=1"`
	GridSizeY       int                       `json:"gridSizeY" validate:"required,min=1"`
	SeatCount       int                       `json:"seatCount" validate:"required,min=1"`
	CreatedBy       string                    `json:"createdBy" validate:"required"`
	LayoutPositions []CreateLayoutPositionDTO `json:"layoutPositions"`
}

type CreateLayoutPositionDTO struct {
	Label    string `json:"label" validate:"required"`
	Row      int    `json:"row" validate:"required,min=1"`
	Col      int    `json:"column" validate:"required,min=1"`
	SeatType string `json:"seatType" validate:"required"`
	IsUsed   bool   `json:"isUsed"`
}
