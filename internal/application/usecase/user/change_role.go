package user

import (
	"cowork_system/internal/application/ports"
	"cowork_system/internal/domain/entity"
	"errors"
)

type ChangeRoleUseCase struct {
	UserRepo ports.UserRepository
}

func NewChangeRoleUseCase(repo ports.UserRepository) *ChangeRoleUseCase {
	return &ChangeRoleUseCase{UserRepo: repo}
}

func (uc *ChangeRoleUseCase) Execute(adminID,targetID string, newRole entity.Role) error {
	adminUser, err := uc.UserRepo.GetByID(adminID)
	if err != nil {
		return err
	}

	if adminUser.Role != entity.Admin {
		return errors.New("only admins can change roles")
	}

	if newRole != entity.Normal && newRole != entity.Admin {
		return errors.New("invalid role")
	}

	return uc.UserRepo.ChangeRole(targetID,newRole)
}