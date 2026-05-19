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

func (s *ServiceTypeServiceImpl) GetAllServiceTypes(ctx context.Context) ([]entities.ServiceTypes, error) {
	serviceTypes, err := s.serviceTypeRepository.GetAllServiceTypes(ctx)
	if err != nil {
		return nil, err
	}
	return serviceTypes, nil
}

func (s *ServiceTypeServiceImpl) CreateServiceType(ctx context.Context, req dto.CreateServiceTypeDTO, createdBy uuid.UUID) error {
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

	err := s.serviceTypeRepository.CreateServiceType(ctx, *serviceType)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceTypeServiceImpl) UpdateServiceType(ctx context.Context, serviceTypeID uuid.UUID, req dto.UpdateServiceTypeDTO) error {
	serviceType, err := s.serviceTypeRepository.GetServiceTypeByID(ctx, serviceTypeID)
	if err != nil {
		if errors.Is(err, errorConst.ErrNotFound) {
			s.log.WithFields(logrus.Fields{
				"service_type_id": serviceTypeID,
				"error":           err,
			}).Error("service type not found")
			return errorConst.ErrNotFound
		}
		return err
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

	err = s.serviceTypeRepository.UpdateServiceType(ctx, *serviceType)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceTypeServiceImpl) DeleteServiceType(ctx context.Context, serviceTypeID uuid.UUID) error {
	err := s.serviceTypeRepository.DeleteServiceType(ctx, serviceTypeID)
	if err != nil {
		return err
	}
	return nil
}