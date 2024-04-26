package repository

import (
	"database/sql"
	"fmt"
	"mime/multipart"

	"github.com/leonardonicola/tickethub/config"
	"github.com/leonardonicola/tickethub/internal/modules/event/domain"
	"github.com/leonardonicola/tickethub/internal/modules/event/dto"
	utils "github.com/leonardonicola/tickethub/internal/pkg/utils"
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
	imageId, err := utils.UploadFileToBucket("tickethub", poster)

	if err != nil {
		return nil, err
	}

	res, err := repo.db.Exec("INSERT INTO events VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		event.ID, event.Title, event.Description, event.Address, event.Date, event.AgeRating,
		event.Genre, imageId.Identifier)

	if err != nil {
		logger.Errorf("EVENT(create): %v", err)
		return nil, fmt.Errorf("EVENT(create): %v", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		// Rollback on image insertion
		if err := utils.DeleteObjectFromBucket("tickethub", imageId.Identifier); err != nil {
			return nil, err
		}
		logger.Errorf("EVENT(create) - no rows affected: %v", err)
		return nil, fmt.Errorf("EVENT(create) - no rows affected: %v", err)
	}
	return &dto.CreateEventOutputDTO{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		Address:     event.Address,
		Date:        event.Date,
		AgeRating:   event.AgeRating,
		Poster:      imageId.Identifier,
		Genre:       event.Genre,
	}, nil
}

func (repo *EventRepositoryImpl) GetMany(query dto.GetManyEventsInputDTO) ([]dto.GetManyEventsOutputDTO, error) {

	events := make([]dto.GetManyEventsOutputDTO, 0)

	const sqlQuery = `
	SELECT id, title, description, date, age_rating, address, poster, genre 
	FROM events
	WHERE title ILIKE '%' || $1 || '%'
	ORDER BY id ASC
	LIMIT 10
	`

	// Execute the query with parameters
	rows, err := repo.db.Query(sqlQuery, query.Search)

	if err != nil {
		logger.Errorf("EVENT(get_many) query: %v", err)
		return nil, fmt.Errorf("Error while querying events: %v", err)
	}

	for rows.Next() {
		var event dto.GetManyEventsOutputDTO

		if err := rows.Scan(&event.ID, &event.Title, &event.Description,
			&event.Address, &event.Date, &event.AgeRating, &event.Genre, &event.Poster); err != nil {
			logger.Errorf("EVENT(get_many) scan: %v", err)
			return nil, fmt.Errorf("Error while reading events: %v", err)
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		logger.Errorf("EVENT(get_many) rows error: %v", err)
		return nil, fmt.Errorf("Error while reading events: %v", err)
	}
	return events, nil
}
