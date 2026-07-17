package repositories

import (
	"context"
	"errors"

	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/ardnh/be-travel-booking-app/internal/domain/repositories"
	"github.com/ardnh/be-travel-booking-app/internal/utils/helpers"
	errorConst "github.com/ardnh/be-travel-booking-app/pkg/errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type scheduleRepositoryImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewScheduleRepository(db *gorm.DB, redis *redis.Client) repositories.ScheduleRepository {
	return &scheduleRepositoryImpl{
		db:    db,
		redis: redis,
	}
}

func (r *scheduleRepositoryImpl) GetScheduleByID(ctx context.Context, scheduleID uuid.UUID) (*entities.Schedules, error) {
	var schedule entities.Schedules
	err := r.db.WithContext(ctx).
		Preload("Layout").
		Preload("ServiceType").
		Preload("OriginPool").
		Preload("DestinationPool").
		Where("schedule_id = ?", scheduleID).
		First(&schedule).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errorConst.ErrNotFound
		}
		return nil, err
	}
	return &schedule, nil
}

func (r *scheduleRepositoryImpl) GetAllSchedules(ctx context.Context, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.Schedules, int64, error) {
	var schedules []entities.Schedules
	var total int64

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize >= 1000 {
		pageSize = 30
	}

	offset := (page - 1) * pageSize

	baseQuery := r.db.WithContext(ctx).
		Model(&entities.Schedules{}).
		Preload("Layout").
		Preload("ServiceType").
		Preload("OriginPool").
		Preload("DestinationPool")

	if search != "" {
		baseQuery = baseQuery.Where("vehicle_type ILIKE ? OR status ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	allowedSort := map[string]bool{
		"vehicle_type":   true,
		"departure_date": true,
		"departure_time": true,
		"price_per_seat": true,
		"total_seat":     true,
		"available_seat": true,
		"status":         true,
		"created_at":     true,
		"updated_at":     true,
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
		Find(&schedules).Error
	if err != nil {
		return nil, 0, err
	}
	return schedules, total, nil
}

func (r *scheduleRepositoryImpl) GetSchedulesByVendorID(ctx context.Context, vendorID uuid.UUID, page int, pageSize int, search string, sortBy string, sortOrder string) ([]entities.Schedules, int64, error) {
	var schedules []entities.Schedules
	var total int64

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize >= 1000 {
		pageSize = 30
	}

	offset := (page - 1) * pageSize

	baseQuery := r.db.WithContext(ctx).
		Model(&entities.Schedules{}).
		Preload("Layout").
		Preload("ServiceType").
		Preload("OriginPool").
		Preload("DestinationPool").
		Where("vendor_id = ?", vendorID)

	if search != "" {
		baseQuery = baseQuery.Where("vehicle_type ILIKE ? OR status ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	allowedSort := map[string]bool{
		"vehicle_type":   true,
		"departure_date": true,
		"departure_time": true,
		"price_per_seat": true,
		"total_seat":     true,
		"available_seat": true,
		"status":         true,
		"created_at":     true,
		"updated_at":     true,
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
		Find(&schedules).Error

	if err != nil {
		return nil, 0, err
	}
	return schedules, total, nil
}

func (r *scheduleRepositoryImpl) CreateSchedule(ctx context.Context, schedule entities.CreateSchedule, total int) error {

	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	// ── validasi kepemilikan, semua di dalam tx ──
	ok, err := r.ValidateServiceType(ctx, schedule.VendorID, schedule.ServiceTypeID)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("service_type does not belong to this vendor")
	}

	allPools := append([]uuid.UUID{schedule.OriginPoolID, schedule.DestinationPoolID})
	found, err := r.ValidatePoolOwnership(ctx, schedule.VendorID, allPools)
	if err != nil {
		return err
	}
	if found != len(allPools) {
		return errors.New("one or more pools do not belong to this vendor")
	}

	seatCount, err := r.CountLayoutSeats(ctx, schedule.LayoutID, schedule.VendorID)
	if err != nil {
		return err
	}
	if seatCount == 0 {
		return errors.New("layout not found, empty, or not owned by this vendor")
	}

	// ── build rows ──
	rows := make([]entities.Schedules, 0, total)
	for _, date := range dates {
		for _, destID := range destIDs {
			for _, t := range d.DepartureTimes {
				price, ok := resolvePrice(t, d.PriceBands)
				if !ok {
					return errors.New("no price band matches departure_time " + t)
				}

				var eta *string
				if d.DurationMinutes != nil {
					eta = helpers.Ptr(helpers.AddMinutes(t, *d.DurationMinutes))
				}

				rows = append(rows, models.Schedules{
					VendorID:             vendorID,
					ServiceTypeID:        serviceTypeID,
					OriginPoolID:         originID,
					DestinationPoolID:    destID,
					LayoutID:             layoutID,
					ScheduleBatchID:      &batchID,
					VehicleType:          d.VehicleType,
					DepartureDate:        date,
					DepartureTime:        t,
					EstimatedArrivalTime: eta,
					PricePerSeat:         price,
					TotalSeat:            seatCount,
					AvailableSeat:        seatCount,
					Status:               "scheduled",
					CreatedBy:            &userID,
				})
			}
		}
	}

}

func (r *scheduleRepositoryImpl) UpdateSchedule(ctx context.Context, schedule entities.Schedules) error {
	err := r.db.WithContext(ctx).Save(&schedule).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *scheduleRepositoryImpl) DeleteSchedule(ctx context.Context, scheduleID uuid.UUID) error {
	err := r.db.WithContext(ctx).Where("schedule_id = ?", scheduleID).Delete(&entities.Schedules{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *scheduleRepositoryImpl) CountLayoutSeats(ctx context.Context, layoutID, vendorID uuid.UUID) (int, error) {
	var count int64
	tx := r.db.WithContext(ctx).Begin()
	err := tx.Table("layout_seats AS ls").
		Joins("JOIN layouts l ON l.layout_id = ls.layout_id").
		Where("ls.layout_id = ? AND l.vendor_id = ? AND ls.is_seat = true", layoutID, vendorID).
		Count(&count).Error
	return int(count), err
}

func (r *scheduleRepositoryImpl) ValidatePoolOwnership(ctx context.Context, vendorID uuid.UUID, poolIDs []uuid.UUID) (int, error) {
	var count int64
	tx := r.db.WithContext(ctx).Begin()
	err := tx.Model(&entities.Pools{}).
		Where("pool_id IN ? AND vendor_id = ?", poolIDs, vendorID).
		Count(&count).Error
	return int(count), err
}

func (r *scheduleRepositoryImpl) ValidateServiceType(ctx context.Context, vendorID, serviceTypeID uuid.UUID) (bool, error) {
	var count int64
	tx := r.db.WithContext(ctx).Begin()
	err := tx.Model(&entities.ServiceTypes{}).
		Where("service_type_id = ? AND vendor_id = ?", serviceTypeID, vendorID).
		Count(&count).Error
	return count > 0, err
}

func (r *scheduleRepositoryImpl) BulkInsert(ctx context.Context, rows []entities.Schedules, overwrite bool) (int64, error) {
	tx := r.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	conflictCols := []clause.Column{
		{Name: "vendor_id"}, {Name: "origin_pool_id"}, {Name: "destination_pool_id"},
		{Name: "departure_date"}, {Name: "departure_time"},
	}

	onConflict := clause.OnConflict{Columns: conflictCols, DoNothing: true}
	if overwrite {
		onConflict = clause.OnConflict{
			Columns: conflictCols,
			DoUpdates: clause.AssignmentColumns([]string{
				"layout_id", "vehicle_type", "estimated_arrival_time",
				"price_per_seat", "total_seat", "available_seat", "updated_at",
			}),
			// jangan timpa jadwal yang sudah ada booking
			Where: clause.Where{Exprs: []clause.Expression{
				gorm.Expr("schedules.available_seat = schedules.total_seat"),
			}},
		}
	}

	res := tx.Clauses(onConflict).CreateInBatches(rows, 500)
	return res.RowsAffected, res.Error
}
