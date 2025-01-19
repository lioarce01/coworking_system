package user

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
	"cowork_system/internal/utils"
)

type CreateUserUseCase struct {
	UserRepo ports.UserRepository
}

func NewCreateUserUseCase(repo ports.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{UserRepo: repo}
}

func (uc *CreateUserUseCase) Execute(user entity.User) (entity.User, error) {
	if user.ID == "" {
		user.ID = utils.GenerateUUID()
	}

	return uc.UserRepo.Create(user)
}