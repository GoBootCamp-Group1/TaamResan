package storage

import (
	"TaamResan/internal/user_roles"
	"context"
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

func (u userRolesRepo) Create(ctx context.Context, ur *user_roles.UserRoles) error {
	panic("implement me")
}

func (u userRolesRepo) Update(ctx context.Context, ur *user_roles.UserRoles) error {
	panic("implement me")
}

func (u userRolesRepo) GetByID(ctx context.Context, id uint) (*user_roles.UserRoles, error) {
	panic("implement me")
}
