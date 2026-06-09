package services

import (
	"context"
	"errors"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type PoolPointServiceImpl struct {
	poolPointRepository repositories.PoolPointRepository
	log                 *logrus.Logger
}

func NewPoolPointServiceImpl(poolPointRepository repositories.PoolPointRepository, log *logrus.Logger) *PoolPointServiceImpl {
	return &PoolPointServiceImpl{
		poolPointRepository: poolPointRepository,
		log:                 log,
	}
}

func (s *PoolPointServiceImpl) GetPoolPointByID(ctx context.Context, poolID uuid.UUID) (*entities.Pools, error) {
	poolPoint, err := s.poolPointRepository.GetPoolPointByID(ctx, poolID)
	if err != nil {
		if errors.Is(err, errorConst.ErrNotFound) {
			s.log.WithFields(logrus.Fields{
				"pool_id": poolID,
				"error":   err,
			}).Error("pool point not found")
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return poolPoint, nil
}

func (s *PoolPointServiceImpl) GetAllPoolPoints(ctx context.Context) ([]entities.Pools, error) {
	poolPoints, err := s.poolPointRepository.GetAllPoolPoints(ctx)
	if err != nil {
		return nil, err
	}
	return poolPoints, nil
}

func (s *PoolPointServiceImpl) GetPoolPointsByVendorID(ctx context.Context, vendorID uuid.UUID) ([]entities.Pools, error) {
	poolPoints, err := s.poolPointRepository.GetPoolPointsByVendorID(ctx, vendorID)
	if err != nil {
		return nil, err
	}
	return poolPoints, nil
}

func (s *PoolPointServiceImpl) CreatePoolPoint(ctx context.Context, req dto.CreatePoolsDTO) (*entities.Pools, error) {
	vendorID, err := uuid.Parse(req.VendorID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"vendor_id": req.VendorID,
			"error":     err,
		}).Error("failed to parse vendor id")
		return nil, errorConst.ErrBadRequest
	}

	poolPoint := &entities.Pools{
		PoolID:      uuid.New(),
		VendorID:    vendorID,
		Name:        req.Name,
		Slug:        req.Slug,
		Address:     req.Address,
		City:        req.City,
		Province:    req.Province,
		District:    req.District,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		OpenTime:    req.OpenTime,
		CloseTime:   req.CloseTime,
		Status:      req.Status,
		Description: req.Description,
		EmbedURL:    req.EmbedURL,
	}

	createdPoolPoint, err := s.poolPointRepository.CreatePoolPoint(ctx, *poolPoint)
	if err != nil {
		return nil, err
	}

	return createdPoolPoint, nil
}

func (s *PoolPointServiceImpl) UpdatePoolPoint(ctx context.Context, poolID uuid.UUID, req dto.UpdatePoolsDTO) (*entities.Pools, error) {
	poolPoint, err := s.poolPointRepository.GetPoolPointByID(ctx, poolID)
	if err != nil {
		if errors.Is(err, errorConst.ErrNotFound) {
			s.log.WithFields(logrus.Fields{
				"pool_id": poolID,
				"error":   err,
			}).Error("pool point not found")
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}

	if req.VendorID != nil {
		vendorID, err := uuid.Parse(*req.VendorID)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"vendor_id": *req.VendorID,
				"error":     err,
			}).Error("failed to parse vendor id")
			return nil, errorConst.ErrBadRequest
		}
		poolPoint.VendorID = vendorID
	}
	if req.Name != nil {
		poolPoint.Name = *req.Name
	}
	if req.Slug != nil {
		poolPoint.Slug = *req.Slug
	}
	if req.Address != nil {
		poolPoint.Address = *req.Address
	}
	if req.City != nil {
		poolPoint.City = *req.City
	}
	if req.Province != nil {
		poolPoint.Province = *req.Province
	}
	if req.District != nil {
		poolPoint.District = *req.District
	}
	if req.Latitude != nil {
		poolPoint.Latitude = *req.Latitude
	}
	if req.Longitude != nil {
		poolPoint.Longitude = *req.Longitude
	}
	if req.OpenTime != nil {
		poolPoint.OpenTime = *req.OpenTime
	}
	if req.CloseTime != nil {
		poolPoint.CloseTime = *req.CloseTime
	}
	if req.Status != nil {
		poolPoint.Status = *req.Status
	}
	if req.Description != nil {
		poolPoint.Description = *req.Description
	}
	if req.EmbedURL != nil {
		poolPoint.EmbedURL = *req.EmbedURL
	}

	updatedPoolPoint, err := s.poolPointRepository.UpdatePoolPoint(ctx, *poolPoint)
	if err != nil {
		return nil, err
	}

	return updatedPoolPoint, nil
}

func (s *PoolPointServiceImpl) DeletePoolPoint(ctx context.Context, poolID uuid.UUID) error {
	err := s.poolPointRepository.DeletePoolPoint(ctx, poolID)
	if err != nil {
		return err
	}
	return nil
}
