package repository

import (
	"database/sql"
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