package services

import (
	"context"
	"errors"
	"testing"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockScheduleRepository struct {
	mock.Mock
}

func (m *MockScheduleRepository) GetScheduleByID(ctx context.Context, scheduleID uuid.UUID) (*entities.Schedules, error) {
	args := m.Called(ctx, scheduleID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Schedules), args.Error(1)
}

func (m *MockScheduleRepository) GetAllSchedules(ctx context.Context) ([]entities.Schedules, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Schedules), args.Error(1)
}

func (m *MockScheduleRepository) GetSchedulesByVendorID(ctx context.Context, vendorID uuid.UUID) ([]entities.Schedules, error) {
	args := m.Called(ctx, vendorID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Schedules), args.Error(1)
}

func (m *MockScheduleRepository) CreateSchedule(ctx context.Context, schedule entities.Schedules) error {
	args := m.Called(ctx, schedule)
	return args.Error(0)
}

func (m *MockScheduleRepository) UpdateSchedule(ctx context.Context, schedule entities.Schedules) error {
	args := m.Called(ctx, schedule)
	return args.Error(0)
}

func (m *MockScheduleRepository) DeleteSchedule(ctx context.Context, scheduleID uuid.UUID) error {
	args := m.Called(ctx, scheduleID)
	return args.Error(0)
}

// ─── GetScheduleByID ──────────────────────────────────

func TestGetScheduleByID_Success(t *testing.T) {
	mockRepo := new(MockScheduleRepository)
	logger := logrus.New()
	service := NewScheduleServiceImpl(mockRepo, logger)

	scheduleID := uuid.New()
	existing := &entities.Schedules{
		ScheduleID:     scheduleID,
		VehicleType:    "Bus",
		DepartureDate:  "2026-06-01",
		DepartureTime:  "08:00:00",
		PricePerSeat:   100000,
		TotalSeat:      40,
		AvailableSeat:  35,
		Status:         "scheduled",
		VendorID:       uuid.New(),
		ServiceTypeID:  uuid.New(),
		LayoutID:       uuid.New(),
		OriginPoolID:   uuid.New(),
		DestinationPoolID: uuid.New(),
	}

	mockRepo.On("GetScheduleByID", mock.Anything, scheduleID).Return(existing, nil)

	schedule, err := service.GetScheduleByID(context.Background(), scheduleID)

	assert.NoError(t, err)
	assert.Equal(t, existing, schedule)
	mockRepo.AssertExpectations(t)
}

func TestGetScheduleByID_NotFound(t *testing.T) {
	mockRepo := new(MockScheduleRepository)
	logger := logrus.New()
	service := NewScheduleServiceImpl(mockRepo, logger)

	scheduleID := uuid.New()
	mockRepo.On("GetScheduleByID", mock.Anything, scheduleID).Return(nil, errors.New("not found"))

	schedule, err := service.GetScheduleByID(context.Background(), scheduleID)

	assert.Error(t, err)
	assert.Nil(t, schedule)
	mockRepo.AssertExpectations(t)
}

// ─── GetAllSchedules ──────────────────────────────────

func TestGetAllSchedules_Success(t *testing.T) {
	mockRepo := new(MockScheduleRepository)
	logger := logrus.New()
	service := NewScheduleServiceImpl(mockRepo, logger)

	schedules := []entities.Schedules{
		{
			ScheduleID:    uuid.New(),
			VehicleType:   "Bus",
			DepartureDate: "2026-06-01",
			DepartureTime: "08:00:00",
			PricePerSeat:  100000,
			TotalSeat:     40,
			AvailableSeat: 35,
			Status:        "scheduled",
		},
		{
			ScheduleID:    uuid.New(),
			VehicleType:   "Minibus",
			DepartureDate: "2026-06-02",
			DepartureTime: "10:00:00",
			PricePerSeat:  75000,
			TotalSeat:     20,
			AvailableSeat: 18,
			Status:        "scheduled",
		},
	}

	mockRepo.On("GetAllSchedules", mock.Anything).Return(schedules, nil)

	result, err := service.GetAllSchedules(context.Background())

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestGetAllSchedules_Error(t *testing.T) {
	mockRepo := new(MockScheduleRepository)
	logger := logrus.New()
	service := NewScheduleServiceImpl(mockRepo, logger)

	mockRepo.On("GetAllSchedules", mock.Anything).Return(nil, errors.New("db error"))

	result, err := service.GetAllSchedules(context.Background())

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// ─── GetSchedulesByVendorID ──────────────────────────────────

func TestGetSchedulesByVendorID_Success(t *testing.T) {
	mockRepo := new(MockScheduleRepository)
	logger := logrus.New()
	service := NewScheduleServiceImpl(mockRepo, logger)

	vendorID := uuid.New()
	schedules := []entities.Schedules{
		{
			ScheduleID:    uuid.New(),
			VendorID:      vendorID,
			VehicleType:   "Bus",
			DepartureDate: "2026-06-01",
			DepartureTime: "08:00:00",
			PricePerSeat:  100000,
			TotalSeat:     40,
			AvailableSeat: 35,
			Status:        "scheduled",
		},
	}

	mockRepo.On("GetSchedulesByVendorID", mock.Anything, vendorID).Return(schedules, nil)

	result, err := service.GetSchedulesByVendorID(context.Background(), vendorID)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, vendorID, result[0].VendorID)
	mockRepo.AssertExpectations(t)
}

func TestGetSchedulesByVendorID_Error(t *testing.T) {
	mockRepo := new(MockScheduleRepository)
	logger := logrus.New()
	service := NewScheduleServiceImpl(mockRepo, logger)

	vendorID := uuid.New()
	mockRepo.On("GetSchedulesByVendorID", mock.Anything, vendorID).Return(nil, errors.New("db error"))

	result, err := service.GetSchedulesByVendorID(context.Background(), vendorID)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// ─── CreateSchedule ──────────────────────────────────

func TestCreateSchedule_Success(t *testing.T) {
	mockRepo := new(MockScheduleRepository)
	logger := logrus.New()
	service := NewScheduleServiceImpl(mockRepo, logger)

	vendorID := uuid.New()
	serviceTypeID := uuid.New()
	layoutID := uuid.New()
	originPoolID := uuid.New()
	destPoolID := uuid.New()

	price := 100000
	totalSeat := 40
	availableSeat := 35
	status := "scheduled"

	req := dto.CreateScheduleDTO{
		VendorID:          vendorID.String(),
		ServiceTypeID:     serviceTypeID.String(),
		OriginPoolID:      originPoolID.String(),
		DestinationPoolID: destPoolID.String(),
		LayoutID:          layoutID.String(),
		VehicleType:       "Bus",
		DepartureDate:     "2026-06-01",
		DepartureTime:     "08:00:00",
		PricePerSeat:      &price,
		TotalSeat:         &totalSeat,
		AvailableSeat:     &availableSeat,
		Status:            &status,
	}

	mockRepo.On("CreateSchedule", mock.Anything, mock.AnythingOfType("entities.Schedules")).Return(nil)

	err := service.CreateSchedule(context.Background(), req)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateSchedule_Error(t *testing.T) {
	mockRepo := new(MockScheduleRepository)
	logger := logrus.New()
	service := NewScheduleServiceImpl(mockRepo, logger)

	req := dto.CreateScheduleDTO{
		VendorID:      uuid.New().String(),
		ServiceTypeID: uuid.New().String(),
		LayoutID:      uuid.New().String(),
		OriginPoolID:  uuid.New().String(),
		DepartureDate: "2026-06-01",
		DepartureTime: "08:00:00",
		PricePerSeat:  float64Ptr(100000),
		TotalSeat:     intPtr(40),
	}

	mockRepo.On("CreateSchedule", mock.Anything, mock.AnythingOfType("entities.Schedules")).Return(errors.New("db error"))

	err := service.CreateSchedule(context.Background(), req)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

// ─── UpdateSchedule ──────────────────────────────────

func TestUpdateSchedule_Success(t *testing.T) {
	mockRepo := new(MockScheduleRepository)
	logger := logrus.New()
	service := NewScheduleServiceImpl(mockRepo, logger)

	scheduleID := uuid.New()
	existing := &entities.Schedules{
		ScheduleID:    scheduleID,
		VehicleType:   "Bus",
		DepartureDate: "2026-06-01",
		DepartureTime: "08:00:00",
		PricePerSeat:  100000,
		TotalSeat:     40,
		AvailableSeat: 35,
		Status:        "scheduled",
	}

	status := "delayed"
	req := dto.UpdateScheduleDTO{
		VehicleType: strPtr("Executive Bus"),
		PricePerSeat: float64Ptr(125000),
		AvailableSeat: intPtr(30),
		Status:        &status,
	}

	mockRepo.On("GetScheduleByID", mock.Anything, scheduleID).Return(existing, nil)
	mockRepo.On("UpdateSchedule", mock.Anything, mock.AnythingOfType("entities.Schedules")).Return(nil)

	err := service.UpdateSchedule(context.Background(), scheduleID, req)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateSchedule_NotFound(t *testing.T) {
	mockRepo := new(MockScheduleRepository)
	logger := logrus.New()
	service := NewScheduleServiceImpl(mockRepo, logger)

	scheduleID := uuid.New()
	req := dto.UpdateScheduleDTO{
		VehicleType: strPtr("Bus"),
	}

	mockRepo.On("GetScheduleByID", mock.Anything, scheduleID).Return(nil, errors.New("not found"))

	err := service.UpdateSchedule(context.Background(), scheduleID, req)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateSchedule_UpdateError(t *testing.T) {
	mockRepo := new(MockScheduleRepository)
	logger := logrus.New()
	service := NewScheduleServiceImpl(mockRepo, logger)

	scheduleID := uuid.New()
	existing := &entities.Schedules{
		ScheduleID:    scheduleID,
		VehicleType:   "Bus",
		DepartureDate: "2026-06-01",
		DepartureTime: "08:00:00",
	}

	req := dto.UpdateScheduleDTO{
		VehicleType: strPtr("Executive Bus"),
	}

	mockRepo.On("GetScheduleByID", mock.Anything, scheduleID).Return(existing, nil)
	mockRepo.On("UpdateSchedule", mock.Anything, mock.AnythingOfType("entities.Schedules")).Return(errors.New("update error"))

	err := service.UpdateSchedule(context.Background(), scheduleID, req)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

// ─── DeleteSchedule ──────────────────────────────────

func TestDeleteSchedule_Success(t *testing.T) {
	mockRepo := new(MockScheduleRepository)
	logger := logrus.New()
	service := NewScheduleServiceImpl(mockRepo, logger)

	scheduleID := uuid.New()
	mockRepo.On("DeleteSchedule", mock.Anything, scheduleID).Return(nil)

	err := service.DeleteSchedule(context.Background(), scheduleID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteSchedule_Error(t *testing.T) {
	mockRepo := new(MockScheduleRepository)
	logger := logrus.New()
	service := NewScheduleServiceImpl(mockRepo, logger)

	scheduleID := uuid.New()
	mockRepo.On("DeleteSchedule", mock.Anything, scheduleID).Return(errors.New("delete error"))

	err := service.DeleteSchedule(context.Background(), scheduleID)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

// ─── Helpers ─────────────────────────────────────────

func strPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func float64Ptr(f float64) *float64 {
	return &f
}
