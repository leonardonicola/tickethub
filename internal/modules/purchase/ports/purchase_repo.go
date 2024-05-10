package ports

import (
	"github.com/leonardonicola/tickethub/internal/modules/purchase/domain"
)

type PurchaseRepository interface {
	UpdatePaymentStatus(status string, id string) error
	Create(*domain.Purchase) error
}
