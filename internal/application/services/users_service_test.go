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

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.Users, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Users), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (*entities.Users, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Users), args.Error(1)
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user entities.Users) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user entities.Users) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	args := m.Called(ctx, userID)
	return args.Error(0)
}

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	logger := logrus.New()
	service := NewUsersServiceImpl(mockRepo, logger)

	req := dto.CreateUserDTO{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
		Phone:    "1234567890",
	}

	mockRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("entities.Users")).Return(nil)

	err := service.CreateUser(context.Background(), req)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)
	logger := logrus.New()
	service := NewUsersServiceImpl(mockRepo, logger)

	req := dto.CreateUserDTO{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
		Phone:    "1234567890",
	}

	mockRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("entities.Users")).Return(errors.New("db error"))

	err := service.CreateUser(context.Background(), req)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	logger := logrus.New()
	service := NewUsersServiceImpl(mockRepo, logger)

	userID := uuid.New()
	existingUser := &entities.Users{
		UserID: userID,
		Name:   "John Doe",
		Phone:  "1234567890",
	}

	req := dto.UpdateUserDTO{
		Name:  "Jane Doe",
		Phone: "0987654321",
	}

	mockRepo.On("GetUserByID", mock.Anything, userID).Return(existingUser, nil)
	mockRepo.On("UpdateUser", mock.Anything, mock.AnythingOfType("entities.Users")).Return(nil)

	err := service.UpdateUser(context.Background(), userID.String(), req)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_InvalidUserID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	logger := logrus.New()
	service := NewUsersServiceImpl(mockRepo, logger)

	req := dto.UpdateUserDTO{
		Name:  "Jane Doe",
		Phone: "0987654321",
	}

	err := service.UpdateUser(context.Background(), "invalid-uuid", req)

	assert.Error(t, err)
	mockRepo.AssertNotCalled(t, "GetUserByID")
}

func TestUpdateUser_UserNotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	logger := logrus.New()
	service := NewUsersServiceImpl(mockRepo, logger)

	userID := uuid.New()

	req := dto.UpdateUserDTO{
		Name:  "Jane Doe",
		Phone: "0987654321",
	}

	mockRepo.On("GetUserByID", mock.Anything, userID).Return(nil, errors.New("not found"))

	err := service.UpdateUser(context.Background(), userID.String(), req)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_UpdateError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	logger := logrus.New()
	service := NewUsersServiceImpl(mockRepo, logger)

	userID := uuid.New()
	existingUser := &entities.Users{
		UserID: userID,
		Name:   "John Doe",
		Phone:  "1234567890",
	}

	req := dto.UpdateUserDTO{
		Name:  "Jane Doe",
		Phone: "0987654321",
	}

	mockRepo.On("GetUserByID", mock.Anything, userID).Return(existingUser, nil)
	mockRepo.On("UpdateUser", mock.Anything, mock.AnythingOfType("entities.Users")).Return(errors.New("update error"))

	err := service.UpdateUser(context.Background(), userID.String(), req)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	logger := logrus.New()
	service := NewUsersServiceImpl(mockRepo, logger)

	userID := uuid.New()

	mockRepo.On("DeleteUser", mock.Anything, userID).Return(nil)

	err := service.DeleteUser(context.Background(), userID.String())

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_InvalidUserID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	logger := logrus.New()
	service := NewUsersServiceImpl(mockRepo, logger)

	err := service.DeleteUser(context.Background(), "invalid-uuid")

	assert.Error(t, err)
	mockRepo.AssertNotCalled(t, "DeleteUser")
}

func TestDeleteUser_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)
	logger := logrus.New()
	service := NewUsersServiceImpl(mockRepo, logger)

	userID := uuid.New()

	mockRepo.On("DeleteUser", mock.Anything, userID).Return(errors.New("delete error"))

	err := service.DeleteUser(context.Background(), userID.String())

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
