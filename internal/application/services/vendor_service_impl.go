package services

import (
	"context"
	"errors"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	"github.com/ardnh/be-travel-booking-app/pkg/constants"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/casbin/casbin/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type VendorServiceImpl struct {
	vendorRepository repositories.VendorRepository
	casbinEnforcer   *casbin.Enforcer
	log              *logrus.Logger
}

func NewVendorServiceImpl(vendorRepository repositories.VendorRepository, casbinEnforcer *casbin.Enforcer, log *logrus.Logger) *VendorServiceImpl {
	return &VendorServiceImpl{
		vendorRepository: vendorRepository,
		casbinEnforcer:   casbinEnforcer,
		log:              log,
	}
}

func (s *VendorServiceImpl) GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*entities.Vendors, error) {
	vendor, err := s.vendorRepository.GetVendorByID(ctx, vendorID)
	if err != nil {
		if errors.Is(err, errorConst.ErrNotFound) {
			s.log.WithFields(logrus.Fields{
				"vendor_id": vendorID,
				"error":     err,
			}).Error("vendor not found")
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return vendor, nil
}

func (s *VendorServiceImpl) GetVendorByOwnerUserID(ctx context.Context, ownerUserID uuid.UUID) (*entities.Vendors, error) {
	vendor, err := s.vendorRepository.GetVendorByOwnerUserID(ctx, ownerUserID)
	if err != nil {
		if errors.Is(err, errorConst.ErrNotFound) {
			s.log.WithFields(logrus.Fields{
				"owner_user_id": ownerUserID,
				"error":         err,
			}).Error("vendor not found for owner user")
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return vendor, nil
}

func (s *VendorServiceImpl) GetAllVendors(ctx context.Context) ([]entities.Vendors, error) {
	vendors, err := s.vendorRepository.GetAllVendors(ctx)
	if err != nil {
		return nil, err
	}
	return vendors, nil
}

func (s *VendorServiceImpl) CreateVendor(ctx context.Context, req dto.CreateVendorDTO) error {

	ownerUserId, err := uuid.Parse(req.OwnerUserID)
	if err != nil {
		return err
	}

	vendor := entities.Vendors{
		VendorID:          uuid.New(),
		OwnerUserID:       ownerUserId,
		BusinessName:      req.BusinessName,
		OwnerName:         req.OwnerName,
		Description:       req.Description,
		FoundedYear:       req.FoundedYear,
		PhoneNumber:       req.PhoneNumber,
		Email:             req.Email,
		HeadOfficeAddress: req.HeadOfficeAddress,
		Status:            req.Status,
	}

	if req.LegalDocumentNumber != nil {
		vendor.LegalDocumentNumber = *req.LegalDocumentNumber
	}

	errCreate := s.vendorRepository.CreateVendor(ctx, vendor)
	if errCreate != nil {
		return errCreate
	}

	_, err = s.casbinEnforcer.AddGroupingPolicy(ownerUserId.String(), constants.RoleBusinessOwner)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"email":  req.Email,
			"userID": ownerUserId,
			"error":  err,
		}).Error("failed to add casbin grouping policy")
		return errorConst.ErrInternalServer
	}

	s.casbinEnforcer.LoadPolicy()

	return nil
}

func (s *VendorServiceImpl) UpdateVendor(ctx context.Context, vendorID uuid.UUID, req dto.UpdateVendorDTO) error {
	vendor, err := s.vendorRepository.GetVendorByID(ctx, vendorID)
	if err != nil {
		if errors.Is(err, errorConst.ErrNotFound) {
			s.log.WithFields(logrus.Fields{
				"vendor_id": vendorID,
				"error":     err,
			}).Error("vendor not found")
			return errorConst.ErrNotFound
		}
		return err
	}

	if req.BusinessName != nil {
		vendor.BusinessName = *req.BusinessName
	}
	if req.OwnerName != nil {
		vendor.OwnerName = *req.OwnerName
	}
	if req.Description != nil {
		vendor.Description = *req.Description
	}
	if req.FoundedYear != nil {
		vendor.FoundedYear = *req.FoundedYear
	}
	if req.PhoneNumber != nil {
		vendor.PhoneNumber = *req.PhoneNumber
	}
	if req.Email != nil {
		vendor.Email = *req.Email
	}
	if req.HeadOfficeAddress != nil {
		vendor.HeadOfficeAddress = *req.HeadOfficeAddress
	}
	if req.LogoURL != nil {
		vendor.LogoURL = *req.LogoURL
	}
	if req.BannerURL != nil {
		vendor.BannerURL = *req.BannerURL
	}
	if req.LegalDocumentNumber != nil {
		vendor.LegalDocumentNumber = *req.LegalDocumentNumber
	}
	if req.IsVerified != nil {
		vendor.IsVerified = *req.IsVerified
	}
	if req.Status != nil {
		vendor.Status = *req.Status
	}

	err = s.vendorRepository.UpdateVendor(ctx, *vendor)
	if err != nil {
		return err
	}

	return nil
}

func (s *VendorServiceImpl) DeleteVendor(ctx context.Context, vendorID uuid.UUID) error {
	err := s.vendorRepository.DeleteVendor(ctx, vendorID)
	if err != nil {
		return err
	}
	return nil
}
