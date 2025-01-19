package ports

import "cowork_system/internal/domain/entity"

type UserRepository interface {
	GetUsers() ([]entity.User, error)
	GetByID(id string) (entity.User, error)
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(id string) error
	ChangeRole(targetID string, newRole entity.Role) error
}