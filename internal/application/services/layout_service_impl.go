package services

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/application/mapper"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	"github.com/ardnh/be-travel-booking-app/internal/domain/services"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type LayoutServiceImpl struct {
	LayoutRepository repositories.LayoutRepository
	log              *logrus.Logger
}

func NewLayoutServiceImpl(layoutRepository repositories.LayoutRepository, log *logrus.Logger) services.LayoutService {
	return &LayoutServiceImpl{
		LayoutRepository: layoutRepository,
		log:              log,
	}
}

func (s *LayoutServiceImpl) GetLayoutById(ctx context.Context, layoutID string) (*dto.LayoutDTO, error) {

	layoutIdUuid, err := uuid.Parse(layoutID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"user_id": layoutID,
			"error":   err,
		}).Error("failed to parse user id")
		return nil, errorConst.ErrInternalServer
	}

	layout, err := s.LayoutRepository.GetLayoutById(ctx, layoutIdUuid)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"layout_id": layoutID,
			"error":     err,
		}).Error("failed to get layout by id")
		return nil, errorConst.ErrInternalServer
	}

	layoutDto := mapper.LayoutToDTO(layout)
	return layoutDto, nil
}

func (s *LayoutServiceImpl) GetLayout(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]*dto.LayoutDTO, int64, error) {

	result, total, err := s.LayoutRepository.GetLayout(ctx, page, pageSize, search, sortBy, sortOrder)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"page":      page,
			"pageSize":  pageSize,
			"search":    search,
			"sortBy":    sortBy,
			"sortOrder": sortOrder,
			"error":     err,
		}).Error("failed to get layout")
		return nil, 0, errorConst.ErrInternalServer
	}

	resultDto := mapper.LayoutsToDTO(result)
	return resultDto, total, nil
}

func (s *LayoutServiceImpl) CreateLayout(ctx context.Context, layout dto.CreateLayoutDTO) error {

	layoutEntity, err := mapper.CreateLayoutDTOToEntity(layout)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed to map layout dto to entity")
		return errorConst.ErrInternalServer
	}

	layoutPositionEntities := mapper.CreateLayoutPositionDTOsToEntities(layout.LayoutPositions)

	err = s.LayoutRepository.CreateLayout(ctx, layoutEntity, layoutPositionEntities)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed to create layout")
		return errorConst.ErrInternalServer
	}

	return nil
}

func (s *LayoutServiceImpl) UpdateLayout(ctx context.Context, layoutID string, layout dto.CreateLayoutDTO) error {

	// Get current layout
	layoutIdUuid, err := uuid.Parse(layoutID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed to parse layout id")
		return errorConst.ErrInternalServer
	}

	currentLayout, err := s.LayoutRepository.GetLayoutById(ctx, layoutIdUuid)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed to get current layout")
		return errorConst.ErrInternalServer
	}

	if currentLayout.GridSizeX != layout.GridSizeX {
		currentLayout.GridSizeX = layout.GridSizeX
	}

	if currentLayout.GridSizeY != layout.GridSizeY {
		currentLayout.GridSizeY = layout.GridSizeY
	}

	if layout.Name != "" {
		currentLayout.Name = layout.Name
	}

	if currentLayout.SeatCount != layout.SeatCount {
		currentLayout.SeatCount = layout.SeatCount
	}

	err = s.LayoutRepository.UpdateLayout(ctx, layoutIdUuid, currentLayout, nil)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("failed to update layout")
		return errorConst.ErrInternalServer
	}

	return nil
}

func (s *LayoutServiceImpl) DeleteLayout(ctx context.Context, layoutID string) error {

}
