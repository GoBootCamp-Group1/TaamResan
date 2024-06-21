package service

import (
	"TaamResan/internal/role"
	"context"
	"errors"
)

type RoleService struct {
	roleOps *role.Ops
}

func NewRoleService(roleOps *role.Ops) *RoleService {
	return &RoleService{roleOps: roleOps}
}

var (
	ErrCreatingRole = errors.New("can not create role")
	ErrDeletingRole = errors.New("can not delete role")
	ErrReadingRole  = errors.New("can not read role")
)

func (s *RoleService) InitializeRoles(ctx context.Context) error {
	roles := []role.Role{
		{ID: role.Customer, Name: role.CUSTOMER},
		{ID: role.Admin, Name: role.ADMIN},
		{ID: role.RestaurantOwner, Name: role.RESTAURANT_OWNER},
		{ID: role.RestaurantOperator, Name: role.RESTAURANT_OPERATOR},
	}

	for _, r := range roles {
		return s.roleOps.Create(ctx, &r)
	}

	return nil
}

func (s *RoleService) Create(ctx context.Context, r *role.Role) error {
	return s.roleOps.Create(ctx, r)
}

func (s *RoleService) Update(ctx context.Context, r *role.Role) error {
	return s.roleOps.Update(ctx, r)
}

func (s *RoleService) Delete(ctx context.Context, id uint) error {
	return s.roleOps.Delete(ctx, id)
}

func (s *RoleService) GetByName(ctx context.Context, name string) (role.Role, error) {
	model, err := s.roleOps.GetByName(ctx, name)
	if err != nil {
		return role.Role{}, ErrReadingRole
	}
	return *model, nil
}
