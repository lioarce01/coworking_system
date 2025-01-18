package space

import (
	"cowork_system/internal/application/ports"
	"errors"
)

type DeleteSpaceUseCase struct {
	SpaceRepo ports.SpaceRepository
}

func NewDeleteSpaceUseCase(repo ports.SpaceRepository) *DeleteSpaceUseCase {
	return &DeleteSpaceUseCase{SpaceRepo: repo}
}

func (uc *DeleteSpaceUseCase) Execute(id uint) error {
	
	space, err := uc.SpaceRepo.GetByID(id)
	if err != nil {
		return err
	}
	if space.ID == 0 {
		return errors.New("space not found")
	}

	err = uc.SpaceRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
