package usecase

import (
	"fmt"

	"github.com/leonardonicola/tickethub/internal/modules/event/dto"
	"github.com/leonardonicola/tickethub/internal/modules/event/ports"
)

type GetManyEventsUseCase struct {
	EventRepository ports.EventRepository
}

func (uc *GetManyEventsUseCase) Execute(query dto.GetManyEventsInputDTO) ([]dto.GetManyEventsOutputDTO, error) {

	events, err := uc.EventRepository.GetMany(query)
	if err != nil {
		return nil, fmt.Errorf("EVENT(get_many): %v", err)
	}
	return events, nil
}
