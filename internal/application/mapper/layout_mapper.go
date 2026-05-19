package mapper

import (
	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

func LayoutsToDTO(layouts []*entities.Layouts) []*dto.LayoutDTO {
	result := make([]*dto.LayoutDTO, 0, len(layouts))
	for _, layout := range layouts {
		result = append(result, LayoutToDTO(layout))
	}
	return result
}

func LayoutToDTO(layout *entities.Layouts) *dto.LayoutDTO {
	createdBy := ""
	if layout.CreatedBy != nil {
		createdBy = layout.CreatedBy.String()
	}
	return &dto.LayoutDTO{
		LayoutID:        layout.LayoutID.String(),
		Name:            layout.Name,
		GridSizeX:       layout.GridSizeX,
		GridSizeY:       layout.GridSizeY,
		SeatCount:       layout.SeatCount,
		CreatedBy:       createdBy,
		LayoutPositions: LayoutPositionsToDTO(layout.Positions),
	}
}

func LayoutPositionsToDTO(layoutPositions []entities.LayoutPositions) []*dto.LayoutPositionDTO {
	result := make([]*dto.LayoutPositionDTO, 0, len(layoutPositions))
	for _, position := range layoutPositions {
		result = append(result, &dto.LayoutPositionDTO{
			LayoutPositionID: position.LayoutPositionID.String(),
			LayoutID:         position.LayoutID.String(),
			Label:            position.Label,
			Row:              position.Row,
			Col:              position.Col,
			PositionType:     position.PositionType,
			IsUsed:           position.IsUsed,
		})
	}
	return result
}

func CreateLayoutDTOToEntity(layout dto.CreateLayoutDTO) (entities.Layouts, error) {
	createdByUUID, err := uuid.Parse(layout.CreatedBy)
	if err != nil {
		return entities.Layouts{}, err
	}

	return entities.Layouts{
		Name:      layout.Name,
		GridSizeX: layout.GridSizeX,
		GridSizeY: layout.GridSizeY,
		SeatCount: layout.SeatCount,
		CreatedBy: &createdByUUID,
	}, nil
}

func CreateLayoutPositionDTOsToEntities(layoutPositions []dto.CreateLayoutPositionDTO) []entities.LayoutPositions {
	result := make([]entities.LayoutPositions, 0, len(layoutPositions))
	for _, position := range layoutPositions {
		result = append(result, entities.LayoutPositions{
			Row:          position.Row,
			Col:          position.Col,
			PositionType: position.PositionType,
			IsUsed:       position.IsUsed,
		})
	}
	return result
}
