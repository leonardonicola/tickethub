package usecase

import (
	"fmt"
	"sync"

	"github.com/leonardonicola/tickethub/internal/modules/purchase/enum"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/ports"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/dto"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/usecase"
	"github.com/stripe/stripe-go/v78"
)

type PurchaseFailedUseCase struct {
	PurchaseRepository             ports.PurchaseRepository
	UpdateAvailableQuantityUseCase *usecase.UpdateAvailableQuantityUseCase
}

// Devolve available_qty para o ticket e muda status do purchase para failed
func (uc *PurchaseFailedUseCase) Execute(session stripe.CheckoutSession) error {

	purchaseId, ok := session.Metadata["purchase_id"]
	if !ok {
		return fmt.Errorf("No purchase_id found for Stripe order: %s", session.ID)
	}

	err := uc.PurchaseRepository.UpdatePaymentStatus(enum.PaymentFailed, purchaseId)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errors := make(chan error, len(session.LineItems.Data))

	for _, val := range session.LineItems.Data {
		wg.Add(1)
		go func(item *stripe.LineItem) {
			defer wg.Done()
			payload := &dto.UpdateTicketAvailableQtyInputDTO{
				ID:       item.ID,
				Quantity: item.Quantity,
				Type:     "increment",
			}
			// Devolve os tickets reservados temporariamente
			err := uc.UpdateAvailableQuantityUseCase.Execute(payload)
			if err != nil {
				errors <- err
				return
			}
		}(val)
	}
	wg.Wait()
	close(errors)

	for err := range errors {
		return err
	}

	return nil

}
