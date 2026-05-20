package services

import (
	"context"

	"github.com/ardnh/be-travel-booking-app/internal/application/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequestDto) (*dto.LoginResponseDto, error)
	Register(ctx context.Context, req dto.RegisterRequestDto) (*dto.RegisterResponseDto, error)
}
