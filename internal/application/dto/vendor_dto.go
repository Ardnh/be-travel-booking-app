package dto

type CreateVendorDTO struct {
	BusinessName        string  `json:"business_name" validate:"required"`
	OwnerUserID         string  `json:"owner_user_id" validate:"required"`
	OwnerName           string  `json:"owner_name" validate:"required"`
	Description         string  `json:"description"`
	FoundedYear         int     `json:"founded_year"`
	PhoneNumber         string  `json:"phone_number"`
	Email               string  `json:"email" validate:"required,email"`
	HeadOfficeAddress   string  `json:"head_office_address"`
	LegalDocumentNumber *string `json:"legal_document_number,omitempty"`
	Status              string  `json:"status"`
}

type UpdateVendorDTO struct {
	BusinessName        *string `json:"business_name,omitempty"`
	OwnerName           *string `json:"owner_name,omitempty"`
	Description         *string `json:"description,omitempty"`
	FoundedYear         *int    `json:"founded_year,omitempty"`
	PhoneNumber         *string `json:"phone_number,omitempty"`
	Email               *string `json:"email,omitempty"`
	HeadOfficeAddress   *string `json:"head_office_address,omitempty"`
	LogoURL             *string `json:"logo_url,omitempty"`
	BannerURL           *string `json:"banner_url,omitempty"`
	LegalDocumentNumber *string `json:"legal_document_number,omitempty"`
	IsVerified          *bool   `json:"is_verified,omitempty"`
	Status              *string `json:"status,omitempty"`
}
