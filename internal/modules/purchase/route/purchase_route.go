package route

import (
	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/handler"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/repository"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/usecase"
	ticket "github.com/leonardonicola/tickethub/internal/modules/ticket/repository"
	user "github.com/leonardonicola/tickethub/internal/modules/user/repository"
	gateway "github.com/leonardonicola/tickethub/internal/pkg/stripe"
)

// TODO: setup das rotas do purchase e DI
func SetupPurchaseRoutes() *handler.PurchaseHandler {
	purchaseRepo := repository.NewPurchaseRepository(config.GetDB())
	ticketRepo := ticket.NewTicketRepository(config.GetDB())
	userRepo := user.NewUserRepository(config.GetDB())

	stripeGateway := gateway.GetStripeGateway()
	createUc := &usecase.PurchaseCreateUseCase{
		PurchaseRepository: purchaseRepo,
		PaymentGateway:     stripeGateway,
		TicketRepository:   ticketRepo,
		UserRepository:     userRepo,
	}
	failedUc := &usecase.PurchaseFailedUseCase{
		PurchaseRepository: purchaseRepo,
		TicketRepository:   ticketRepo,
	}

	purchaseHandler := handler.NewPurchaseHandler(failedUc,
		createUc)
	return purchaseHandler
}
