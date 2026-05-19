package services

import (
	"context"
	"errors"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/application/mapper"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type LayoutPositionServiceImpl struct {
	layoutPositionRepository repositories.LayoutPositionRepository
	log                      *logrus.Logger
}

func NewLayoutPositionServiceImpl(layoutPositionRepository repositories.LayoutPositionRepository, log *logrus.Logger) *LayoutPositionServiceImpl {
	return &LayoutPositionServiceImpl{
		layoutPositionRepository: layoutPositionRepository,
		log:                      log,
	}
}

func (s *LayoutPositionServiceImpl) GetLayoutPositionByID(ctx context.Context, layoutPositionID uuid.UUID) (*entities.LayoutPositions, error) {
	position, err := s.layoutPositionRepository.GetLayoutPositionByID(ctx, layoutPositionID)
	if err != nil {
		if errors.Is(err, errorConst.ErrNotFound) {
			s.log.WithFields(logrus.Fields{
				"layout_position_id": layoutPositionID,
				"error":              err,
			}).Error("layout position not found")
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return position, nil
}

func (s *LayoutPositionServiceImpl) GetLayoutPositionsByLayoutID(ctx context.Context, layoutID uuid.UUID) ([]entities.LayoutPositions, error) {
	positions, err := s.layoutPositionRepository.GetLayoutPositionsByLayoutID(ctx, layoutID)
	if err != nil {
		return nil, err
	}
	return positions, nil
}

func (s *LayoutPositionServiceImpl) GetAllLayoutPositions(ctx context.Context) ([]entities.LayoutPositions, error) {
	positions, err := s.layoutPositionRepository.GetAllLayoutPositions(ctx)
	if err != nil {
		return nil, err
	}
	return positions, nil
}

func (s *LayoutPositionServiceImpl) CreateLayoutPosition(ctx context.Context, req dto.CreateLayoutPositionDTO, layoutID uuid.UUID) error {
	position, err := mapper.CreateLayoutPositionDTOToEntity(req)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"layout_id": layoutID,
			"error":     err,
		}).Error("failed to map layout position dto to entity")
		return errorConst.ErrInternalServer
	}
	position.LayoutID = layoutID

	err = s.layoutPositionRepository.CreateLayoutPosition(ctx, position)
	if err != nil {
		return err
	}
	return nil
}

func (s *LayoutPositionServiceImpl) UpdateLayoutPosition(ctx context.Context, layoutPositionID uuid.UUID, req dto.UpdateLayoutPositionDTO) error {
	position, err := s.layoutPositionRepository.GetLayoutPositionByID(ctx, layoutPositionID)
	if err != nil {
		if errors.Is(err, errorConst.ErrNotFound) {
			s.log.WithFields(logrus.Fields{
				"layout_position_id": layoutPositionID,
				"error":              err,
			}).Error("layout position not found")
			return errorConst.ErrNotFound
		}
		return err
	}

	if req.Label != nil {
		position.Label = *req.Label
	}
	if req.Row != nil {
		position.Row = *req.Row
	}
	if req.Col != nil {
		position.Col = *req.Col
	}
	if req.PositionType != nil {
		position.PositionType = *req.PositionType
	}
	if req.IsUsed != nil {
		position.IsUsed = *req.IsUsed
	}

	err = s.layoutPositionRepository.UpdateLayoutPosition(ctx, *position)
	if err != nil {
		return err
	}
	return nil
}

func (s *LayoutPositionServiceImpl) DeleteLayoutPosition(ctx context.Context, layoutPositionID uuid.UUID) error {
	err := s.layoutPositionRepository.DeleteLayoutPosition(ctx, layoutPositionID)
	if err != nil {
		return err
	}
	return nil
}
