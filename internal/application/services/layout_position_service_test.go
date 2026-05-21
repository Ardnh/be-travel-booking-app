package services

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
// 	"github.com/ardnh/be-travel-booking-app/internal/domain/entities"
// 	"github.com/google/uuid"
// 	"github.com/sirupsen/logrus"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockLayoutPositionRepository struct {
// 	mock.Mock
// }

// func (m *MockLayoutPositionRepository) GetLayoutPositionByID(ctx context.Context, layoutPositionID uuid.UUID) (*entities.LayoutPositions, error) {
// 	args := m.Called(ctx, layoutPositionID)
// 	if args.Get(0) == nil {
// 		return nil, args.Error(1)
// 	}
// 	return args.Get(0).(*entities.LayoutPositions), args.Error(1)
// }

// func (m *MockLayoutPositionRepository) GetLayoutPositionsByLayoutID(ctx context.Context, layoutID uuid.UUID) ([]entities.LayoutPositions, error) {
// 	args := m.Called(ctx, layoutID)
// 	if args.Get(0) == nil {
// 		return nil, args.Error(1)
// 	}
// 	return args.Get(0).([]entities.LayoutPositions), args.Error(1)
// }

// func (m *MockLayoutPositionRepository) GetAllLayoutPositions(ctx context.Context) ([]entities.LayoutPositions, error) {
// 	args := m.Called(ctx)
// 	if args.Get(0) == nil {
// 		return nil, args.Error(1)
// 	}
// 	return args.Get(0).([]entities.LayoutPositions), args.Error(1)
// }

// func (m *MockLayoutPositionRepository) CreateLayoutPosition(ctx context.Context, layoutPosition entities.LayoutPositions) error {
// 	args := m.Called(ctx, layoutPosition)
// 	return args.Error(0)
// }

// func (m *MockLayoutPositionRepository) UpdateLayoutPosition(ctx context.Context, layoutPosition entities.LayoutPositions) error {
// 	args := m.Called(ctx, layoutPosition)
// 	return args.Error(0)
// }

// func (m *MockLayoutPositionRepository) DeleteLayoutPosition(ctx context.Context, layoutPositionID uuid.UUID) error {
// 	args := m.Called(ctx, layoutPositionID)
// 	return args.Error(0)
// }

// // ─── GetLayoutPositionByID ───────────────────────────────

// func TestGetLayoutPositionByID_Success(t *testing.T) {
// 	mockRepo := new(MockLayoutPositionRepository)
// 	logger := logrus.New()
// 	service := NewLayoutPositionServiceImpl(mockRepo, logger)

// 	layoutID := uuid.New()
// 	layoutPosID := uuid.New()
// 	existing := &entities.LayoutPositions{
// 		LayoutPositionID: layoutPosID,
// 		LayoutID:         layoutID,
// 		Label:            "A1",
// 		Row:              1,
// 		Col:              1,
// 		PositionType:     "seat",
// 		IsUsed:           true,
// 	}

// 	mockRepo.On("GetLayoutPositionByID", mock.Anything, layoutPosID).Return(existing, nil)

// 	position, err := service.GetLayoutPositionByID(context.Background(), layoutPosID)

// 	assert.NoError(t, err)
// 	assert.Equal(t, existing, position)
// 	mockRepo.AssertExpectations(t)
// }

// func TestGetLayoutPositionByID_NotFound(t *testing.T) {
// 	mockRepo := new(MockLayoutPositionRepository)
// 	logger := logrus.New()
// 	service := NewLayoutPositionServiceImpl(mockRepo, logger)

// 	layoutPosID := uuid.New()
// 	mockRepo.On("GetLayoutPositionByID", mock.Anything, layoutPosID).Return(nil, errors.New("not found"))

// 	position, err := service.GetLayoutPositionByID(context.Background(), layoutPosID)

// 	assert.Error(t, err)
// 	assert.Nil(t, position)
// 	mockRepo.AssertExpectations(t)
// }

// // ─── GetLayoutPositionsByLayoutID ───────────────────────────────

// func TestGetLayoutPositionsByLayoutID_Success(t *testing.T) {
// 	mockRepo := new(MockLayoutPositionRepository)
// 	logger := logrus.New()
// 	service := NewLayoutPositionServiceImpl(mockRepo, logger)

// 	layoutID := uuid.New()
// 	positions := []entities.LayoutPositions{
// 		{
// 			LayoutPositionID: uuid.New(),
// 			LayoutID:         layoutID,
// 			Label:            "A1",
// 			Row:              1,
// 			Col:              1,
// 			PositionType:     "seat",
// 			IsUsed:           true,
// 		},
// 		{
// 			LayoutPositionID: uuid.New(),
// 			LayoutID:         layoutID,
// 			Label:            "A2",
// 			Row:              1,
// 			Col:              2,
// 			PositionType:     "seat",
// 			IsUsed:           true,
// 		},
// 	}

// 	mockRepo.On("GetLayoutPositionsByLayoutID", mock.Anything, layoutID).Return(positions, nil)

// 	result, err := service.GetLayoutPositionsByLayoutID(context.Background(), layoutID)

// 	assert.NoError(t, err)
// 	assert.Len(t, result, 2)
// 	mockRepo.AssertExpectations(t)
// }

// func TestGetLayoutPositionsByLayoutID_Error(t *testing.T) {
// 	mockRepo := new(MockLayoutPositionRepository)
// 	logger := logrus.New()
// 	service := NewLayoutPositionServiceImpl(mockRepo, logger)

// 	layoutID := uuid.New()
// 	mockRepo.On("GetLayoutPositionsByLayoutID", mock.Anything, layoutID).Return(nil, errors.New("db error"))

// 	result, err := service.GetLayoutPositionsByLayoutID(context.Background(), layoutID)

// 	assert.Error(t, err)
// 	assert.Nil(t, result)
// 	mockRepo.AssertExpectations(t)
// }

// // ─── GetAllLayoutPositions ───────────────────────────────

// func TestGetAllLayoutPositions_Success(t *testing.T) {
// 	mockRepo := new(MockLayoutPositionRepository)
// 	logger := logrus.New()
// 	service := NewLayoutPositionServiceImpl(mockRepo, logger)

// 	positions := []entities.LayoutPositions{
// 		{
// 			LayoutPositionID: uuid.New(),
// 			Label:            "A1",
// 			Row:              1,
// 			Col:              1,
// 			PositionType:     "seat",
// 			IsUsed:           true,
// 		},
// 	}

// 	mockRepo.On("GetAllLayoutPositions", mock.Anything).Return(positions, nil)

// 	result, err := service.GetAllLayoutPositions(context.Background())

// 	assert.NoError(t, err)
// 	assert.Len(t, result, 1)
// 	mockRepo.AssertExpectations(t)
// }

// func TestGetAllLayoutPositions_Error(t *testing.T) {
// 	mockRepo := new(MockLayoutPositionRepository)
// 	logger := logrus.New()
// 	service := NewLayoutPositionServiceImpl(mockRepo, logger)

// 	mockRepo.On("GetAllLayoutPositions", mock.Anything).Return(nil, errors.New("db error"))

// 	result, err := service.GetAllLayoutPositions(context.Background())

// 	assert.Error(t, err)
// 	assert.Nil(t, result)
// 	mockRepo.AssertExpectations(t)
// }

// // ─── CreateLayoutPosition ───────────────────────────────

// func TestCreateLayoutPosition_Success(t *testing.T) {
// 	mockRepo := new(MockLayoutPositionRepository)
// 	logger := logrus.New()
// 	service := NewLayoutPositionServiceImpl(mockRepo, logger)

// 	layoutID := uuid.New()
// 	req := dto.CreateLayoutPositionDTO{
// 		LayoutID:     layoutID.String(),
// 		Label:        "A1",
// 		Row:          1,
// 		Col:          1,
// 		PositionType: "seat",
// 		IsUsed:       true,
// 	}

// 	mockRepo.On("CreateLayoutPosition", mock.Anything, mock.AnythingOfType("entities.LayoutPositions")).Return(nil)

// 	err := service.CreateLayoutPosition(context.Background(), req, layoutID)

// 	assert.NoError(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func TestCreateLayoutPosition_Error(t *testing.T) {
// 	mockRepo := new(MockLayoutPositionRepository)
// 	logger := logrus.New()
// 	service := NewLayoutPositionServiceImpl(mockRepo, logger)

// 	layoutID := uuid.New()
// 	req := dto.CreateLayoutPositionDTO{
// 		LayoutID:     layoutID.String(),
// 		Label:        "A1",
// 		Row:          1,
// 		Col:          1,
// 		PositionType: "seat",
// 		IsUsed:       true,
// 	}

// 	mockRepo.On("CreateLayoutPosition", mock.Anything, mock.AnythingOfType("entities.LayoutPositions")).Return(errors.New("db error"))

// 	err := service.CreateLayoutPosition(context.Background(), req, layoutID)

// 	assert.Error(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// // ─── UpdateLayoutPosition ───────────────────────────────

// func TestUpdateLayoutPosition_Success(t *testing.T) {
// 	mockRepo := new(MockLayoutPositionRepository)
// 	logger := logrus.New()
// 	service := NewLayoutPositionServiceImpl(mockRepo, logger)

// 	layoutPosID := uuid.New()
// 	existing := &entities.LayoutPositions{
// 		LayoutPositionID: layoutPosID,
// 		Label:            "A1",
// 		Row:              1,
// 		Col:              1,
// 		PositionType:     "seat",
// 		IsUsed:           true,
// 	}

// 	req := dto.UpdateLayoutPositionDTO{
// 		Label:  strPtr("B1"),
// 		Row:    intPtr(2),
// 		Col:    intPtr(1),
// 		IsUsed: boolPtr(false),
// 	}

// 	mockRepo.On("GetLayoutPositionByID", mock.Anything, layoutPosID).Return(existing, nil)
// 	mockRepo.On("UpdateLayoutPosition", mock.Anything, mock.AnythingOfType("entities.LayoutPositions")).Return(nil)

// 	err := service.UpdateLayoutPosition(context.Background(), layoutPosID, req)

// 	assert.NoError(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func TestUpdateLayoutPosition_NotFound(t *testing.T) {
// 	mockRepo := new(MockLayoutPositionRepository)
// 	logger := logrus.New()
// 	service := NewLayoutPositionServiceImpl(mockRepo, logger)

// 	layoutPosID := uuid.New()
// 	req := dto.UpdateLayoutPositionDTO{
// 		Label: strPtr("B1"),
// 	}

// 	mockRepo.On("GetLayoutPositionByID", mock.Anything, layoutPosID).Return(nil, errors.New("not found"))

// 	err := service.UpdateLayoutPosition(context.Background(), layoutPosID, req)

// 	assert.Error(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func TestUpdateLayoutPosition_UpdateError(t *testing.T) {
// 	mockRepo := new(MockLayoutPositionRepository)
// 	logger := logrus.New()
// 	service := NewLayoutPositionServiceImpl(mockRepo, logger)

// 	layoutPosID := uuid.New()
// 	existing := &entities.LayoutPositions{
// 		LayoutPositionID: layoutPosID,
// 		Label:            "A1",
// 		Row:              1,
// 		Col:              1,
// 		PositionType:     "seat",
// 		IsUsed:           true,
// 	}

// 	req := dto.UpdateLayoutPositionDTO{
// 		Label: strPtr("B1"),
// 	}

// 	mockRepo.On("GetLayoutPositionByID", mock.Anything, layoutPosID).Return(existing, nil)
// 	mockRepo.On("UpdateLayoutPosition", mock.Anything, mock.AnythingOfType("entities.LayoutPositions")).Return(errors.New("update error"))

// 	err := service.UpdateLayoutPosition(context.Background(), layoutPosID, req)

// 	assert.Error(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// // ─── DeleteLayoutPosition ───────────────────────────────

// func TestDeleteLayoutPosition_Success(t *testing.T) {
// 	mockRepo := new(MockLayoutPositionRepository)
// 	logger := logrus.New()
// 	service := NewLayoutPositionServiceImpl(mockRepo, logger)

// 	layoutPosID := uuid.New()
// 	mockRepo.On("DeleteLayoutPosition", mock.Anything, layoutPosID).Return(nil)

// 	err := service.DeleteLayoutPosition(context.Background(), layoutPosID)

// 	assert.NoError(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func TestDeleteLayoutPosition_Error(t *testing.T) {
// 	mockRepo := new(MockLayoutPositionRepository)
// 	logger := logrus.New()
// 	service := NewLayoutPositionServiceImpl(mockRepo, logger)

// 	layoutPosID := uuid.New()
// 	mockRepo.On("DeleteLayoutPosition", mock.Anything, layoutPosID).Return(errors.New("delete error"))

// 	err := service.DeleteLayoutPosition(context.Background(), layoutPosID)

// 	assert.Error(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// // ─── Helpers ───────────────────────────────────────────

// func boolPtr(b bool) *bool {
// 	return &b
// }
