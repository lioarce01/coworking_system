package user

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
	"errors"
)

type UpdateUserUseCase struct {
	UserRepo ports.UserRepository
}

func NewUpdateSpaceUseCase(repo ports.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{UserRepo: repo}
}

func (uc *UpdateUserUseCase) Execute(id string, user entity.User) (entity.User, error) {
	existingUser, err := uc.UserRepo.GetByID(id)
	if err !=nil {
		return entity.User{}, err
	}

	if existingUser.ID == "" {
		return entity.User{}, errors.New("user not found")
	}

	user.ID = existingUser.ID

	updatedUser, err := uc.UserRepo.Update(user)
	if err != nil {
		return entity.User{}, err
	}
	return updatedUser, nil
}