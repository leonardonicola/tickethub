package ports

import (
	"github.com/leonardonicola/tickethub/internal/modules/ticket/domain"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/dto"
)

type TicketRepository interface {
	Create(*domain.Ticket) (*dto.CreateTicketOutputDTO, error)
	CreateTicketProduct(*dto.TicketProduct) error
}
