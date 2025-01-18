package repository

import (
	"cowork_system/internal/domain/entity"

	"gorm.io/gorm"
)

type GormSpaceRepository struct {
	DB *gorm.DB
}

func NewSpaceRepository(db *gorm.DB) *GormSpaceRepository {
	return &GormSpaceRepository{DB: db}
}

func (repo *GormSpaceRepository) ListAvailableSpaces() ([]entity.Space, error) {
	var spaces []entity.Space
	if err := repo.DB.Find(&spaces).Error; err != nil {
		return nil, err
	}
	return spaces, nil
}

func (repo *GormSpaceRepository) Create(space entity.Space) (entity.Space, error) {
	if err := repo.DB.Create(&space).Error; err != nil {
		return entity.Space{}, err
	}
	return space, nil
}
