package domain

import (
	"errors"

	"github.com/leonardonicola/tickethub/internal/modules/purchase/enum"
)

type Purchase struct {
	ID            string
	TicketID      string
	UserID        string
	Quantity      int64
	Status        string
	PaymentMethod string
}

func NewPurchase(id, ticketId, userId, status,
	paymentMethod string, quantity int64) (*Purchase, error) {
	possibleStatus := map[string]struct{}{
		enum.PaymentCanceled:   struct{}{},
		enum.PaymentFailed:     struct{}{},
		enum.PaymentProcessing: struct{}{},
		enum.PaymentSucceeded:  struct{}{},
	}

	_, ok := possibleStatus[status]

	if !ok {
		return nil, errors.New("Unknown purchase status!")
	}

	return &Purchase{
		ID:            id,
		TicketID:      ticketId,
		UserID:        userId,
		Quantity:      quantity,
		Status:        status,
		PaymentMethod: paymentMethod,
	}, nil
}
