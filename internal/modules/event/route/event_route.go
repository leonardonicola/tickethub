package route

import (
	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/event/handler"
	"github.com/leonardonicola/tickethub/internal/modules/event/repository"
	"github.com/leonardonicola/tickethub/internal/modules/event/usecase"
	genreRepo "github.com/leonardonicola/tickethub/internal/modules/genre/repository"
	genreUc "github.com/leonardonicola/tickethub/internal/modules/genre/usecase"
)

func SetupEventRoutes() *handler.EventHandler {
	pgRepo := repository.NewEventRepository(config.GetDB())
	genreRepoImpl := genreRepo.NewGenreRepository(config.GetDB())
	getGenreUc := &genreUc.GetGenreByIdUseCase{
		GenreRepository: genreRepoImpl,
	}
	createUc := &usecase.CreateEventUseCase{
		EventRepository:     pgRepo,
		GetGenreByIdUseCase: *getGenreUc,
	}
	getManyUc := &usecase.GetManyEventsUseCase{
		EventRepository: pgRepo,
	}
	eventHandler := handler.NewEventHandler(createUc, getManyUc)
	return eventHandler
}
