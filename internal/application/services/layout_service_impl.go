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

}

func (s *LayoutServiceImpl) UpdateLayout(ctx context.Context, layoutID string, layout dto.CreateLayoutDTO) error {

}

func (s *LayoutServiceImpl) DeleteLayout(ctx context.Context, layoutID string) error {

}
