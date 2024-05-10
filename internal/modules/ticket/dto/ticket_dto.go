package dto

type CreateTicketInputDTO struct {
	Name        string `json:"name" validate:"required,gt=2,lt=100"`
	Price       int64  `json:"price" validate:"required,numeric"`
	TotalQty    uint32 `json:"total_qty" validate:"required,numeric"`
	Description string `json:"description" validate:"required,gt=3,lt=250"`
	MaxPerUser  uint16 `json:"max_per_user" validate:"required,number"`
	EventId     string `json:"event_id" validate:"required,uuid"`
}

type CreateTicketOutputDTO struct {
	ID           string
	Name         string
	Price        int64
	TotalQty     uint32
	AvailableQty uint32
	Description  string
	MaxPerUser   uint16
	EventId      string
}

type GetTicketByIdDTO struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	TotalQty     uint32 `json:"total_qty"`
	Description  string `json:"description"`
	MaxPerUser   uint16 `json:"max_per_user"`
	AvailableQty uint32 `json:"available_qty"`
	EventId      string `json:"event_id"`
}

type UpdateTicketAvailableQtyInputDTO struct {
	ID       string
	Quantity int64
	Type     string
}
