package dto

type CreateUserDTO struct {
	Name    string `json:"name" validate:"required,gte=1,lte=200"`
	Surname string `json:"surname" validate:"required,gte=1,lte=200"`
	Email   string `json:"email" validate:"required,email"`
	Address string `json:"address" validate:"required,gte=1,lte=254"`
	CPF     string `json:"cpf" validate:"required,gte=11,lte=11"`
}
