package route

import (
	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/event/handler"
	"github.com/leonardonicola/tickethub/internal/modules/event/repository"
	"github.com/leonardonicola/tickethub/internal/modules/event/usecase"
)

func SetupEventRoutes() *handler.EventHandler {
	pgRepo := repository.NewEventRepository(config.GetDB())
	createUc := &usecase.CreateEventUseCase{
		EventRepository: pgRepo,
	}
	getManyUc := &usecase.GetManyEventsUseCase{
		EventRepository: pgRepo,
	}
	eventHandler := handler.NewEventHandler(createUc, getManyUc)
	return eventHandler
}
