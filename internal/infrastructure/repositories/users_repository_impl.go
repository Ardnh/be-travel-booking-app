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

type userRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewUsersRepository(db *gorm.DB, redis *redis.Client) repositories.UserRepository {
	return &userRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *userRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*entities.Users, error) {
	var user entities.Users
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetUserByID(ctx context.Context, userID uuid.UUID) (*entities.Users, error) {
	var user entities.Users
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) CreateUser(ctx context.Context, user *entities.Users) error {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryImpl) UpdateUser(ctx context.Context, user *entities.Users) error {
	err := r.db.WithContext(ctx).Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryImpl) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&entities.Users{}).Error
	if err != nil {
		return err
	}
	return nil
}
