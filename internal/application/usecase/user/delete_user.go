package user

import (
	"cowork_system/internal/application/ports"
	"errors"
)

type DeleteUserUseCase struct {
	UserRepo ports.UserRepository
}

func NewDeleteUserUseCase(repo ports.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{UserRepo: repo}
}

func (uc *DeleteUserUseCase) Execute(id string) error {
	user, err := uc.UserRepo.GetByID(id)
	if err != nil {
		return err
	}
	if user.ID == "" {
		return errors.New("user not found")
	}

	err = uc.UserRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}