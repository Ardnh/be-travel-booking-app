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
		LayoutID:     layout.LayoutID.String(),
		Name:         layout.Name,
		GridSizeX:    layout.GridSizeX,
		GridSizeY:    layout.GridSizeY,
		SeatCount:    layout.SeatCount,
		LayoutConfig: layout.LayoutConfig,
		CreatedBy:    createdBy,
	}
}

func CreateLayoutDTOToEntity(layout dto.CreateLayoutDTO) (entities.Layouts, error) {
	createdByUUID, err := uuid.Parse(layout.CreatedBy)
	if err != nil {
		return entities.Layouts{}, err
	}

	if layout.LayoutConfig == nil {
		layout.LayoutConfig = entities.LayoutConfig{} // empty slice, bukan nil
	}

	return entities.Layouts{
		Name:         layout.Name,
		GridSizeX:    layout.GridSizeX,
		GridSizeY:    layout.GridSizeY,
		SeatCount:    layout.SeatCount,
		CreatedBy:    &createdByUUID,
		LayoutConfig: layout.LayoutConfig,
	}, nil
}
