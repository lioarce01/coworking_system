// internal/infrastructure/repository/space_repository.go
package repository

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"

	"gorm.io/gorm"
)

type GormSpaceRepository struct {
    DB *gorm.DB
}

func NewGormSpaceRepository(db *gorm.DB) ports.SpaceRepository {
    return &GormSpaceRepository{DB: db}
}

func (r *GormSpaceRepository) ListAvailableSpaces() ([]entity.Space, error) {
    var spaces []entity.Space
    result := r.DB.Find(&spaces)
    if result.Error != nil {
        return nil, result.Error
    }
    return spaces, nil
}

func (r *GormSpaceRepository) Create(space entity.Space) (entity.Space, error) {
    result := r.DB.Create(&space)
    if result.Error != nil {
        return entity.Space{}, result.Error
    }
    return space, nil
}
