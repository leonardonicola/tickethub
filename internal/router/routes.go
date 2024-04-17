package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardonicola/tickethub/internal/handler"
)

func InitRoutes(router *gin.Engine) {
	handler.InitHandler()
	prefix := "/api/v1"

	v1 := router.Group(prefix)
	{
		v1.POST("/register", handler.CreatePlayerHandler)
	}
}
