package usecase

import (
	"github.com/google/uuid"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/domain"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/dto"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/enum"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/ports"
	ticketDTO "github.com/leonardonicola/tickethub/internal/modules/ticket/dto"
	ticketPort "github.com/leonardonicola/tickethub/internal/modules/ticket/ports"
	userPort "github.com/leonardonicola/tickethub/internal/modules/user/ports"
)

type PurchaseCreateUseCase struct {
	PurchaseRepository ports.PurchaseRepository
	TicketRepository   ticketPort.TicketRepository
	UserRepository     userPort.UserRepository
	PaymentGateway     ports.PaymentGateway
}

// TODO: criar purchase, reduzir available qty do ticket
func (uc *PurchaseCreateUseCase) Execute(dto *dto.CreatePurchaseDTO) (string, error) {

	if _, err := uc.UserRepository.GetById(dto.UserID); err != nil {
		return "", err
	}

	if _, err := uc.TicketRepository.GetTicketById(dto.TicketID); err != nil {
		return "", err
	}

	id := uuid.NewString()
	purchase, err := domain.NewPurchase(id, dto.TicketID,
		dto.UserID, enum.PaymentProcessing,
		dto.PaymentMethod, dto.Quantity)

	if err != nil {
		return "", err
	}

	if err := uc.PurchaseRepository.Create(purchase); err != nil {
		return "", err
	}

	payload := &ticketDTO.UpdateTicketAvailableQtyInputDTO{
		ID:       dto.TicketID,
		Quantity: dto.Quantity,
		Type:     "decrement",
	}
	// Reserva quantidades temporariamente
	err = uc.TicketRepository.UpdateAvailableQuantity(payload)
	if err != nil {
		return "", err
	}

	ticketProduct, err := uc.TicketRepository.GetProductByTicketId(dto.TicketID)

	if err != nil {
		return "", err
	}

	tickets := []*ticketDTO.CreatePaymentDTO{
		&ticketDTO.CreatePaymentDTO{
			StripeID: ticketProduct.StripeID,
			PriceID:  ticketProduct.PriceID,
			Quantity: dto.Quantity,
		},
	}
	metadata := map[string]string{
		"purchase_id": purchase.ID,
	}

	link, err := uc.PaymentGateway.CreatePaymentLink(tickets, metadata)

	if err != nil {
		return "", err
	}

	return link, nil
}
