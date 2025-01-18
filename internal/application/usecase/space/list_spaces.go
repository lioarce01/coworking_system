package space

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
)

type ListSpacesUseCase struct {
    SpaceRepo ports.SpaceRepository
}

func NewListSpacesUseCase(repo ports.SpaceRepository) *ListSpacesUseCase {
    return &ListSpacesUseCase{SpaceRepo: repo}
}

func (uc *ListSpacesUseCase) Execute() ([]entity.Space, error) {
    return uc.SpaceRepo.ListAvailableSpaces()
}
