package port

import "github.com/leonardonicola/tickethub/internal/modules/genre/domain"

type GenreRepository interface {
	GetById(id string) (*domain.Genre, error)
}
