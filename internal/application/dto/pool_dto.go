package dto

import "time"

type CreatePoolsDTO struct {
	VendorID    string  `json:"vendor_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Slug        string  `json:"slug" validate:"required"`
	Address     string  `json:"address" validate:"required"`
	City        string  `json:"city" validate:"required"`
	Province    string  `json:"province" validate:"required"`
	Latitude    float64 `json:"latitude" validate:"required"`
	Longitude   float64 `json:"longitude" validate:"required"`
	OpenTime    string  `json:"open_time" validate:"required"`
	CloseTime   string  `json:"close_time" validate:"required"`
	Status      string  `json:"status"`
	Description *string `json:"description,omitempty"`
}

type UpdatePoolsDTO struct {
	VendorID    *string  `json:"vendor_id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Slug        *string  `json:"slug,omitempty"`
	Address     *string  `json:"address,omitempty"`
	City        *string  `json:"city,omitempty"`
	Province    *string  `json:"province,omitempty"`
	Latitude    *float64 `json:"latitude,omitempty"`
	Longitude   *float64 `json:"longitude,omitempty"`
	OpenTime    *string  `json:"open_time,omitempty"`
	CloseTime   *string  `json:"close_time,omitempty"`
	Status      *string  `json:"status,omitempty"`
	Description *string  `json:"description,omitempty"`
}

type PoolsResponseDTO struct {
	PoolID      string    `json:"pool_id"`
	VendorID    string    `json:"vendor_id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	OpenTime    string    `json:"open_time"`
	CloseTime   string    `json:"close_time"`
	Status      string    `json:"status"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
