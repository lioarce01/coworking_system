package space

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
	"errors"
)

type GetSpaceUseCase struct {
	SpaceRepo ports.SpaceRepository
}

func NewGetSpaceUseCase(repo ports.SpaceRepository) *GetSpaceUseCase {
	return &GetSpaceUseCase{SpaceRepo: repo}
}

func (uc *GetSpaceUseCase) Execute(id string) (entity.Space, error) {
	space, err := uc.SpaceRepo.GetByID(id)
	if err != nil {
		return entity.Space{}, err
	}
	if space.ID == "" {
		return entity.Space{}, errors.New("space not found")
	}
	return space, nil
}
