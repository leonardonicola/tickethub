package usecase

import (
	"github.com/leonardonicola/tickethub/internal/modules/genre/domain"
	"github.com/leonardonicola/tickethub/internal/modules/genre/port"
)

type GetGenreByIdUseCase struct {
	GenreRepository port.GenreRepository
}

func (uc *GetGenreByIdUseCase) Execute(id string) (*domain.Genre, error) {
	return uc.GenreRepository.GetById(id)
}
