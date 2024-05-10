package usecase

import (
	"github.com/leonardonicola/tickethub/internal/modules/ticket/dto"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/ports"
)

type GetTicketProductUseCase struct {
	TicketRepository ports.TicketRepository
}

type GetTicketByIdUseCase struct {
	TicketRepository ports.TicketRepository
}

func (uc *GetTicketProductUseCase) Execute(id string) (*dto.TicketProduct, error) {
	return uc.TicketRepository.GetProductByTicketId(id)
}
func (uc *GetTicketByIdUseCase) Execute(id string) (*dto.GetTicketByIdDTO, error) {
	return uc.TicketRepository.GetTicketById(id)
}
