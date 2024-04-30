package usecase

import (
	"github.com/leonardonicola/tickethub/internal/modules/event/dto"
	"github.com/leonardonicola/tickethub/internal/modules/event/ports"
)

type GetEventByIdUseCase struct {
	EventRepository ports.EventRepository
}

func (uc *GetEventByIdUseCase) Execute(id string) (*dto.GetEventByIdOutputDTO, error) {
	return uc.EventRepository.GetById(id)
}
