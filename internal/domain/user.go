package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID      string
	Name    string
	Surname string
	Email   string
	CPF     string
	Address string
}

func NewUser(name, surname, email, address, cpf string) *User {
	id := uuid.NewString()
	return &User{
		ID:      id,
		Name:    name,
		Surname: surname,
		Email:   email,
		CPF:     cpf,
		Address: address,
	}
}
