package dto

type CreateTicketInputDTO struct {
	ID          string
	Name        string  `json:"name" validate:"required,gt=2,lt=100"`
	Price       float32 `json:"price" validate:"required,numeric"`
	TotalQty    uint32  `json:"total_qty" validate:"required,numeric"`
	Description string  `json:"description" validate:"required,gt=3,lt=250"`
	MaxPerUser  uint16  `json:"max_per_user" validate:"required,number"`
	EventId     string  `json:"event_id" validate:"required,uuid"`
}

type CreateTicketOutputDTO struct {
	ID           string
	Name         string
	Price        float32
	TotalQty     uint32
	AvailableQty uint32
	Description  string
	MaxPerUser   uint16
	EventId      string
}
