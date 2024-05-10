package usecase

import (
	"github.com/leonardonicola/tickethub/internal/modules/user/dto"
	"github.com/leonardonicola/tickethub/internal/modules/user/ports"
)

type GetUserByIdUseCase struct {
	UserRepository ports.UserRepository
}

func (uc *GetUserByIdUseCase) Execute(id string) (*dto.GetUserOutputDTO, error) {
	return uc.UserRepository.GetById(id)
}
