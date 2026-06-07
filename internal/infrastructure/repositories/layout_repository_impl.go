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

type LayoutRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewLayoutRepository(db *gorm.DB, redis *redis.Client) repositories.LayoutRepository {
	return &LayoutRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *LayoutRepositoryImpl) GetLayoutById(ctx context.Context, layoutID uuid.UUID) (*entities.Layouts, error) {
	var layout entities.Layouts

	err := r.db.WithContext(ctx).First(&layout, "layout_id = ?", layoutID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}

	return &layout, nil
}

func (r *LayoutRepositoryImpl) GetLayout(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]*entities.Layouts, int64, error) {

	var (
		layouts []*entities.Layouts
		total   int64
	)

	// --- Default pagination guard
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize >= 1000 {
		pageSize = 30
	}

	offset := (page - 1) * pageSize

	// --- Base query TANPA preload (untuk count & filter)
	baseQuery := r.db.WithContext(ctx).
		Model(&entities.Layouts{})

	// --- Search (by name)
	if search != "" {
		baseQuery = baseQuery.Where("name ILIKE ?", "%"+search+"%")
	}

	// --- Hitung total data dulu (WAJIB sebelum limit)
	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// --- Validasi sortBy agar tidak SQL injection
	allowedSort := map[string]bool{
		"name":       true,
		"created_at": true,
		"seat_count": true,
	}

	if !allowedSort[sortBy] {
		sortBy = "created_at"
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	// --- Query data + preload + pagination
	err := baseQuery.
		// Preload("CreatedBy").
		// Preload("LayoutPositions").
		Order(sortBy + " " + sortOrder).
		Limit(pageSize).
		Offset(offset).
		Find(&layouts).Error

	if err != nil {
		return nil, 0, err
	}

	return layouts, total, nil
}

func (r *LayoutRepositoryImpl) CreateLayout(ctx context.Context, layout entities.Layouts) (*entities.Layouts, error) {
	if err := r.db.WithContext(ctx).Create(&layout).Error; err != nil {
		return nil, err
	}
	return &layout, nil
}

func (r *LayoutRepositoryImpl) UpdateLayout(ctx context.Context, layout entities.Layouts) (*entities.Layouts, error) {
	if err := r.db.WithContext(ctx).Model(&entities.Layouts{}).
		Where("layout_id = ?", layout.LayoutID).
		Updates(layout).Error; err != nil {
		return nil, err
	}
	return &layout, nil
}

func (r *LayoutRepositoryImpl) DeleteLayout(ctx context.Context, layoutID uuid.UUID) error {

	return r.db.WithContext(ctx).
		Where("layout_id = ?", layoutID).
		Delete(&entities.Layouts{}).Error
}
