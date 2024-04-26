package route

import (
	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/user/handler"
	"github.com/leonardonicola/tickethub/internal/modules/user/repository"
	"github.com/leonardonicola/tickethub/internal/modules/user/usecase"
)

func SetupUserRoutes() *handler.UserHandler {

	pgRepo := repository.NewUserRepository(config.GetDB())
	regUc := usecase.RegisterUseCase{
		Repository: pgRepo,
	}
	userHdlr := handler.NewUserHandler(regUc)
	return userHdlr
}
