package services

import (
	"context"
	"errors"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/config"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	"github.com/ardnh/be-travel-booking-app/internal/domain/services"
	casbin_utils "github.com/ardnh/be-travel-booking-app/internal/utils/casbin"
	jwt_utils "github.com/ardnh/be-travel-booking-app/internal/utils/jwt"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/casbin/casbin/v3"
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

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
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

	permissions := casbin_utils.GetUserPermissions(s.casbinEnforcer, user.UserID.String())

	s.log.WithField("user_id", user.UserID).Info("login successful")

	return &dto.LoginResponseDto{
		Token:       *token,
		ExpireDate:  *expiredTimeISO,
		Permissions: permissions,
	}, nil
}
