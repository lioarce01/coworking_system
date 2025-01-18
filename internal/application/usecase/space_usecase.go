package usecase

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
	"errors"
)

type SpaceUsecase struct {
	Repo ports.SpaceRepository
}

func NewSpaceUsecase(repo ports.SpaceRepository) *SpaceUsecase {
	return &SpaceUsecase{Repo: repo}
}

func (uc *SpaceUsecase) GetSpaces() ([]entity.Space, error) {
	spaces, err := uc.Repo.ListAvailableSpaces()
	if err != nil {
		return nil, err
	}
	return spaces, nil
}

func (uc *SpaceUsecase) CreateSpace(space entity.Space) (entity.Space, error) {
	if space.Name == "" {
		return entity.Space{}, errors.New("name is required")
	}
	createdSpace, err := uc.Repo.Create(space)
	if err != nil {
		return entity.Space{}, err
	}
	return createdSpace, nil
}
