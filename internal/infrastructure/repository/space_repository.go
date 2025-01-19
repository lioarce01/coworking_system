package repository

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
	"errors"

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

    result := r.DB.Model(&existingSpace).Updates(entity.Space{
        Name:        space.Name,
        Description: space.Description,
        Capacity:    space.Capacity,
        IsAvailable: space.IsAvailable,
        Price:       space.Price,
        Location:    space.Location,
    })
    if result.Error != nil {
        return entity.Space{}, result.Error
    }

    return existingSpace, nil
}



func (r *GormSpaceRepository) ListAvailableSpaces() ([]entity.Space, error) {
	var spaces []entity.Space
	result := r.DB.Where("is_available = ?", true).Find(&spaces)
	if result.Error != nil {
		return nil, result.Error
	}
	return spaces, nil
}

func (r *GormSpaceRepository) Create(space entity.Space) (entity.Space, error) {
	if space.Capacity <= 0 {
		return entity.Space{}, errors.New("capacity must be greater than 0")
	}
	result := r.DB.Create(&space)
	if result.Error != nil {
		return entity.Space{}, result.Error
	}
	return space, nil
}

func (r *GormSpaceRepository) CountActiveReservations(spaceID string) (int, error) {
	var count int64
	result := r.DB.Model(&entity.Reservation{}).
		Where("space_id = ? AND status = ?", spaceID, entity.Confirmed).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}

func (r *GormSpaceRepository) SetAvailability(spaceID string, isAvailable bool) error {
	result := r.DB.Model(&entity.Space{}).
		Where("id = ?", spaceID).
		Update("is_available", isAvailable)
	if result.Error != nil {
		return result.Error
	}
	return nil
}