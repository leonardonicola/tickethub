package user_test

import (
	"testing"

	"github.com/leonardonicola/tickethub/internal/user/domain"
	"github.com/leonardonicola/tickethub/internal/user/dto"
	"github.com/leonardonicola/tickethub/internal/user/usecase"
)

type MockRepository struct{}

func (m *MockRepository) Create(user *domain.User) (*domain.User, error) {
	return &domain.User{
		ID:      user.ID,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Address: user.Address,
		CPF:     user.CPF,
	}, nil
}

func TestUserService(t *testing.T) {
	mockRepo := &MockRepository{}

	registerUc := usecase.RegisterUseCase{
		Repository: mockRepo,
	}
	fakeUser := dto.CreateUserInputDTO{
		Name:     "Faked",
		Surname:  "Surname",
		Email:    "leonicola@hotmail.com",
		CPF:      "12345678910",
		Address:  "Av. Paulista",
		Password: "SuperSecret123123",
	}
	t.Run("success: register user", func(t *testing.T) {
		res, err := registerUc.Execute(fakeUser)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if res.ID == "" {
			t.Errorf("Wanted: %v\n Received: %v\n", fakeUser, res)
		}

	})
}
