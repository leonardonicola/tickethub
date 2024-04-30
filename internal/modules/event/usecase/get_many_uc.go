package usecase

import (
	"github.com/leonardonicola/tickethub/internal/modules/event/dto"
	"github.com/leonardonicola/tickethub/internal/modules/event/ports"
)

type GetManyEventsUseCase struct {
	EventRepository ports.EventRepository
}

func (uc *GetManyEventsUseCase) Execute(query dto.GetManyEventsInputDTO) ([]dto.GetManyEventsOutputDTO, error) {

	return uc.EventRepository.GetMany(query)
}
