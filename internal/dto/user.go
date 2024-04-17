package dto

type CreateUserInputDTO struct {
	Name    string `json:"name" validate:"required,gte=1,lte=200"`
	Surname string `json:"surname" validate:"required,gte=1,lte=200"`
	Email   string `json:"email" validate:"required,email"`
	Address string `json:"address" validate:"required,gte=1,lte=254"`
	CPF     string `json:"cpf" validate:"required,gte=11,lte=11"`
}

type CreateUserOutputDTO struct {
	ID      string
	Name    string
	Surname string
	Email   string
	Address string
	CPF     string
}

type Login struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
