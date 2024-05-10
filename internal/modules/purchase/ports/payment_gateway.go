package ports

import (
	ticketDTO "github.com/leonardonicola/tickethub/internal/modules/ticket/dto"
)

type PaymentGateway interface {
	CreatePaymentLink([]*ticketDTO.CreatePaymentDTO, map[string]string) (string, error)
}
