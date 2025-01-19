package space

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
	"errors"
)

type UpdateSpaceUseCase struct {
	SpaceRepo ports.SpaceRepository
}

func NewUpdateSpaceUseCase(repo ports.SpaceRepository) *UpdateSpaceUseCase {
	return &UpdateSpaceUseCase{SpaceRepo: repo}
}

func (uc *UpdateSpaceUseCase) Execute(id string, space entity.Space) (entity.Space, error) {
	
	existingSpace, err := uc.SpaceRepo.GetByID(id)
	if err != nil {
		return entity.Space{}, err
	}
	if existingSpace.ID == "" {
		return entity.Space{}, errors.New("space not found")
	}

	
	space.ID = existingSpace.ID

	updatedSpace, err := uc.SpaceRepo.Update(space)
	if err != nil {
		return entity.Space{}, err
	}

	return updatedSpace, nil
}
