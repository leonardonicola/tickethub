package router

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/leonardonicola/tickethub/config"
	userHandler "github.com/leonardonicola/tickethub/internal/user/handler"
	userRepo "github.com/leonardonicola/tickethub/internal/user/repository"
	userUC "github.com/leonardonicola/tickethub/internal/user/usecase"
	"github.com/leonardonicola/tickethub/pkg/validation"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	logger *config.Logger
)

func InitRoutes() (*gin.Engine, error) {
	logger = config.NewLogger()
	router := gin.Default()
	prefix := "/api/v1"

	middleware, err := InitAuthMiddleware()

	if err != nil {
		logger.Errorf("authMiddleware error: %v", err)
		return nil, fmt.Errorf("authMiddleware error: %v", err)
	}

	if err := validation.InitValidator(); err != nil {
		logger.Errorf("validation error: %v", err)
		return nil, fmt.Errorf("validation error: %v", err)
	}

	// Handles not found routes
	router.NoRoute(middleware.MiddlewareFunc(), func(ctx *gin.Context) {
		claims := jwt.ExtractClaims(ctx)
		log.Printf("NoRoute Claims: %v\n", claims)
		ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	userRepo := userRepo.NewUserRepository(config.GetDB())
	userUc := userUC.RegisterUseCase{
		Repository: userRepo,
	}
	userHdlr := userHandler.NewUserHandler(userUc)

	v1 := router.Group(prefix)
	v1.POST("/login", middleware.LoginHandler)
	v1.POST("/register", userHdlr.RegisterHandler)
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// With auth middleware
	v1.Use(middleware.MiddlewareFunc())
	{
		v1.GET("/refresh_token", middleware.RefreshHandler)
		v1.POST("/logout", middleware.LogoutHandler)
	}

	return router, nil
}