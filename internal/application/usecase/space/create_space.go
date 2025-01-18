package space

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
)

type CreateSpaceUseCase struct {
    SpaceRepo ports.SpaceRepository
}

func NewCreateSpaceUseCase(repo ports.SpaceRepository) *CreateSpaceUseCase {
    return &CreateSpaceUseCase{SpaceRepo: repo}
}

func (uc *CreateSpaceUseCase) Execute(space entity.Space) (entity.Space, error) {
    return uc.SpaceRepo.Create(space)
}
