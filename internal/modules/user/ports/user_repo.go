package ports

import (
	"github.com/leonardonicola/tickethub/internal/modules/user/domain"
	"github.com/leonardonicola/tickethub/internal/modules/user/dto"
)

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	GetById(id string) (*dto.GetUserOutputDTO, error)
	// Delete(id string) bool
	// Update(user *domain.User) (*domain.User, error)
}
