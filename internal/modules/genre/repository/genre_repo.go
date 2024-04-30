package repository

import (
	"database/sql"
	"fmt"

	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/genre/domain"
)

type GenreRepositoryImpl struct {
	db *sql.DB
}

var (
	logger *config.Logger
)

func NewGenreRepository(db *sql.DB) *GenreRepositoryImpl {
	logger = config.NewLogger()
	return &GenreRepositoryImpl{
		db: db,
	}
}

func (repo *GenreRepositoryImpl) GetById(id string) (*domain.Genre, error) {
	var genre domain.Genre
	row := repo.db.QueryRow("SELECT * FROM genres WHERE id = $1", id)

	if err := row.Scan(&genre.ID, &genre.Name); err != nil {
		logger.Errorf("GENRE(get_by_id): %v", err)
		return nil, fmt.Errorf("Error while querying for genre: %v", err)
	}

	return &genre, nil

}
