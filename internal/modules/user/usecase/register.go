package usecase

import (
	"github.com/leonardonicola/tickethub/internal/modules/user/domain"
	"github.com/leonardonicola/tickethub/internal/modules/user/dto"
	"github.com/leonardonicola/tickethub/internal/modules/user/ports"
)

type RegisterUseCase struct {
	Repository ports.UserRepository
}

func (uc *RegisterUseCase) Execute(payload *dto.CreateUserInputDTO) (*dto.CreateUserOutputDTO, error) {
	user := domain.NewUser(payload.Name, payload.Surname, payload.Email, payload.Address, payload.CPF, payload.Password)

	newUser, err := uc.Repository.Create(user)
	if err != nil {
		return nil, err
	}

	output := &dto.CreateUserOutputDTO{
		ID:      newUser.ID,
		Name:    newUser.Name,
		Surname: newUser.Surname,
		Email:   newUser.Email,
		Address: newUser.Address,
		CPF:     newUser.CPF,
	}

	return output, nil
}
