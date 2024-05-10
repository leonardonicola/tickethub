package dto

type CreatePurchaseDTO struct {
	TicketID      string `json:"ticket_id" validate:"required,uuid"`
	UserID        string `json:"user_id" validate:"required,uuid"`
	Quantity      int64  `json:"quantity" validate:"required,number"`
	PaymentMethod string `json:"payment_method" validate:"required"`
}
