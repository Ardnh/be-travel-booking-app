package repositories

import (
	"context"
	"errors"
	"strings"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type poolPointRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPoolPointRepository(db *gorm.DB, redis *redis.Client) repositories.PoolPointRepository {
	return &poolPointRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *poolPointRepositoryImpl) GetPoolPointByID(ctx context.Context, poolID uuid.UUID) (*entities.Pools, error) {
	var poolPoint entities.Pools
	err := r.db.WithContext(ctx).Where("pool_id = ?", poolID).First(&poolPoint).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return &poolPoint, nil
}

func (r *poolPointRepositoryImpl) GetAllPoolPoints(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.Pools, int64, error) {
	var poolPoints []entities.Pools
	var total int64

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize >= 1000 {
		pageSize = 30
	}

	offset := (page - 1) * pageSize

	baseQuery := r.db.WithContext(ctx).Model(&entities.Pools{})

	if search != "" {
		baseQuery = baseQuery.Where("name ILIKE ? OR city ILIKE ? OR province ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	allowedSort := map[string]bool{
		"name":       true,
		"city":       true,
		"province":   true,
		"district":   true,
		"status":     true,
		"created_at": true,
	}

	if !allowedSort[sortBy] {
		sortBy = "created_at"
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	err := baseQuery.
		Order(sortBy + " " + sortOrder).
		Limit(pageSize).
		Offset(offset).
		Find(&poolPoints).Error

	if err != nil {
		return nil, 0, err
	}

	return poolPoints, total, nil
}

func (r *poolPointRepositoryImpl) GetPoolPointsByVendorID(ctx context.Context, vendorID uuid.UUID, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.Pools, int64, error) {
	var poolPoints []entities.Pools
	var total int64

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize >= 1000 {
		pageSize = 30
	}

	offset := (page - 1) * pageSize

	baseQuery := r.db.WithContext(ctx).
		Model(&entities.Pools{}).
		Where("vendor_id = ?", vendorID)

	if search != "" {
		baseQuery = baseQuery.Where("name ILIKE ? OR city ILIKE ? OR province ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	allowedSort := map[string]bool{
		"name":       true,
		"city":       true,
		"province":   true,
		"district":   true,
		"status":     true,
		"created_at": true,
	}

	if !allowedSort[sortBy] {
		sortBy = "created_at"
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	err := baseQuery.
		Order(sortBy + " " + sortOrder).
		Limit(pageSize).
		Offset(offset).
		Find(&poolPoints).Error

	if err != nil {
		return nil, 0, err
	}

	return poolPoints, total, nil
}

func (r *poolPointRepositoryImpl) GetAvailableLocationsByVendorID(ctx context.Context, vendorID uuid.UUID, locationType string) ([]entities.PoolLocationEntity, error) {
	locationType = strings.TrimSpace(strings.ToLower(locationType))
	if locationType == "" {
		return []entities.PoolLocationEntity{}, nil
	}

	locationColumns := map[string]string{
		"city":     "city",
		"province": "province",
		"district": "district",
		"distict":  "district",
	}

	column, ok := locationColumns[locationType]
	if !ok {
		return nil, errorConst.ErrBadRequest
	}

	var locations []entities.PoolLocationEntity
	err := r.db.WithContext(ctx).
		Model(&entities.Pools{}).
		Select(column+" as city, count(*) as total_pool").
		Where("vendor_id = ?", vendorID).
		Where(column+" <> ?", "").
		Group(column).
		Order(column + " ASC").
		Scan(&locations).Error

	if err != nil {
		return nil, err
	}

	for i := range locations {
		var pools []entities.Pools
		if err := r.db.WithContext(ctx).
			Where("vendor_id = ? AND "+column+" = ?", vendorID, locations[i].City).
			Find(&pools).Error; err != nil {
			return nil, err
		}
		locations[i].Pools = pools
	}

	return locations, nil
}

func (r *poolPointRepositoryImpl) CreatePoolPoint(ctx context.Context, poolPoint entities.Pools) (*entities.Pools, error) {
	err := r.db.WithContext(ctx).Create(&poolPoint).Error
	if err != nil {
		return nil, err
	}
	return &poolPoint, nil
}

func (r *poolPointRepositoryImpl) UpdatePoolPoint(ctx context.Context, poolPoint entities.Pools) (*entities.Pools, error) {
	if err := r.db.WithContext(ctx).Model(&entities.Pools{}).
		Where("pool_id = ?", poolPoint.PoolID).
		Updates(poolPoint).Error; err != nil {
		return nil, err
	}
	return &poolPoint, nil
}

func (r *poolPointRepositoryImpl) DeletePoolPoint(ctx context.Context, poolID uuid.UUID) error {
	err := r.db.WithContext(ctx).Where("pool_id = ?", poolID).Delete(&entities.Pools{}).Error
	if err != nil {
		return err
	}
	return nil
}
