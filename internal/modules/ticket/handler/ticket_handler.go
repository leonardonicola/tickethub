package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/dto"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/usecase"
	"github.com/leonardonicola/tickethub/internal/pkg/validation"
)

var (
	logger *config.Logger
)

type TicketHandler struct {
	CreateTicketUseCase *usecase.CreateTicketUseCase
}

func NewTicketHandler(createUc usecase.CreateTicketUseCase) *TicketHandler {
	logger = config.NewLogger()
	return &TicketHandler{
		CreateTicketUseCase: &createUc,
	}
}

// @Summary Create ticket
// @Description	Create an ticket
// @Tags ticket
// @Accept json
// @Produce json
// @SecurityDefinitions.apikey JWT
// @Success 200 {object} dto.CreateTicketInputDTO
// @Router /ticket [post]
func (h *TicketHandler) CreateHandler(ctx *gin.Context) {
	var payload dto.CreateTicketInputDTO

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Corpo inválido",
		})
		return
	}

	errs := validation.Validate(payload)

	if errs != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, errs)
		return
	}

	ticket, err := h.CreateTicketUseCase.Execute(&payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	logger.Info("Ticket criado com sucesso!")
	ctx.JSON(http.StatusCreated, ticket)
}
