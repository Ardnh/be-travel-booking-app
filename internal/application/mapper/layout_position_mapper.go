package mapper

import (
	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
)

func LayoutPositionToDTO(position entities.LayoutPositions) *dto.LayoutPositionDTO {
	return &dto.LayoutPositionDTO{
		LayoutPositionID: position.LayoutPositionID.String(),
		LayoutID:         position.LayoutID.String(),
		Label:            position.Label,
		Row:              position.Row,
		Col:              position.Col,
		PositionType:     position.PositionType,
		IsUsed:           position.IsUsed,
	}
}

func LayoutPositionsToDTO(positions []entities.LayoutPositions) []*dto.LayoutPositionDTO {
	result := make([]*dto.LayoutPositionDTO, 0, len(positions))
	for _, p := range positions {
		result = append(result, LayoutPositionToDTO(p))
	}
	return result
}

func CreateLayoutPositionDTOToEntity(req dto.CreateLayoutPositionDTO) (entities.LayoutPositions, error) {
	layoutID, err := uuid.Parse(req.LayoutID)
	if err != nil {
		return entities.LayoutPositions{}, err
	}
	return entities.LayoutPositions{
		LayoutID:     layoutID,
		Label:        req.Label,
		Row:          req.Row,
		Col:          req.Col,
		PositionType: req.PositionType,
		IsUsed:       req.IsUsed,
	}, nil
}
