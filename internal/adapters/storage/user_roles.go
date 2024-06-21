package storage

import (
	"TaamResan/internal/adapters/storage/mappers"
	"TaamResan/internal/user_roles"
	"context"
	"errors"
	"gorm.io/gorm"
)

type userRolesRepo struct {
	db *gorm.DB
}

func NewUserRolesRepo(db *gorm.DB) user_roles.Repo {
	return &userRolesRepo{
		db: db,
	}
}

var (
	ErrDeletingUserRole = errors.New("error deleting user role")
	ErrUserRoleNotFound = errors.New("user role not found")
)

func (r *userRolesRepo) Create(ctx context.Context, ur *user_roles.UserRoles) error {
	entity := mappers.DomainToUserRolesEntity(ur)
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *userRolesRepo) Update(ctx context.Context, ur *user_roles.UserRoles) error {
	entity := mappers.DomainToUserRolesEntity(ur)
	return r.db.WithContext(ctx).Save(&entity).Error
}

func (r *userRolesRepo) Delete(ctx context.Context, ur *user_roles.UserRoles) error {
	if err := r.db.WithContext(ctx).Model(&user_roles.UserRoles{}).Delete(&ur).Error; err != nil {
		return ErrDeletingUserRole
	}
	return nil
}

func (r *userRolesRepo) Get(ctx context.Context, id uint) (*user_roles.UserRoles, error) {
	var entity user_roles.UserRoles
	if err := r.db.WithContext(ctx).Model(&user_roles.UserRoles{}).Where("id = ?", id).First(&entity).Error; err != nil {
		return nil, ErrUserRoleNotFound
	}
	return &entity, nil
}
