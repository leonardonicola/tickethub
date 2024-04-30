package dto

type GetManyEventsInputDTO struct {
	// Caso não passem descrição, é obrigatório passar título
	Search string `form:"search" validate:"required"`
	Limit  uint8  `form:"limit" validate:"required"`
	Page   uint8  `form:"page" validate:"required"`
}

type GetManyEventsOutputDTO struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Address string `json:"address"`
	Date    string `json:"date"`
	Poster  string `json:"poster_url"`
	Genre   string `json:"genre"`
}
