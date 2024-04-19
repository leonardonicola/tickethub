package router

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/leonardonicola/tickethub/internal/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes() (*gin.Engine, error) {
	router := gin.Default()
	handler.InitHandler()
	prefix := "/api/v1"

	middleware, err := InitAuthMiddleware()

	if err != nil {
		return nil, fmt.Errorf("authMiddleware error: %v", err)
	}

	// Handles not found routes
	router.NoRoute(middleware.MiddlewareFunc(), func(ctx *gin.Context) {
		claims := jwt.ExtractClaims(ctx)
		log.Printf("NoRoute Claims: %v\n", claims)
		ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	v1 := router.Group(prefix)
	v1.POST("/login", middleware.LoginHandler)
	v1.POST("/register", handler.RegisterHandler)
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// With auth middleware
	v1.Use(middleware.MiddlewareFunc())
	{
		v1.GET("/refresh_token", middleware.RefreshHandler)
		v1.POST("/logout", middleware.LogoutHandler)
	}

	return router, nil
}
