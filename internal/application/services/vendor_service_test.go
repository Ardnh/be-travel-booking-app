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

// type MockVendorRepository struct {
// 	mock.Mock
// }

// func (m *MockVendorRepository) GetVendorByID(ctx context.Context, vendorID uuid.UUID) (*entities.Vendors, error) {
// 	args := m.Called(ctx, vendorID)
// 	if args.Get(0) == nil {
// 		return nil, args.Error(1)
// 	}
// 	return args.Get(0).(*entities.Vendors), args.Error(1)
// }

// func (m *MockVendorRepository) GetAllVendors(ctx context.Context) ([]entities.Vendors, error) {
// 	args := m.Called(ctx)
// 	if args.Get(0) == nil {
// 		return nil, args.Error(1)
// 	}
// 	return args.Get(0).([]entities.Vendors), args.Error(1)
// }

// func (m *MockVendorRepository) CreateVendor(ctx context.Context, vendor entities.Vendors) error {
// 	args := m.Called(ctx, vendor)
// 	return args.Error(0)
// }

// func (m *MockVendorRepository) UpdateVendor(ctx context.Context, vendor entities.Vendors) error {
// 	args := m.Called(ctx, vendor)
// 	return args.Error(0)
// }

// func (m *MockVendorRepository) DeleteVendor(ctx context.Context, vendorID uuid.UUID) error {
// 	args := m.Called(ctx, vendorID)
// 	return args.Error(0)
// }

// // ─── GetVendorByID ────────────────────────────────────────

// func TestGetVendorByID_Success(t *testing.T) {
// 	mockRepo := new(MockVendorRepository)
// 	logger := logrus.New()
// 	service := NewVendorServiceImpl(mockRepo, logger)

// 	vendorID := uuid.New()
// 	existingVendor := &entities.Vendors{
// 		VendorID:     vendorID,
// 		BusinessName: "Travel Corp",
// 		OwnerName:    "John Owner",
// 	}

// 	mockRepo.On("GetVendorByID", mock.Anything, vendorID).Return(existingVendor, nil)

// 	vendor, err := service.GetVendorByID(context.Background(), vendorID)

// 	assert.NoError(t, err)
// 	assert.Equal(t, existingVendor, vendor)
// 	mockRepo.AssertExpectations(t)
// }

// func TestGetVendorByID_NotFound(t *testing.T) {
// 	mockRepo := new(MockVendorRepository)
// 	logger := logrus.New()
// 	service := NewVendorServiceImpl(mockRepo, logger)

// 	vendorID := uuid.New()
// 	mockRepo.On("GetVendorByID", mock.Anything, vendorID).Return(nil, errors.New("not found"))

// 	vendor, err := service.GetVendorByID(context.Background(), vendorID)

// 	assert.Error(t, err)
// 	assert.Nil(t, vendor)
// 	mockRepo.AssertExpectations(t)
// }

// // ─── GetAllVendors ────────────────────────────────────────

// func TestGetAllVendors_Success(t *testing.T) {
// 	mockRepo := new(MockVendorRepository)
// 	logger := logrus.New()
// 	service := NewVendorServiceImpl(mockRepo, logger)

// 	vendors := []entities.Vendors{
// 		{
// 			VendorID:     uuid.New(),
// 			BusinessName: "Travel Corp",
// 			OwnerName:    "John Owner",
// 		},
// 		{
// 			VendorID:     uuid.New(),
// 			BusinessName: "Booking Inc",
// 			OwnerName:    "Jane Owner",
// 		},
// 	}

// 	mockRepo.On("GetAllVendors", mock.Anything).Return(vendors, nil)

// 	result, err := service.GetAllVendors(context.Background())

// 	assert.NoError(t, err)
// 	assert.Len(t, result, 2)
// 	assert.Equal(t, "Travel Corp", result[0].BusinessName)
// 	assert.Equal(t, "Booking Inc", result[1].BusinessName)
// 	mockRepo.AssertExpectations(t)
// }

// func TestGetAllVendors_Error(t *testing.T) {
// 	mockRepo := new(MockVendorRepository)
// 	logger := logrus.New()
// 	service := NewVendorServiceImpl(mockRepo, logger)

// 	mockRepo.On("GetAllVendors", mock.Anything).Return(nil, errors.New("db error"))

// 	result, err := service.GetAllVendors(context.Background())

// 	assert.Error(t, err)
// 	assert.Nil(t, result)
// 	mockRepo.AssertExpectations(t)
// }

// // ─── CreateVendor ────────────────────────────────────────

// func TestCreateVendor_Success(t *testing.T) {
// 	mockRepo := new(MockVendorRepository)
// 	logger := logrus.New()
// 	service := NewVendorServiceImpl(mockRepo, logger)

// 	legalDocNumber := "1234567890"
// 	req := dto.CreateVendorDTO{
// 		BusinessName:        "Travel Corp",
// 		OwnerName:           "John Owner",
// 		Description:         "A travel company",
// 		FoundedYear:         2020,
// 		PhoneNumber:         "1234567890",
// 		Email:               "owner@travel.com",
// 		HeadOfficeAddress:   "123 Main St",
// 		LegalDocumentNumber: &legalDocNumber,
// 		Status:              "active",
// 	}

// 	mockRepo.On("CreateVendor", mock.Anything, mock.AnythingOfType("entities.Vendors")).Return(nil)

// 	err := service.CreateVendor(context.Background(), req)

// 	assert.NoError(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func TestCreateVendor_Error(t *testing.T) {
// 	mockRepo := new(MockVendorRepository)
// 	logger := logrus.New()
// 	service := NewVendorServiceImpl(mockRepo, logger)

// 	req := dto.CreateVendorDTO{
// 		BusinessName: "Travel Corp",
// 		OwnerName:    "John Owner",
// 		Email:        "owner@travel.com",
// 	}

// 	mockRepo.On("CreateVendor", mock.Anything, mock.AnythingOfType("entities.Vendors")).Return(errors.New("db error"))

// 	err := service.CreateVendor(context.Background(), req)

// 	assert.Error(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// // ─── UpdateVendor ────────────────────────────────────────

// func TestUpdateVendor_Success(t *testing.T) {
// 	mockRepo := new(MockVendorRepository)
// 	logger := logrus.New()
// 	service := NewVendorServiceImpl(mockRepo, logger)

// 	vendorID := uuid.New()
// 	existingVendor := &entities.Vendors{
// 		VendorID:     vendorID,
// 		BusinessName: "Travel Corp",
// 		OwnerName:    "John Owner",
// 	}

// 	req := dto.UpdateVendorDTO{
// 		BusinessName:  strPtr("Travel Corp Updated"),
// 		OwnerName:     strPtr("Jane Owner"),
// 		PhoneNumber:   strPtr("0987654321"),
// 		HeadOfficeAddress: strPtr("456 Elm St"),
// 	}

// 	mockRepo.On("GetVendorByID", mock.Anything, vendorID).Return(existingVendor, nil)
// 	mockRepo.On("UpdateVendor", mock.Anything, mock.AnythingOfType("entities.Vendors")).Return(nil)

// 	err := service.UpdateVendor(context.Background(), vendorID, req)

// 	assert.NoError(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func TestUpdateVendor_VendorNotFound(t *testing.T) {
// 	mockRepo := new(MockVendorRepository)
// 	logger := logrus.New()
// 	service := NewVendorServiceImpl(mockRepo, logger)

// 	vendorID := uuid.New()
// 	req := dto.UpdateVendorDTO{
// 		BusinessName: strPtr("Updated Name"),
// 	}

// 	mockRepo.On("GetVendorByID", mock.Anything, vendorID).Return(nil, errors.New("not found"))

// 	err := service.UpdateVendor(context.Background(), vendorID, req)

// 	assert.Error(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func TestUpdateVendor_UpdateError(t *testing.T) {
// 	mockRepo := new(MockVendorRepository)
// 	logger := logrus.New()
// 	service := NewVendorServiceImpl(mockRepo, logger)

// 	vendorID := uuid.New()
// 	existingVendor := &entities.Vendors{
// 		VendorID:     vendorID,
// 		BusinessName: "Travel Corp",
// 		OwnerName:    "John Owner",
// 	}

// 	req := dto.UpdateVendorDTO{
// 		BusinessName: strPtr("Travel Corp Updated"),
// 	}

// 	mockRepo.On("GetVendorByID", mock.Anything, vendorID).Return(existingVendor, nil)
// 	mockRepo.On("UpdateVendor", mock.Anything, mock.AnythingOfType("entities.Vendors")).Return(errors.New("update error"))

// 	err := service.UpdateVendor(context.Background(), vendorID, req)

// 	assert.Error(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// // ─── DeleteVendor ────────────────────────────────────────

// func TestDeleteVendor_Success(t *testing.T) {
// 	mockRepo := new(MockVendorRepository)
// 	logger := logrus.New()
// 	service := NewVendorServiceImpl(mockRepo, logger)

// 	vendorID := uuid.New()

// 	mockRepo.On("DeleteVendor", mock.Anything, vendorID).Return(nil)

// 	err := service.DeleteVendor(context.Background(), vendorID)

// 	assert.NoError(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func TestDeleteVendor_Error(t *testing.T) {
// 	mockRepo := new(MockVendorRepository)
// 	logger := logrus.New()
// 	service := NewVendorServiceImpl(mockRepo, logger)

// 	vendorID := uuid.New()

// 	mockRepo.On("DeleteVendor", mock.Anything, vendorID).Return(errors.New("delete error"))

// 	err := service.DeleteVendor(context.Background(), vendorID)

// 	assert.Error(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func strPtr(s string) *string {
// 	return &s
// }
