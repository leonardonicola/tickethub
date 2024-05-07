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
	// Injecting stripe (external dependecy) on the usecase is not good practice
	PaymentGateway ports.TicketPaymentGateway[*dto.CreateTicketOutputDTO, *dto.TicketProduct]
}

func (uc *CreateTicketUseCase) Execute(payload *dto.CreateTicketInputDTO) (*dto.CreateTicketOutputDTO, error) {

	if _, err := uc.GetEventByIdUseCase.Execute(payload.EventId); err != nil {
		return nil, err
	}

	id := uuid.NewString()
	ticket := domain.NewTicket(id, payload.Name, payload.Description,
		payload.EventId, payload.Price, payload.MaxPerUser, payload.TotalQty)

	createdTicket, err := uc.TicketRepository.Create(ticket)

	if err != nil {
		return nil, err
	}

	stripeTicket, err := uc.PaymentGateway.CreateProduct(createdTicket)
	if err != nil {
		return nil, err
	}

	err = uc.TicketRepository.CreateTicketProduct(stripeTicket)

	if err != nil {
		return nil, err
	}

	return createdTicket, nil
}
