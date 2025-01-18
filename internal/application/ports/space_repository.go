package ports

import "cowork_system/internal/domain/entity"

type SpaceRepository interface {
	ListAvailableSpaces() ([]entity.Space, error)
	Create(space entity.Space) (entity.Space, error)
	GetByID(id uint) (entity.Space, error)
	Update(space entity.Space) (entity.Space, error)
	Delete(id uint) error
}
