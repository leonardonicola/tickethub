package dto

import "time"

type GetEventByIdOutputDTO struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Address   string    `json:"address"`
	Date      string    `json:"date"`
	Poster    string    `json:"poster_url"`
	Genre     string    `json:"genre"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
