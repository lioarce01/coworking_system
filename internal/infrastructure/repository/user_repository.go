package repository

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) ports.UserRepository {
	return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) Create(user entity.User) (entity.User, error) {
	result := r.DB.Create(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}

func (r *GormUserRepository) GetByID(id string) (entity.User, error) {
	var user entity.User
	result := r.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}

func (r *GormUserRepository) GetUsers() ([]entity.User, error) {
	var users []entity.User
	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *GormUserRepository) Update(user entity.User) (entity.User, error) {
	existingUser, err := r.GetByID(user.ID)
	if err != nil {
		return entity.User{}, err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	existingUser.Phone = user.Phone

	result := r.DB.Save(&existingUser)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return existingUser, nil
}
func (r *GormUserRepository) Delete(id string) error {
	var user entity.User
	result := r.DB.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
