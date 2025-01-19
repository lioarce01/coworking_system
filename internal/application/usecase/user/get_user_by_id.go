package user

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
	"errors"
)

type GetUserUseCase struct {
	UserRepo ports.UserRepository
}

func NewGetUserUseCase(repo ports.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{UserRepo: repo}
}

func (uc *GetUserUseCase) Execute(id string) (entity.User, error) {
	user, err := uc.UserRepo.GetByID(id)
	if err != nil {
		return entity.User{}, err
	}

	if user.ID == "" {
		return entity.User{}, errors.New("user not found")
	}

	return user, nil
}