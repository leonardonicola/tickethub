package usecase

import (
	"github.com/leonardonicola/tickethub/internal/modules/purchase/ports"
	"github.com/stripe/stripe-go/v78"
)

type PurchaseSucceededUseCase struct {
	PurchaseRepository ports.PurchaseRepository
}

// TODO: implementar envio de e-mail pro user, atualizar payment status
func (uc *PurchaseSucceededUseCase) Execute(session stripe.CheckoutSession) error {
	return nil
}
