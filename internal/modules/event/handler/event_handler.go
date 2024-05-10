package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/event/dto"
	"github.com/leonardonicola/tickethub/internal/modules/event/usecase"
	"github.com/leonardonicola/tickethub/internal/pkg/validation"
)

var (
	logger *config.Logger
)

type EventHandler struct {
	CreateEventUseCase   *usecase.CreateEventUseCase
	GetManyEventsUseCase *usecase.GetManyEventsUseCase
}

func NewEventHandler(createUc *usecase.CreateEventUseCase, getManyUc *usecase.GetManyEventsUseCase) *EventHandler {
	logger = config.NewLogger()
	return &EventHandler{
		CreateEventUseCase:   createUc,
		GetManyEventsUseCase: getManyUc,
	}
}

// @Summary Create Event
// @Description Create an show/event/festival/concert
// @Tags event
// @Accept json
// @Produce json
// @SecurityDefinitions.apikey JWT
// @Success 200 {object} dto.CreateEventOutputDTO
// @Router /event [post]
func (uc *EventHandler) CreateEventHandler(ctx *gin.Context) {
	var payload *dto.CreateEventInputDTO
	if err := ctx.Bind(&payload); err != nil {
		logger.Errorf("EVENT(create) handler: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	errs := validation.Validate(payload)

	if errs != nil {
		logger.Errorf("EVENT(create) handler: %v", errs)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, errs)
		return
	}
	event, err := uc.CreateEventUseCase.Execute(payload)
	if err != nil {
		logger.Errorf("EVENT(create) use case: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	logger.Info("event created successfully")
	ctx.JSON(http.StatusCreated, event)
}

// @Summary Get many events
// @Description Get events base on title search
// @Tags event
// @Param search query string true "Search"
// @Param limit query uint8 true "Limit"
// @Param page query uint8 true "Page"
// @Produce json
// @SecurityDefinitions.apikey JWT
// @Success 200 {array} dto.GetManyEventsOutputDTO
// @Failure 406
// @Failure 500
// @Router /event [get]
func (uc *EventHandler) GetManyHandler(ctx *gin.Context) {
	var query dto.GetManyEventsInputDTO
	if err := ctx.BindQuery(&query); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	errs := validation.Validate(query)

	if errs != nil {
		logger.Errorf("EVENT(create) handler: %v", errs)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, errs)
		return
	}

	events, err := uc.GetManyEventsUseCase.Execute(query)

	if err != nil {
		logger.Errorf("get many events error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, events)
}
