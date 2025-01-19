package space

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
	"cowork_system/internal/utils"
)

type CreateSpaceUseCase struct {
    SpaceRepo ports.SpaceRepository
}

func NewCreateSpaceUseCase(repo ports.SpaceRepository) *CreateSpaceUseCase {
    return &CreateSpaceUseCase{SpaceRepo: repo}
}

func (uc *CreateSpaceUseCase) Execute(space entity.Space) (entity.Space, error) {
	if space.ID == "" {
		space.ID = utils.GenerateUUID()
	}
	return uc.SpaceRepo.Create(space)
}
