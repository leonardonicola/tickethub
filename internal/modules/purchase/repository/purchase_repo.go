package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/purchase/domain"
)

type PurchaseRepositoryImpl struct {
	db *sql.DB
}

var (
	logger *config.Logger
)

func NewPurchaseRepository(db *sql.DB) *PurchaseRepositoryImpl {
	logger = config.NewLogger()
	return &PurchaseRepositoryImpl{db: db}
}

func (repo *PurchaseRepositoryImpl) UpdatePaymentStatus(status string, id string) error {
	const sqlQuery = `
  UPDATE purchase
  SET status = $1
  WHERE id = $2
  `

	res, err := repo.db.Exec(sqlQuery, status, id)

	if err != nil {
		logger.Errorf("PURCHASE(update_status): %v", err)
		return fmt.Errorf("Error while updating payment status: %v", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		logger.Error("PURCHASE(update_status): no rows affected")
		return errors.New("No rows affected while updating purchase status")
	}

	return nil
}

func (repo *PurchaseRepositoryImpl) Create(dto *domain.Purchase) error {
	const sqlQuery = `
	INSERT INTO purchase
	(id, ticket_id, user_id, quantity, status, payment_method)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	res, err := repo.db.Exec(sqlQuery, dto.ID, dto.TicketID,
		dto.UserID, dto.Quantity, dto.Status, dto.PaymentMethod)

	if err != nil {
		logger.Errorf("PURCHASE(create): %v", err)
		return fmt.Errorf("Error while creating a purchase: %v", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		logger.Error("PURCHASE(create): no rows affected")
		return errors.New("No rows affected while creating purchase")
	}

	return nil
}
