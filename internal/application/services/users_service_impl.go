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

type UsersServiceImpl struct {
	userRepository repositories.UserRepository
	log            *logrus.Logger
}

func NewUsersServiceImpl(userRepository repositories.UserRepository, log *logrus.Logger) *UsersServiceImpl {
	return &UsersServiceImpl{
		userRepository: userRepository,
		log:            log,
	}
}

func (s *UsersServiceImpl) CreateUser(ctx context.Context, user dto.CreateUserDTO) error {

	userEntity := &entities.Users{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
	}

	err := s.userRepository.CreateUser(ctx, userEntity)
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

	err = s.userRepository.UpdateUser(ctx, user)
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
