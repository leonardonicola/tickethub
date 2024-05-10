package route

import (
	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/handler"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/repository"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/usecase"
	ticket "github.com/leonardonicola/tickethub/internal/modules/ticket/repository"
	ticketUc "github.com/leonardonicola/tickethub/internal/modules/ticket/usecase"
	user "github.com/leonardonicola/tickethub/internal/modules/user/repository"
	userUc "github.com/leonardonicola/tickethub/internal/modules/user/usecase"
	gateway "github.com/leonardonicola/tickethub/internal/pkg/stripe"
)

// TODO: setup das rotas do purchase e DI
func SetupPurchaseRoutes() *handler.PurchaseHandler {
	purchaseRepo := repository.NewPurchaseRepository(config.GetDB())
	ticketRepo := ticket.NewTicketRepository(config.GetDB())
	userRepo := user.NewUserRepository(config.GetDB())

	stripeGateway := gateway.GetStripeGateway()
	updateAvlQtyUc := &ticketUc.UpdateAvailableQuantityUseCase{
		TicketRepository: ticketRepo,
	}
	getTicketProductUc := &ticketUc.GetTicketProductUseCase{
		TicketRepository: ticketRepo,
	}
	getTicketByIdUc := &ticketUc.GetTicketByIdUseCase{
		TicketRepository: ticketRepo,
	}
	getUserByIdUc := &userUc.GetUserByIdUseCase{
		UserRepository: userRepo,
	}
	createUc := &usecase.PurchaseCreateUseCase{
		PurchaseRepository:             purchaseRepo,
		PaymentGateway:                 stripeGateway,
		UpdateAvailableQuantityUseCase: updateAvlQtyUc,
		GetTicketProductUseCase:        getTicketProductUc,
		GetTicketByIdUseCase:           getTicketByIdUc,
		GetUserByIdUseCase:             getUserByIdUc,
	}
	failedUc := &usecase.PurchaseFailedUseCase{
		PurchaseRepository:             purchaseRepo,
		UpdateAvailableQuantityUseCase: updateAvlQtyUc,
	}

	purchaseHandler := handler.NewPurchaseHandler(failedUc,
		createUc, getTicketByIdUc, getUserByIdUc)
	return purchaseHandler
}
