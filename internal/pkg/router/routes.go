package router

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/leonardonicola/tickethub/config"
	eventRoutes "github.com/leonardonicola/tickethub/internal/modules/event/route"
	purchaseRoutes "github.com/leonardonicola/tickethub/internal/modules/purchase/route"
	ticketRoutes "github.com/leonardonicola/tickethub/internal/modules/ticket/route"
	userRoutes "github.com/leonardonicola/tickethub/internal/modules/user/route"
	"github.com/leonardonicola/tickethub/internal/pkg/validation"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	logger *config.Logger
)

func InitRoutes() (*gin.Engine, error) {
	logger = config.NewLogger()
	router := gin.Default()
	const PREFIX = "/api/v1"
	// 8 MB - left shift operator - 8 x (2 elevado a 20)
	router.MaxMultipartMemory = 8 << 20

	authMiddleware, err := InitAuthMiddleware()

	if err != nil {
		logger.Errorf("authMiddleware error: %v", err)
		return nil, fmt.Errorf("authMiddleware error: %v", err)
	}

	if err := validation.InitValidator(); err != nil {
		logger.Errorf("validation error: %v", err)
		return nil, fmt.Errorf("validation error: %v", err)
	}

	// Handles not found routes
	router.NoRoute(authMiddleware.MiddlewareFunc(), func(ctx *gin.Context) {
		claims := jwt.ExtractClaims(ctx)
		log.Printf("NoRoute Claims: %v\n", claims)
		ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	userHandlers := userRoutes.SetupUserRoutes()
	eventHandlers := eventRoutes.SetupEventRoutes()
	ticketHandlers := ticketRoutes.SetupTicketRoutes()
	purchaseHandlers := purchaseRoutes.SetupPurchaseRoutes()

	v1 := router.Group(PREFIX)
	v1.POST("/login", authMiddleware.LoginHandler)
	v1.POST("/register", userHandlers.RegisterHandler)
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1.POST("/webhook", purchaseHandlers.StripeWebhook)
	// With auth middleware
	v1.Use(authMiddleware.MiddlewareFunc())
	{
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)
		v1.POST("/logout", authMiddleware.LogoutHandler)

		// Event
		v1.POST("/event", eventHandlers.CreateEventHandler)
		v1.GET("/event", eventHandlers.GetManyHandler)

		// Ticket
		v1.POST("/ticket", ticketHandlers.CreateHandler)

		// Purchase
		v1.POST("/purchase", purchaseHandlers.Create)
	}

	return router, nil
}
