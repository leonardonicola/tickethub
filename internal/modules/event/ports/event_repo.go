package ports

import (
	"mime/multipart"

	"github.com/leonardonicola/tickethub/internal/modules/event/domain"
	"github.com/leonardonicola/tickethub/internal/modules/event/dto"
)

type EventRepository interface {
	Create(event *domain.Event, poster *multipart.FileHeader) (*dto.CreateEventOutputDTO, error)
	GetMany(search dto.GetManyEventsInputDTO) ([]dto.GetManyEventsOutputDTO, error)
}
