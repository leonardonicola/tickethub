package usecase

import (
	"github.com/leonardonicola/tickethub/internal/modules/event/domain"
	"github.com/leonardonicola/tickethub/internal/modules/event/dto"
	"github.com/leonardonicola/tickethub/internal/modules/event/ports"
	"github.com/leonardonicola/tickethub/internal/modules/genre/usecase"
)

type CreateEventUseCase struct {
	EventRepository     ports.EventRepository
	GetGenreByIdUseCase usecase.GetGenreByIdUseCase
}

// Verificar se genre_id pertence a algum record antes de criar evento
func (uc *CreateEventUseCase) Execute(payload *dto.CreateEventInputDTO) (*dto.CreateEventOutputDTO, error) {
	_, err := uc.GetGenreByIdUseCase.Execute(payload.GenreID)
	if err != nil {
		return nil, err
	}
	event, err := domain.NewEvent(
		payload.Title, payload.Description, payload.Address, payload.Date,
		payload.GenreID, payload.AgeRating)
	if err != nil {
		return nil, err
	}
	createdEvent, err := uc.EventRepository.Create(event, payload.Poster)
	if err != nil {
		return nil, err
	}
	return createdEvent, nil

}
