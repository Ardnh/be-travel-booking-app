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

type layoutPositionRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewLayoutPositionRepository(db *gorm.DB, redis *redis.Client) repositories.LayoutPositionRepository {
	return &layoutPositionRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *layoutPositionRepositoryImpl) GetLayoutPositionByID(ctx context.Context, layoutPositionID uuid.UUID) (*entities.LayoutPositions, error) {
	var position entities.LayoutPositions
	err := r.db.WithContext(ctx).Where("layout_position_id = ?", layoutPositionID).First(&position).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return &position, nil
}

func (r *layoutPositionRepositoryImpl) GetLayoutPositionsByLayoutID(ctx context.Context, layoutID uuid.UUID) ([]entities.LayoutPositions, error) {
	var positions []entities.LayoutPositions
	err := r.db.WithContext(ctx).Where("layout_id = ?", layoutID).Find(&positions).Error
	if err != nil {
		return nil, err
	}
	return positions, nil
}

func (r *layoutPositionRepositoryImpl) GetAllLayoutPositions(ctx context.Context) ([]entities.LayoutPositions, error) {
	var positions []entities.LayoutPositions
	err := r.db.WithContext(ctx).Find(&positions).Error
	if err != nil {
		return nil, err
	}
	return positions, nil
}

func (r *layoutPositionRepositoryImpl) CreateLayoutPosition(ctx context.Context, layoutPosition entities.LayoutPositions) error {
	err := r.db.WithContext(ctx).Create(&layoutPosition).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *layoutPositionRepositoryImpl) UpdateLayoutPosition(ctx context.Context, layoutPosition entities.LayoutPositions) error {
	err := r.db.WithContext(ctx).Save(&layoutPosition).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *layoutPositionRepositoryImpl) DeleteLayoutPosition(ctx context.Context, layoutPositionID uuid.UUID) error {
	err := r.db.WithContext(ctx).Where("layout_position_id = ?", layoutPositionID).Delete(&entities.LayoutPositions{}).Error
	if err != nil {
		return err
	}
	return nil
}
