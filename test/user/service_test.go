package user_test

import (
	"reflect"
	"testing"

	"github.com/leonardonicola/tickethub/internal/modules/user/domain"
	"github.com/leonardonicola/tickethub/internal/modules/user/dto"
	"github.com/leonardonicola/tickethub/internal/modules/user/usecase"
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

func (m *MockRepository) GetById(id string) (*dto.GetUserOutputDTO, error) {
	return nil, nil
}

func TestUserService(t *testing.T) {
	mockRepo := &MockRepository{}

	registerUc := usecase.RegisterUseCase{
		Repository: mockRepo,
	}
	fakeUser := &dto.CreateUserInputDTO{
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
		vInput := reflect.ValueOf(*fakeUser)
		vOutput := reflect.ValueOf(*res)

		// Iterate over fields of CreateUserInputDTO
		for i := range vInput.NumField() {
			fieldInput := vInput.Field(i)
			fieldOutput := vOutput.FieldByName(vInput.Type().Field(i).Name)

			if !fieldOutput.IsValid() {
				continue
			}

			// Compare values if field exists in CreateUserOutputDTO
			if fieldInput.Interface() != fieldOutput.Interface() {
				t.Errorf("WANTED: %s\n HAD: %s", fieldInput, fieldOutput)
			}
		}

	})
}
