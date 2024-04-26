package usecase

import (
	"github.com/leonardonicola/tickethub/internal/modules/event/domain"
	"github.com/leonardonicola/tickethub/internal/modules/event/dto"
	"github.com/leonardonicola/tickethub/internal/modules/event/ports"
)

type CreateEventUseCase struct {
	EventRepository ports.EventRepository
}

func (uc *CreateEventUseCase) Execute(payload *dto.CreateEventInputDTO) (*dto.CreateEventOutputDTO, error) {
	event, err := domain.NewEvent(
		payload.Title, payload.Description, payload.Address, payload.Date,
		payload.Genre, payload.AgeRating)
	if err != nil {
		return nil, err
	}
	createdEvent, err := uc.EventRepository.Create(event, payload.Poster)
	if err != nil {
		return nil, err
	}
	return createdEvent, nil

}
