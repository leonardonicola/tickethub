package usecase

import (
	"github.com/google/uuid"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/domain"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/dto"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/enum"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/ports"
	ticket "github.com/leonardonicola/tickethub/internal/modules/ticket/dto"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/usecase"
	userUc "github.com/leonardonicola/tickethub/internal/modules/user/usecase"
)

type PurchaseCreateUseCase struct {
	PurchaseRepository             ports.PurchaseRepository
	UpdateAvailableQuantityUseCase *usecase.UpdateAvailableQuantityUseCase
	GetTicketByIdUseCase           *usecase.GetTicketByIdUseCase
	GetUserByIdUseCase             *userUc.GetUserByIdUseCase
	GetTicketProductUseCase        *usecase.GetTicketProductUseCase
	PaymentGateway                 ports.PaymentGateway
}

// TODO: criar purchase, reduzir available qty do ticket
func (uc *PurchaseCreateUseCase) Execute(dto *dto.CreatePurchaseDTO) (string, error) {

	if _, err := uc.GetUserByIdUseCase.Execute(dto.UserID); err != nil {
		return "", err
	}

	if _, err := uc.GetTicketByIdUseCase.Execute(dto.TicketID); err != nil {
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

	payload := &ticket.UpdateTicketAvailableQtyInputDTO{
		ID:       dto.TicketID,
		Quantity: dto.Quantity,
		Type:     "decrement",
	}
	// Reserva quantidades temporariamente
	err = uc.UpdateAvailableQuantityUseCase.Execute(payload)
	if err != nil {
		return "", err
	}

	ticketProduct, err := uc.GetTicketProductUseCase.Execute(dto.TicketID)

	if err != nil {
		return "", err
	}

	tickets := []*ticket.CreatePaymentDTO{
		&ticket.CreatePaymentDTO{
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
