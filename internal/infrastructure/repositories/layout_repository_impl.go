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

func (r *LayoutRepositoryImpl) GetLayoutById(ctx context.Context, layoutID uuid.UUID) (*entities.Layout, error) {
	var layout entities.Layout

	err := r.db.WithContext(ctx).
		Preload("Creator").
		Preload("LayoutPositions").
		First(&layout, "layout_id = ?", layoutID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}

	return &layout, nil
}

func (r *LayoutRepositoryImpl) GetLayout(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]*entities.Layout, int64, error) {

	var (
		layouts []*entities.Layout
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
		Model(&entities.Layout{})

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
		Preload("Creator").
		Preload("LayoutPositions").
		Order(sortBy + " " + sortOrder).
		Limit(pageSize).
		Offset(offset).
		Find(&layouts).Error

	if err != nil {
		return nil, 0, err
	}

	return layouts, total, nil
}

func (r *LayoutRepositoryImpl) CreateLayout(ctx context.Context, layout *entities.Layout, layoutPositions []*entities.LayoutPosition) error {

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// 1. Create layout dulu → LayoutID keisi
		if err := tx.Create(layout).Error; err != nil {
			return err
		}

		// 2. Inject FK ke semua layout positions
		for _, lp := range layoutPositions {
			lp.LayoutID = layout.LayoutID
		}

		// 3. Batch insert (jauh lebih cepat)
		if err := tx.CreateInBatches(layoutPositions, 100).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *LayoutRepositoryImpl) UpdateLayout(ctx context.Context, layout *entities.Layout, layoutPositions []*entities.LayoutPosition) error {

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// 1. Update parent (hanya kolom perlu)
		if err := tx.Model(&entities.Layout{}).
			Where("layout_id = ?", layout.LayoutID).
			Updates(layout).Error; err != nil {
			return err
		}

		// 2. Delete semua children lama
		if err := tx.
			Where("layout_id = ?", layout.LayoutID).
			Delete(&entities.LayoutPosition{}).Error; err != nil {
			return err
		}

		// 3. Inject FK
		for _, lp := range layoutPositions {
			lp.LayoutID = layout.LayoutID
		}

		// 4. Insert ulang batch
		if len(layoutPositions) > 0 {
			if err := tx.CreateInBatches(layoutPositions, 100).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *LayoutRepositoryImpl) DeleteLayout(ctx context.Context, layoutID uuid.UUID) error {

	return r.db.WithContext(ctx).
		Where("layout_id = ?", layoutID).
		Delete(&entities.Layout{}).Error
}
