package dto

type LoginRequestDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=1,max=30"`
}

type LoginResponseDto struct {
	Token       string   `json:"token"`
	ExpireDate  string   `json:"expire_date"`
	Permissions []string `json:"permissions"`
}
