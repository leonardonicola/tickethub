package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Events são os shows/festivais/eventos

// Eles também contam com:
// - Tipos de ingressos/tickets e seus preços

type Event struct {
	ID          string
	Title       string
	Description string
	Address     string
	Date        string
	AgeRating   uint8
	GenreID     string
}

func NewEvent(title, description, address, date, genreId string, agerating uint8) (*Event, error) {
	id := uuid.NewString()

	if _, err := time.Parse("2006-01-02", date); err != nil {
		return nil, errors.New("EVENT - invalid date")
	}
	return &Event{
		ID:          id,
		Title:       title,
		Description: description,
		Address:     address,
		Date:        date,
		GenreID:     genreId,
		AgeRating:   agerating,
	}, nil
}
