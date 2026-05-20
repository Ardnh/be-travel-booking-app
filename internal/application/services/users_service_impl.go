package services

import (
	"context"
	"errors"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	casbin_utils "github.com/ardnh/be-travel-booking-app/internal/utils/casbin"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/casbin/casbin/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UsersServiceImpl struct {
	userRepository      repositories.UserRepository
	userRolesRepository repositories.UserRolesRepository
	enforcer            *casbin.Enforcer
	log                 *logrus.Logger
}

func NewUsersServiceImpl(userRepository repositories.UserRepository, userRolesRepository repositories.UserRolesRepository, enforcer *casbin.Enforcer, log *logrus.Logger) *UsersServiceImpl {
	return &UsersServiceImpl{
		userRepository:      userRepository,
		userRolesRepository: userRolesRepository,
		enforcer:            enforcer,
		log:                 log,
	}
}

func (s *UsersServiceImpl) CreateUser(ctx context.Context, user dto.CreateUserDTO) error {

	userEntity := &entities.Users{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.Password,
		Phone:        user.Phone,
	}

	err := s.userRepository.CreateUser(ctx, *userEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *UsersServiceImpl) UpdateUser(ctx context.Context, userId string, req dto.UpdateUserDTO) error {

	userIdUuid, err := uuid.Parse(userId)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"user_id": userId,
			"error":   err,
		}).Error("failed to parse user id")
		return errorConst.ErrInternalServer
	}

	user, err := s.userRepository.GetUserByID(ctx, userIdUuid)
	if err != nil {
		if errors.Is(err, errorConst.ErrNotFound) {
			s.log.WithFields(logrus.Fields{
				"user_id": userId,
				"error":   err,
			}).Error("failed to get user by id")
			return errorConst.ErrUnauthorized
		}
		return errorConst.ErrInternalServer
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	err = s.userRepository.UpdateUser(ctx, *user)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"user_id": userId,
			"error":   err,
		}).Error("failed to get user by id")
		return errorConst.ErrInternalServer
	}

	return nil
}

func (s *UsersServiceImpl) DeleteUser(ctx context.Context, userID string) error {

	userIdUuid, err := uuid.Parse(userID)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"user_id": userID,
			"error":   err,
		}).Error("failed to parse user id")
		return errorConst.ErrInternalServer
	}

	err = s.userRepository.DeleteUser(ctx, userIdUuid)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"user_id": userID,
			"error":   err,
		}).Error("failed to delete user")
		return errorConst.ErrInternalServer
	}

	return nil
}

func (s *UsersServiceImpl) GetUserProfile(ctx context.Context, userID string) (*dto.UserProfileDTO, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		s.log.WithField("user_id", userID).Error("invalid user ID format")
		return nil, errorConst.ErrBadRequest
	}

	user, err := s.userRepository.GetUserByID(ctx, userUUID)
	if err != nil {
		if errors.Is(err, errorConst.ErrNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}

	userRoles, err := s.userRolesRepository.GetUserRolesByUserID(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	var roles []dto.UserRoleDTO
	for _, ur := range userRoles {
		roles = append(roles, dto.UserRoleDTO{
			UserRoleID: ur.UserRoleID.String(),
			Role:       ur.Role,
		})
	}

	permissions := casbin_utils.GetUserPermissions(s.enforcer, user.UserID.String())

	return &dto.UserProfileDTO{
		UserID:      user.UserID.String(),
		Name:        user.Name,
		Email:       user.Email,
		Phone:       user.Phone,
		AvatarURL:   user.AvatarURL,
		IsActive:    user.IsActive,
		Roles:       roles,
		Permissions: permissions,
		CreatedAt:   user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (s *UsersServiceImpl) uuidToString(id *uuid.UUID) *string {
	if id == nil {
		return nil
	}
	str := id.String()
	return &str
}
