package dto

type CreateServiceTypeDTO struct {
	Name               string `json:"name" validate:"required"`
	UniqueCode         string `json:"unique_code" validate:"required"`
	Description        string `json:"description"`
	NeedChair          bool   `json:"need_chair"`
	NeedPickupAddress  bool   `json:"need_pickup_address"`
	NeedDropoffAddress bool   `json:"need_dropoff_address"`
	DisplayOrder       int    `json:"display_order"`
	Status             bool   `json:"status"`
}

type UpdateServiceTypeDTO struct {
	Name               *string `json:"name,omitempty"`
	UniqueCode         *string `json:"unique_code,omitempty"`
	Description        *string `json:"description,omitempty"`
	NeedChair          *bool   `json:"need_chair,omitempty"`
	NeedPickupAddress  *bool   `json:"need_pickup_address,omitempty"`
	NeedDropoffAddress *bool   `json:"need_dropoff_address,omitempty"`
	DisplayOrder       *int    `json:"display_order,omitempty"`
	Status             *bool   `json:"status,omitempty"`
}