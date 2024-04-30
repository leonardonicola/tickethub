package domain

type Ticket struct {
	ID          string
	Name        string
	Price       float32
	TotalQty    uint32
	Description string
	MaxPerUser  uint16
	EventId     string
}

func NewTicket(id, name, description, eventId string, price float32,
	maxPerUser uint16, totalQty uint32) *Ticket {

	return &Ticket{
		ID:          id,
		Name:        name,
		Description: description,
		MaxPerUser:  maxPerUser,
		Price:       price,
		TotalQty:    totalQty,
		EventId:     eventId,
	}
}
