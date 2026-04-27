package mapper

import (
	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
)

func LayoutsToDTO(layouts []*entities.Layout) []*dto.LayoutDTO {
	result := make([]*dto.LayoutDTO, 0, len(layouts))
	for _, layout := range layouts {
		result = append(result, LayoutToDTO(layout))
	}
	return result
}

func LayoutToDTO(layout *entities.Layout) *dto.LayoutDTO {
	return &dto.LayoutDTO{
		LayoutID:        layout.LayoutID.String(),
		Name:            layout.Name,
		GridSizeX:       layout.GridSizeX,
		GridSizeY:       layout.GridSizeY,
		SeatCount:       layout.SeatCount,
		CreatedBy:       layout.CreatedBy.String(),
		LayoutPositions: LayoutPositionsToDTO(layout.LayoutPositions),
	}
}

func LayoutPositionsToDTO(layoutPositions []entities.LayoutPosition) []*dto.LayoutPositionDTO {
	result := make([]*dto.LayoutPositionDTO, 0, len(layoutPositions))
	for _, position := range layoutPositions {
		result = append(result, &dto.LayoutPositionDTO{
			LayoutPositionID: position.LayoutPositionID.String(),
			LayoutID:         position.LayoutID.String(),
			Row:              position.Row,
			Col:              position.Col,
			PositionType:     position.PositionType,
			IsUsed:           position.IsUsed,
		})
	}
	return result
}
