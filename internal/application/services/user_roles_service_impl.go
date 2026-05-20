package services

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserRolesServiceImpl struct {
	userRolesRepository repositories.UserRolesRepository
	log                 *logrus.Logger
}

func NewUserRolesServiceImpl(userRolesRepository repositories.UserRolesRepository, log *logrus.Logger) *UserRolesServiceImpl {
	return &UserRolesServiceImpl{
		userRolesRepository: userRolesRepository,
		log:                 log,
	}
}

func (s *UserRolesServiceImpl) GetUserRoleByID(ctx context.Context, userRoleID string) (*dto.UserRoleDTO, error) {
	userRoleUUID, err := uuid.Parse(userRoleID)
	if err != nil {
		s.log.WithField("user_role_id", userRoleID).Error("invalid user role ID")
		return nil, errorConst.ErrBadRequest
	}

	userRole, err := s.userRolesRepository.GetUserRoleByID(ctx, userRoleUUID)
	if err != nil {
		return nil, err
	}

	return &dto.UserRoleDTO{
		UserRoleID: userRole.UserRoleID.String(),
		UserID:     userRole.UserID.String(),
		Role:       userRole.Role,
		VendorID:   s.uuidToString(userRole.VendorID),
		PoolID:     s.uuidToString(userRole.PoolID),
	}, nil
}

func (s *UserRolesServiceImpl) GetUserRolesByUserID(ctx context.Context, userID string) ([]dto.UserRoleDTO, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		s.log.WithField("user_id", userID).Error("invalid user ID")
		return nil, errorConst.ErrBadRequest
	}

	userRoles, err := s.userRolesRepository.GetUserRolesByUserID(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	var result []dto.UserRoleDTO
	for _, ur := range userRoles {
		result = append(result, dto.UserRoleDTO{
			UserRoleID: ur.UserRoleID.String(),
			UserID:     ur.UserID.String(),
			Role:       ur.Role,
			VendorID:   s.uuidToString(ur.VendorID),
			PoolID:     s.uuidToString(ur.PoolID),
		})
	}
	return result, nil
}

func (s *UserRolesServiceImpl) GetUserRolesByVendorID(ctx context.Context, vendorID string) ([]dto.UserRoleDTO, error) {
	vendorUUID, err := uuid.Parse(vendorID)
	if err != nil {
		s.log.WithField("vendor_id", vendorID).Error("invalid vendor ID")
		return nil, errorConst.ErrBadRequest
	}

	userRoles, err := s.userRolesRepository.GetUserRolesByVendorID(ctx, vendorUUID)
	if err != nil {
		return nil, err
	}

	var result []dto.UserRoleDTO
	for _, ur := range userRoles {
		result = append(result, dto.UserRoleDTO{
			UserRoleID: ur.UserRoleID.String(),
			UserID:     ur.UserID.String(),
			Role:       ur.Role,
			VendorID:   s.uuidToString(ur.VendorID),
			PoolID:     s.uuidToString(ur.PoolID),
		})
	}
	return result, nil
}

func (s *UserRolesServiceImpl) CreateUserRole(ctx context.Context, req dto.CreateUserRoleDTO) (*dto.UserRoleDTO, error) {
	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		s.log.WithField("user_id", req.UserID).Error("invalid user ID")
		return nil, errorConst.ErrBadRequest
	}

	userRole := entities.UserRoles{
		UserRoleID: uuid.New(),
		UserID:     userUUID,
		Role:       req.Role,
	}

	if req.VendorID != nil {
		vendorUUID, err := uuid.Parse(*req.VendorID)
		if err != nil {
			s.log.WithField("vendor_id", *req.VendorID).Error("invalid vendor ID")
			return nil, errorConst.ErrBadRequest
		}
		userRole.VendorID = &vendorUUID
	}

	if req.PoolID != nil {
		poolUUID, err := uuid.Parse(*req.PoolID)
		if err != nil {
			s.log.WithField("pool_id", *req.PoolID).Error("invalid pool ID")
			return nil, errorConst.ErrBadRequest
		}
		userRole.PoolID = &poolUUID
	}

	err = s.userRolesRepository.CreateUserRole(ctx, userRole)
	if err != nil {
		return nil, err
	}

	return &dto.UserRoleDTO{
		UserRoleID: userRole.UserRoleID.String(),
		UserID:     userRole.UserID.String(),
		Role:       userRole.Role,
		VendorID:   s.uuidToString(userRole.VendorID),
		PoolID:     s.uuidToString(userRole.PoolID),
	}, nil
}

func (s *UserRolesServiceImpl) UpdateUserRole(ctx context.Context, userRoleID string, req dto.UpdateUserRoleDTO) (*dto.UserRoleDTO, error) {
	userRoleUUID, err := uuid.Parse(userRoleID)
	if err != nil {
		s.log.WithField("user_role_id", userRoleID).Error("invalid user role ID")
		return nil, errorConst.ErrBadRequest
	}

	existingRole, err := s.userRolesRepository.GetUserRoleByID(ctx, userRoleUUID)
	if err != nil {
		return nil, err
	}

	if req.Role != nil {
		existingRole.Role = *req.Role
	}

	if req.VendorID != nil {
		vendorUUID, err := uuid.Parse(*req.VendorID)
		if err != nil {
			s.log.WithField("vendor_id", *req.VendorID).Error("invalid vendor ID")
			return nil, errorConst.ErrBadRequest
		}
		existingRole.VendorID = &vendorUUID
	}

	if req.PoolID != nil {
		poolUUID, err := uuid.Parse(*req.PoolID)
		if err != nil {
			s.log.WithField("pool_id", *req.PoolID).Error("invalid pool ID")
			return nil, errorConst.ErrBadRequest
		}
		existingRole.PoolID = &poolUUID
	}

	err = s.userRolesRepository.UpdateUserRole(ctx, *existingRole)
	if err != nil {
		return nil, err
	}

	return &dto.UserRoleDTO{
		UserRoleID: existingRole.UserRoleID.String(),
		UserID:     existingRole.UserID.String(),
		Role:       existingRole.Role,
		VendorID:   s.uuidToString(existingRole.VendorID),
		PoolID:     s.uuidToString(existingRole.PoolID),
	}, nil
}

func (s *UserRolesServiceImpl) DeleteUserRole(ctx context.Context, userRoleID string) error {
	userRoleUUID, err := uuid.Parse(userRoleID)
	if err != nil {
		s.log.WithField("user_role_id", userRoleID).Error("invalid user role ID")
		return errorConst.ErrBadRequest
	}

	err = s.userRolesRepository.DeleteUserRole(ctx, userRoleUUID)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserRolesServiceImpl) uuidToString(id *uuid.UUID) *string {
	if id == nil {
		return nil
	}
	str := id.String()
	return &str
}