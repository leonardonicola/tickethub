package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/dto"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/usecase"
	ticketUc "github.com/leonardonicola/tickethub/internal/modules/ticket/usecase"
	userUc "github.com/leonardonicola/tickethub/internal/modules/user/usecase"
	"github.com/leonardonicola/tickethub/internal/pkg/validation"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/webhook"
)

type PurchaseHandler struct {
	PurchaseFailedUseCase *usecase.PurchaseFailedUseCase
	// PurchaseSucceededUseCase *usecase.PurchaseSucceededUseCase
	PurchaseCreateUseCase *usecase.PurchaseCreateUseCase
	GetTicketByIdUseCase  *ticketUc.GetTicketByIdUseCase
	GetUserByIdUseCase    *userUc.GetUserByIdUseCase
}

var (
	stripeWebhookKey string
	logger           *config.Logger
)

func NewPurchaseHandler(purchaseFailedUc *usecase.PurchaseFailedUseCase,
	// purchaseSucceededUc *usecase.PurchaseSucceededUseCase,
	purchaseCreateUc *usecase.PurchaseCreateUseCase,
	getTicketByIdUc *ticketUc.GetTicketByIdUseCase,
	getUserByIdUc *userUc.GetUserByIdUseCase) *PurchaseHandler {

	stripeWebhookKey = os.Getenv("STRIPE_WEBHOOK_KEY")
	logger = config.NewLogger()
	return &PurchaseHandler{
		PurchaseFailedUseCase: purchaseFailedUc,
		// PurchaseSucceededUseCase: purchaseSucceededUc,
		PurchaseCreateUseCase: purchaseCreateUc,
		GetTicketByIdUseCase:  getTicketByIdUc,
		GetUserByIdUseCase:    getUserByIdUc,
	}
}

func (h *PurchaseHandler) Create(ctx *gin.Context) {

	var payload *dto.CreatePurchaseDTO
	if err := ctx.Bind(&payload); err != nil {
		logger.Errorf("PURCHASE(create) handler: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	errs := validation.Validate(payload)

	if errs != nil {
		logger.Errorf("PURCHASE(create) handler: %v", errs)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, errs)
		return
	}

	paymentLink, err := h.PurchaseCreateUseCase.Execute(payload)

	if err != nil {
		logger.Errorf("EVENT(create) usecase: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	logger.Info("Purchase create successfully")
	ctx.JSON(http.StatusCreated, paymentLink)
}

func (h *PurchaseHandler) StripeWebhook(ctx *gin.Context) {
	payload, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		logger.Errorf("Error on webhook: %v", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}
	event, err := webhook.ConstructEvent(payload, ctx.GetHeader("Stripe-Signature"), stripeWebhookKey)

	if err != nil {
		logger.Errorf("Error on webhook: %v", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var session stripe.CheckoutSession
	err = json.Unmarshal(event.Data.Raw, &session)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	switch event.Type {
	case stripe.EventTypeCheckoutSessionAsyncPaymentFailed:
		if err := h.PurchaseFailedUseCase.Execute(session); err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	ctx.Status(http.StatusOK)
}
