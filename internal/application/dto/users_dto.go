package dto

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type UpdateUserDTO struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
