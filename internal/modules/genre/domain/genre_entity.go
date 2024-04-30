package domain

import "github.com/google/uuid"

type Genre struct {
	ID   string
	Name string
}

func NewGenre(name string) *Genre {
	id := uuid.NewString()

	return &Genre{
		ID:   id,
		Name: name,
	}
}
