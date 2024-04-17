package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardonicola/tickethub/internal/dto"
)

func CreatePlayerHandler(ctx *gin.Context) {
	request := dto.CreateUserDTO{}
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Errorf("Erro ao criar usuário: %v", err)
		return
	}
	ctx.JSON(200, request)
}
