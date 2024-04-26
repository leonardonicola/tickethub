package dto

type GetManyEventsInputDTO struct {
	// Caso não passem descrição, é obrigatório passar título
	Search string `form:"search" validate:"required"`
}

type GetManyEventsOutputDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Date        string `json:"date"`
	AgeRating   uint8  `json:"age_rating"`
	Poster      string `json:"poster_url"`
	Genre       string `json:"genre"`
}
