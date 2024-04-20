package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/user/dto"
	"github.com/leonardonicola/tickethub/internal/user/usecase"
	"github.com/leonardonicola/tickethub/pkg/validation"
)

var (
	logger *config.Logger
)

type UserHandler struct {
	RegisterUseCase *usecase.RegisterUseCase
}

func NewUserHandler(regUc usecase.RegisterUseCase) *UserHandler {
	logger = config.NewLogger()
	return &UserHandler{
		RegisterUseCase: &regUc,
	}
}

// @Summary		Register
// @Description	Create an account
// @Tags			auth
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.CreateUserOutputDTO
// @Router			/register [post]
func (h *UserHandler) RegisterHandler(ctx *gin.Context) {
	payload := dto.CreateUserInputDTO{}
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	errs := validation.Validate(payload)
	if errs != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errs)
		return
	}
	user, err := h.RegisterUseCase.Execute(payload)
	if err != nil {
		logger.Errorf("register db error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	logger.Info("user created successfully")
	ctx.JSON(200, user)
}
