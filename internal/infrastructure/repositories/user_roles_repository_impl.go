package repositories

import (
	"context"
	"errors"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type userRolesRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUserRolesRepository(db *gorm.DB, redis *redis.Client) repositories.UserRolesRepository {
	return &userRolesRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *userRolesRepositoryImpl) GetUserRoleByID(ctx context.Context, userRoleID uuid.UUID) (*entities.UserRoles, error) {
	var userRole entities.UserRoles
	err := r.db.WithContext(ctx).Where("user_role_id = ?", userRoleID).First(&userRole).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return &userRole, nil
}

func (r *userRolesRepositoryImpl) GetUserRolesByUserID(ctx context.Context, userID uuid.UUID) ([]entities.UserRoles, error) {
	var userRoles []entities.UserRoles
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&userRoles).Error
	if err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (r *userRolesRepositoryImpl) GetUserRolesByVendorID(ctx context.Context, vendorID uuid.UUID) ([]entities.UserRoles, error) {
	var userRoles []entities.UserRoles
	err := r.db.WithContext(ctx).Where("vendor_id = ?", vendorID).Find(&userRoles).Error
	if err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (r *userRolesRepositoryImpl) CreateUserRole(ctx context.Context, userRole entities.UserRoles) error {
	err := r.db.WithContext(ctx).Create(&userRole).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRolesRepositoryImpl) UpdateUserRole(ctx context.Context, userRole entities.UserRoles) error {
	err := r.db.WithContext(ctx).Save(&userRole).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRolesRepositoryImpl) DeleteUserRole(ctx context.Context, userRoleID uuid.UUID) error {
	err := r.db.WithContext(ctx).Where("user_role_id = ?", userRoleID).Delete(&entities.UserRoles{}).Error
	if err != nil {
		return err
	}
	return nil
}
