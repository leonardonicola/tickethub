package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID       string
	Name     string
	Surname  string
	Email    string
	CPF      string
	Password string
	Address  string
}

func NewUser(name, surname, email, address, cpf, password string) *User {
	id := uuid.NewString()
	return &User{
		ID:       id,
		Name:     name,
		Surname:  surname,
		Email:    email,
		CPF:      cpf,
		Address:  address,
		Password: password,
	}
}
