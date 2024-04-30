package route

import (
	"github.com/leonardonicola/tickethub/config"
	eventRepo "github.com/leonardonicola/tickethub/internal/modules/event/repository"
	eventUcs "github.com/leonardonicola/tickethub/internal/modules/event/usecase"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/handler"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/repository"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/usecase"
)

func SetupTicketRoutes() *handler.TicketHandler {
	ticketRepo := repository.NewTicketRepository(config.GetDB())
	eventRepoImpl := eventRepo.NewEventRepository(config.GetDB())

	getEventById := &eventUcs.GetEventByIdUseCase{
		EventRepository: eventRepoImpl,
	}
	createUc := &usecase.CreateTicketUseCase{
		TicketRepository:    ticketRepo,
		GetEventByIdUseCase: *getEventById,
	}
	ticketHandler := handler.NewTicketHandler(*createUc)
	return ticketHandler
}
