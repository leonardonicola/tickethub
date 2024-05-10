package route

import (
	"github.com/leonardonicola/tickethub/config"
	eventRepo "github.com/leonardonicola/tickethub/internal/modules/event/repository"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/handler"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/repository"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/usecase"
	gateway "github.com/leonardonicola/tickethub/internal/pkg/stripe"
)

func SetupTicketRoutes() *handler.TicketHandler {
	ticketRepo := repository.NewTicketRepository(config.GetDB())
	eventRepoImpl := eventRepo.NewEventRepository(config.GetDB())

	stripeGateway := gateway.GetStripeGateway()
	createUc := &usecase.CreateTicketUseCase{
		TicketRepository: ticketRepo,
		EventRepository:  eventRepoImpl,
		PaymentGateway:   stripeGateway,
	}
	updateUc := &usecase.UpdateAvailableQuantityUseCase{
		TicketRepository: ticketRepo,
	}
	ticketHandler := handler.NewTicketHandler(createUc, updateUc)
	return ticketHandler
}
