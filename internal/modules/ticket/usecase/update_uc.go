package usecase

import (
	"github.com/leonardonicola/tickethub/internal/modules/ticket/dto"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/ports"
)

type UpdateAvailableQuantityUseCase struct {
	TicketRepository ports.TicketRepository
}

func (uc *UpdateAvailableQuantityUseCase) Execute(payload *dto.UpdateTicketAvailableQtyInputDTO) error {
	return uc.TicketRepository.UpdateAvailableQuantity(payload)
}
