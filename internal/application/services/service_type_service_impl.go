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

type ServiceTypeServiceImpl struct {
	serviceTypeRepository repositories.ServiceTypeRepository
	log                    *logrus.Logger
}

func NewServiceTypeServiceImpl(serviceTypeRepository repositories.ServiceTypeRepository, log *logrus.Logger) *ServiceTypeServiceImpl {
	return &ServiceTypeServiceImpl{
		serviceTypeRepository: serviceTypeRepository,
		log:                   log,
	}
}

func (s *ServiceTypeServiceImpl) GetServiceTypeByID(ctx context.Context, serviceTypeID uuid.UUID) (*entities.ServiceTypes, error) {
	serviceType, err := s.serviceTypeRepository.GetServiceTypeByID(ctx, serviceTypeID)
	if err != nil {
		if errors.Is(err, errorConst.ErrNotFound) {
			s.log.WithFields(logrus.Fields{
				"service_type_id": serviceTypeID,
				"error":           err,
			}).Error("service type not found")
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return serviceType, nil
}

func (s *ServiceTypeServiceImpl) GetAllServiceTypes(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.ServiceTypes, int64, error) {
	serviceTypes, total, err := s.serviceTypeRepository.GetAllServiceTypes(ctx, page, pageSize, search, sortBy, sortOrder)
	if err != nil {
		return nil, 0, err
	}
	return serviceTypes, total, nil
}

func (s *ServiceTypeServiceImpl) CreateServiceType(ctx context.Context, req dto.CreateServiceTypeDTO, createdBy uuid.UUID) (*entities.ServiceTypes, error) {
	serviceType := &entities.ServiceTypes{
		ServiceTypeID:      uuid.New(),
		Name:               req.Name,
		UniqueCode:         req.UniqueCode,
		Description:        req.Description,
		NeedChair:          req.NeedChair,
		NeedPickupAddress:  req.NeedPickupAddress,
		NeedDropoffAddress: req.NeedDropoffAddress,
		DisplayOrder:       req.DisplayOrder,
		Status:             req.Status,
		CreatedBy:          createdBy,
	}

	createdServiceType, err := s.serviceTypeRepository.CreateServiceType(ctx, *serviceType)
	if err != nil {
		return nil, err
	}

	return createdServiceType, nil
}

func (s *ServiceTypeServiceImpl) UpdateServiceType(ctx context.Context, serviceTypeID uuid.UUID, req dto.UpdateServiceTypeDTO) (*entities.ServiceTypes, error) {
	serviceType, err := s.serviceTypeRepository.GetServiceTypeByID(ctx, serviceTypeID)
	if err != nil {
		if errors.Is(err, errorConst.ErrNotFound) {
			s.log.WithFields(logrus.Fields{
				"service_type_id": serviceTypeID,
				"error":           err,
			}).Error("service type not found")
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}

	if req.Name != nil {
		serviceType.Name = *req.Name
	}
	if req.UniqueCode != nil {
		serviceType.UniqueCode = *req.UniqueCode
	}
	if req.Description != nil {
		serviceType.Description = *req.Description
	}
	if req.NeedChair != nil {
		serviceType.NeedChair = *req.NeedChair
	}
	if req.NeedPickupAddress != nil {
		serviceType.NeedPickupAddress = *req.NeedPickupAddress
	}
	if req.NeedDropoffAddress != nil {
		serviceType.NeedDropoffAddress = *req.NeedDropoffAddress
	}
	if req.DisplayOrder != nil {
		serviceType.DisplayOrder = *req.DisplayOrder
	}
	if req.Status != nil {
		serviceType.Status = *req.Status
	}

	updatedServiceType, err := s.serviceTypeRepository.UpdateServiceType(ctx, *serviceType)
	if err != nil {
		return nil, err
	}

	return updatedServiceType, nil
}

func (s *ServiceTypeServiceImpl) DeleteServiceType(ctx context.Context, serviceTypeID uuid.UUID) error {
	err := s.serviceTypeRepository.DeleteServiceType(ctx, serviceTypeID)
	if err != nil {
		return err
	}
	return nil
}