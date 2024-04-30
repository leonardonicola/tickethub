package usecase

import (
	"github.com/google/uuid"
	"github.com/leonardonicola/tickethub/internal/modules/event/usecase"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/domain"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/dto"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/ports"
)

type CreateTicketUseCase struct {
	TicketRepository    ports.TicketRepository
	GetEventByIdUseCase usecase.GetEventByIdUseCase
}

func (uc *CreateTicketUseCase) Execute(payload *dto.CreateTicketInputDTO) (*dto.CreateTicketOutputDTO, error) {

	if _, err := uc.GetEventByIdUseCase.Execute(payload.EventId); err != nil {
		return nil, err
	}

	id := uuid.NewString()
	ticket := domain.NewTicket(id, payload.Name, payload.Description,
		payload.EventId, payload.Price, payload.MaxPerUser, payload.TotalQty)

	return uc.TicketRepository.Create(ticket)

}
