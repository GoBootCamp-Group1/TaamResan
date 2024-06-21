package user_roles

import "context"

type Repo interface {
	Create(ctx context.Context, ur *UserRoles) error
	Update(ctx context.Context, ur *UserRoles) error
	GetByID(ctx context.Context, id uint) (*UserRoles, error)
}

type UserRoles struct {
	ID     uint
	UserID uint
	RoleID uint
}
