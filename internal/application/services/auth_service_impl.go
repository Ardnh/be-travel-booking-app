package services

import (
	"context"
	"errors"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/config"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	"github.com/ardnh/be-travel-booking-app/internal/domain/services"
	jwt_utils "github.com/ardnh/be-travel-booking-app/internal/utils/jwt"
	"github.com/ardnh/be-travel-booking-app/pkg/constants"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/casbin/casbin/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	userRepository repositories.UserRepository
	log            *logrus.Logger
	appConfig      *config.Config
	casbinEnforcer *casbin.Enforcer
}

func NewAuthService(
	userRepository repositories.UserRepository,
	log *logrus.Logger,
	appConfig *config.Config,
	casbinEnforcer *casbin.Enforcer,
) services.AuthService {
	return &AuthServiceImpl{
		userRepository: userRepository,
		log:            log,
		appConfig:      appConfig,
		casbinEnforcer: casbinEnforcer,
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, req dto.LoginRequestDto) (*dto.LoginResponseDto, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		// ErrNotFound dari repo sudah di-wrap jadi ErrUnauthorized
		// agar tidak bisa di-enumerate (email terdaftar atau tidak)
		if errors.Is(err, errorConst.ErrNotFound) {
			s.log.WithField("email", req.Email).Warn("login attempt with unregistered email")
			return nil, errorConst.ErrUnauthorized
		}
		return nil, errorConst.ErrInternalServer
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		s.log.WithField("email", req.Email).Warn("invalid credentials provided")
		return nil, errorConst.ErrUnauthorized
	}

	secretKey := []byte(s.appConfig.App.JWTSecret)
	token, expiredTimeISO, err := jwt_utils.GenerateToken(secretKey, user.UserID.String())
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"user_id": user.UserID,
			"error":   err,
		}).Error("failed to generate token")
		return nil, errorConst.ErrInternalServer
	}

	s.log.WithField("user_id", user.UserID).Info("login successful")

	return &dto.LoginResponseDto{
		Token:      *token,
		ExpireDate: *expiredTimeISO,
	}, nil
}

func (s *AuthServiceImpl) Register(ctx context.Context, req dto.RegisterRequestDto) (*dto.RegisterResponseDto, error) {
	existingUser, err := s.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, errorConst.ErrNotFound) {
		s.log.WithFields(logrus.Fields{
			"email": req.Email,
			"error": err,
		}).Error("failed to check existing user")
		return nil, errorConst.ErrNotFound
	}

	if existingUser != nil {
		s.log.WithField("email", req.Email).Warn("registration attempt with existing email")
		return nil, errorConst.ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithField("email", req.Email).Error("failed to hash password")
		return nil, errorConst.ErrInternalServer
	}

	user := entities.Users{
		UserID:       uuid.New(),
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Phone:        req.Phone,
	}

	err = s.userRepository.CreateUser(ctx, user)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"email": req.Email,
			"error": err,
		}).Error("failed to create user")
		return nil, errorConst.ErrInternalServer
	}

	// Add role grouping to Casbin
	_, err = s.casbinEnforcer.AddGroupingPolicy(user.UserID.String(), constants.RoleDailyUser)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"email":  req.Email,
			"userID": user.UserID,
			"error":  err,
		}).Error("failed to add casbin grouping policy")
		return nil, errorConst.ErrInternalServer
	}

	return nil, nil
}
