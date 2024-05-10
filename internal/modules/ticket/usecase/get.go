package usecase

import (
	"github.com/leonardonicola/tickethub/internal/modules/ticket/dto"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/ports"
)

type GetTicketProductUseCase struct {
	TicketRepository ports.TicketRepository
}

type GetTicketUseCase struct {
	TicketRepository ports.TicketRepository
}

func (uc *GetTicketProductUseCase) Execute(id string) (*dto.TicketProduct, error) {
	return uc.TicketRepository.GetProductByTicketId(id)
}
func (uc *GetTicketUseCase) Execute(id string) (*dto.GetTicketDTO, error) {
	return uc.TicketRepository.GetTicketById(id)
}
