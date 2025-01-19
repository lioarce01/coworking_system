package user

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
)

type GetUsersUseCase struct {
	UserRepo ports.UserRepository
}

func NewGetUsersUseCase(repo ports.UserRepository) *GetUsersUseCase {
	return &GetUsersUseCase{UserRepo: repo}
}

func (uc *GetUsersUseCase) Execute() ([]entity.User, error) {
	return uc.UserRepo.GetUsers()
}