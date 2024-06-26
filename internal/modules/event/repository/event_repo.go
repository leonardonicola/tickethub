package repository

import (
	"database/sql"
	"fmt"
	"mime/multipart"

	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/event/domain"
	"github.com/leonardonicola/tickethub/internal/modules/event/dto"
	bucket "github.com/leonardonicola/tickethub/internal/pkg/s3"
)

type EventRepositoryImpl struct {
	db *sql.DB
}

var (
	logger *config.Logger
)

func NewEventRepository(db *sql.DB) *EventRepositoryImpl {
	logger = config.NewLogger()
	return &EventRepositoryImpl{
		db: db,
	}
}

func (repo *EventRepositoryImpl) Create(event *domain.Event, poster *multipart.FileHeader) (*dto.CreateEventOutputDTO, error) {
	imageId, err := bucket.UploadFileToBucket("tickethub", poster)

	if err != nil {
		return nil, err
	}

	const sqlQuery = `
		INSERT INTO events 
		(id, title, description, address, date, age_rating, genre_id, poster) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	res, err := repo.db.Exec(sqlQuery,
		event.ID, event.Title, event.Description, event.Address, event.Date, event.AgeRating,
		event.GenreID, imageId.Identifier)

	if err != nil {
		logger.Errorf("EVENT(create): %v", err)
		return nil, fmt.Errorf("EVENT(create): %v", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		// Rollback on image insertion
		if err := bucket.DeleteObjectFromBucket("tickethub", imageId.Identifier); err != nil {
			return nil, err
		}
		logger.Errorf("EVENT(create) - no rows affected: %v", err)
		return nil, fmt.Errorf("Error while creating event: %v", err)
	}
	return &dto.CreateEventOutputDTO{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		Address:     event.Address,
		Date:        event.Date,
		AgeRating:   event.AgeRating,
		Poster:      imageId.Identifier,
		GenreID:     event.GenreID,
	}, nil
}

func (repo *EventRepositoryImpl) GetById(id string) (*dto.GetEventByIdOutputDTO, error) {

	var event dto.GetEventByIdOutputDTO
	const sqlQuery = `
	SELECT id, title, date, address, poster, created_at, updated_at	
	FROM events
	WHERE id = $1
	LIMIT 1
	`
	row := repo.db.QueryRow(sqlQuery, id)
	err := row.Scan(&event.ID, &event.Title, &event.Date, &event.Address, &event.Poster, &event.CreatedAt, &event.UpdatedAt)
	if err != nil {
		logger.Errorf("EVENT(get by id): %v", err)
		return nil, fmt.Errorf("Evento inexistente")
	}
	return &event, nil
}

func (repo *EventRepositoryImpl) GetMany(query dto.GetManyEventsInputDTO) ([]dto.GetManyEventsOutputDTO, error) {
	events := make([]dto.GetManyEventsOutputDTO, 0)
	offset := (query.Page - 1) * query.Limit
	// Pesquisa por eventos que contenham o titulo ou descrição com a query do usuario
	// que vão acontecer no futuro
	const sqlQuery = `
		SELECT e.id, e.title, e.date, e.address, e.poster, g.name, e.created_at, e.updated_at
		FROM events e 
		INNER JOIN genres g
		ON g.id = e.genre_id
		WHERE e.date > NOW() 
		AND unaccent(e.searchable) ILIKE '%' || unaccent($1) || '%'  
		ORDER BY e.date ASC
		LIMIT $2
		OFFSET $3
	`
	// Execute the query with parameters
	rows, err := repo.db.Query(sqlQuery, query.Search, query.Limit, offset)
	if err != nil {
		logger.Errorf("EVENT(get_many) query: %v", err)
		return nil, fmt.Errorf("Error while getting events: %v", err)
	}

	for rows.Next() {
		var event dto.GetManyEventsOutputDTO

		if err := rows.Scan(&event.ID, &event.Title,
			&event.Date, &event.Address, &event.Poster, &event.Genre, &event.CreatedAt, &event.UpdatedAt); err != nil {
			logger.Errorf("EVENT(get_many) scan: %v", err)
			return nil, fmt.Errorf("Error while getting events: %v", err)
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		logger.Errorf("EVENT(get_many) rows error: %v", err)
		return nil, fmt.Errorf("Error while reading events: %v", err)
	}
	return events, nil
}
