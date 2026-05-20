package dto

type LoginRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=1,max=30"`
}

type LoginResponseDto struct {
	Token      string `json:"token"`
	ExpireDate string `json:"expire_date"`
}

type RegisterRequestDto struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=30"`
	Phone    string `json:"phone"`
}

type RegisterResponseDto struct {
	Message string `json:"message"`
}
