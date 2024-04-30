package dto

import "mime/multipart"

type CreateEventInputDTO struct {
	Title       string                `form:"title" validate:"required,gt=2,lt=170"`
	Description string                `form:"description" validate:"required,gt=1,lt=600"`
	Address     string                `form:"address" validate:"required,gt=3,lt=254"`
	Date        string                `form:"date" validate:"required,datetime=2006-01-02"`
	AgeRating   uint8                 `form:"age_rating" validate:"required,number,gte=1,lte=60"`
	Poster      *multipart.FileHeader `form:"poster" validate:"required"`
	GenreID     string                `form:"genre_id" validate:"required,gt=2,lt=50,uuid"`
}

type CreateEventOutputDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Date        string `json:"date"`
	AgeRating   uint8  `json:"age_rating"`
	Poster      string `json:"poster_url"`
	GenreID     string `json:"genre_id"`
}
