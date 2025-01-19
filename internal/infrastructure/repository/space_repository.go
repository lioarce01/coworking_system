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

func (r *GormSpaceRepository) Delete(id string) error {
    var space entity.Space
    result := r.DB.Where("id = ?", id).Delete(&space)
    if result.Error != nil {
        return result.Error
    }

    return nil
}

func (r *GormSpaceRepository) GetByID(id string) (entity.Space, error) {
	var space entity.Space
	result := r.DB.Where("id = ?", id).First(&space)
	if result.Error != nil {
		return entity.Space{}, result.Error
	}
	return space, nil
}


func (r *GormSpaceRepository) Update(space entity.Space) (entity.Space, error) {
    existingSpace, err := r.GetByID(space.ID)
    if err != nil {
        return entity.Space{}, err 
    }

    existingSpace.Name = space.Name
    existingSpace.Description = space.Description
    existingSpace.Capacity = space.Capacity
    existingSpace.IsAvailable = space.IsAvailable
    existingSpace.Price = space.Price
    existingSpace.Location = space.Location

    result := r.DB.Save(&existingSpace)
    if result.Error != nil {
        return entity.Space{}, result.Error
    }

    return existingSpace, nil
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
