package dto

type LayoutPositionDTO struct {
	LayoutPositionID string `json:"layout_position_id"`
	LayoutID         string `json:"layout_id"`
	Label            string `json:"label"`
	Row              int    `json:"row"`
	Col              int    `json:"col"`
	PositionType     string `json:"position_type"`
	IsUsed           bool   `json:"is_used"`
}

type CreateLayoutPositionDTO struct {
	LayoutID     string `json:"layout_id" validate:"required"`
	Label        string `json:"label" validate:"required"`
	Row          int    `json:"row" validate:"required,min=1"`
	Col          int    `json:"col" validate:"required,min=1"`
	PositionType string `json:"position_type" validate:"required"`
	IsUsed       bool   `json:"is_used"`
}

type UpdateLayoutPositionDTO struct {
	Label        *string `json:"label,omitempty"`
	Row          *int    `json:"row,omitempty"`
	Col          *int    `json:"col,omitempty"`
	PositionType *string `json:"position_type,omitempty"`
	IsUsed       *bool   `json:"is_used,omitempty"`
}
