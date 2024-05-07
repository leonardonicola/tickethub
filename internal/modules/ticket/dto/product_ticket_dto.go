package dto

type TicketProduct struct {
	PriceID     string
	Description string
	TicketID    string
	StripeID    string
	Name        string
	Metadata    any
}

type CreatePaymentDTO struct {
	StripeID string
	PriceID  string
	Quantity int64
}
