package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/domain"
	"github.com/leonardonicola/tickethub/internal/modules/ticket/dto"
)

var (
	logger *config.Logger
)

// Implement ticket repo
type TicketRepositoryImpl struct {
	DB *sql.DB
}

func NewTicketRepository(db *sql.DB) *TicketRepositoryImpl {
	logger = config.NewLogger()
	return &TicketRepositoryImpl{
		DB: db,
	}
}

func (repo *TicketRepositoryImpl) CreateTicketProduct(ticketProd *dto.TicketProduct) error {
	const sqlQuery = `
		INSERT INTO stripe_ticket
		(id, ticket_id, price_id, name, description)	
		VALUES ($1, $2, $3, $4, $5)
	`

	res, err := repo.DB.Exec(sqlQuery, ticketProd.StripeID,
		ticketProd.TicketID, ticketProd.PriceID, ticketProd.Name, ticketProd.Description)

	if err != nil {
		logger.Errorf("STRIPE_TICKET(create): %v", err)
		return fmt.Errorf("Erro ao criar stripe ticket: %v", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		logger.Error("STRIPE_TICKET(create): no rows affected")
		return errors.New("Nenhum linha afetada ao criar stripe ticket")
	}

	return nil
}

func (repo *TicketRepositoryImpl) Create(ticket *domain.Ticket) (*dto.CreateTicketOutputDTO, error) {
	const sqlQuery = `
		INSERT INTO tickets
		(id, name, price, total_qty, available_qty, description, max_per_user, event_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	res, err := repo.DB.Exec(sqlQuery, ticket.ID, ticket.Name,
		ticket.Price, ticket.TotalQty, ticket.TotalQty, ticket.Description, ticket.MaxPerUser,
		ticket.EventId)

	if err != nil {
		logger.Errorf("TICKET(create): %v", err)
		return nil, fmt.Errorf("Erro ao criar ticket: %v", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		logger.Errorf("TICKET(create): %v", err)
		return nil, fmt.Errorf("Nenhuma linha afetada: %v", err)
	}
	return &dto.CreateTicketOutputDTO{
		ID:           ticket.ID,
		Name:         ticket.Name,
		Price:        ticket.Price,
		TotalQty:     ticket.TotalQty,
		AvailableQty: ticket.TotalQty,
		Description:  ticket.Description,
		MaxPerUser:   ticket.MaxPerUser,
		EventId:      ticket.EventId,
	}, nil
}

func (repo *TicketRepositoryImpl) GetTicketById(id string) (*dto.GetTicketDTO, error) {
	var ticket dto.GetTicketDTO

	const sqlQuery = `
	SELECT id, name, price, total_qty, available_qty, description, max_per_user, event_id
	FROM tickets
	WHERE id = $1
	LIMIT 1
	`
	row := repo.DB.QueryRow(sqlQuery, id)
	if err := row.Scan(&ticket.ID, &ticket.Name, &ticket.Price,
		&ticket.TotalQty, &ticket.AvailableQty, &ticket.Description, &ticket.MaxPerUser, &ticket.EventId); err != nil {
		logger.Errorf("TICKET(get_by_id): %v", err)
		return nil, fmt.Errorf("Error to get ticket by id: %v", err)
	}

	return &ticket, nil

}

func (repo *TicketRepositoryImpl) GetProductByTicketId(id string) (*dto.TicketProduct, error) {
	var ticketProduct dto.TicketProduct
	const sqlQuery = `
	SELECT id, price_id, ticket_id, name, description
	FROM stripe_ticket
	WHERE ticket_id = $1
	LIMIT 1
	`
	row := repo.DB.QueryRow(sqlQuery, id)
	err := row.Scan(&ticketProduct.StripeID, &ticketProduct.PriceID, &ticketProduct.TicketID,
		&ticketProduct.Name, &ticketProduct.Description)
	if err != nil {
		logger.Errorf("TICKET_PRODUCT(get_by_ticket_id): %v", err)
		return nil, fmt.Errorf("Error while getting product by ticket id: %v", err)
	}
	return &ticketProduct, nil
}

func (repo *TicketRepositoryImpl) UpdateAvailableQuantity(dto *dto.UpdateTicketAvailableQtyInputDTO) error {
	var operator string

	if dto.Type == "increment" {
		operator = "+"
	} else {
		operator = "-"
	}
	sqlQuery := fmt.Sprintf(`
	UPDATE tickets 	
	SET available_qty = available_qty %s $1
	WHERE id = $2
	`, operator)

	res, err := repo.DB.Exec(sqlQuery, dto.Quantity, dto.ID)
	if err != nil {
		logger.Errorf("TICKET(update_quantity): %v", err)
		return fmt.Errorf("Error while updating available quantity: %v", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		logger.Error("TICKET(update_quantity): No rows affected")
		return errors.New("No rows affected on ticket updating quantity")
	}
	return nil
}
