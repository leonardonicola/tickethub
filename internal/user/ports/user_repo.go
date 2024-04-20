package ports

import "github.com/leonardonicola/tickethub/internal/user/domain"

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	// Delete(id string) bool
	// Update(user *domain.User) (*domain.User, error)
}
