package storage

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/role"
	"context"
	"errors"
	"gorm.io/gorm"
)

type roleRepo struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) role.Repo {
	return &roleRepo{
		db: db,
	}
}

var (
	ErrCreatingRole      = errors.New("error creating role")
	ErrUpdatingRole      = errors.New("error updating role")
	ErrDeletingRole      = errors.New("error deleting role")
	ErrRoleNotFound      = errors.New("role doesn't exist")
	ErrRoleExists        = errors.New("role already exists")
	ErrNonUniqueRoleName = errors.New("role with this name exists")
)

func (r roleRepo) Create(ctx context.Context, role *role.Role) error {
	var existingRole1 entities.Role
	var existingRole2 entities.Role
	err1 := r.db.WithContext(ctx).Model(&entities.Role{}).Where("id = ?", role.ID).First(&existingRole1).Error
	err2 := r.db.WithContext(ctx).Model(&entities.Role{}).Where("name = ?", role.Name).First(&existingRole2).Error

	if err1 != nil && err2 != nil {
		if errors.Is(err1, gorm.ErrRecordNotFound) && errors.Is(err2, gorm.ErrRecordNotFound) {
			entity := mappers.DomainToRoleEntity(role)
			if err := r.db.WithContext(ctx).Create(&entity).Error; err != nil {
				return ErrCreatingRole
			}
			return nil
		}
		return ErrCreatingRole
	}

	return ErrRoleExists
}

func (r roleRepo) Update(ctx context.Context, role *role.Role) error {
	var existingRole entities.Role
	err := r.db.WithContext(ctx).Model(&entities.Role{}).Where("id = ?", role.ID).First(&existingRole).Error
	if err != nil {
		return ErrRoleNotFound
	}

	if existingRole.Name != role.Name {
		var existingRole2 entities.Role
		err = r.db.WithContext(ctx).Model(&entities.Role{}).Where("name = ?", role.Name).First(&existingRole2).Error
		if err == nil {
			return ErrNonUniqueRoleName
		}

		existingRole.Name = role.Name
		if err := r.db.WithContext(ctx).Model(&entities.Role{}).Where("id = ?", role.ID).Updates(&existingRole).Error; err != nil {
			return ErrUpdatingRole
		}
	}

	return nil
}

func (r roleRepo) Delete(ctx context.Context, id uint) error {
	var existingRole entities.Role
	if err := r.db.WithContext(ctx).Model(&entities.Role{}).Where("id = ?", id).First(&existingRole).Error; err != nil {
		return ErrRoleNotFound
	}
	if err := r.db.WithContext(ctx).Model(&entities.Role{}).Delete(&existingRole).Error; err != nil {
		return ErrDeletingRole
	}
	return nil
}

func (r roleRepo) GetByName(ctx context.Context, name string) (*role.Role, error) {
	var existingRole entities.Role
	err := r.db.WithContext(ctx).Model(&entities.Role{}).Where("name = ?", name).First(&existingRole).Error
	if err != nil {
		return nil, ErrRoleNotFound
	}
	model := mappers.RoleEntityToDomain(&existingRole)
	return model, nil
}
