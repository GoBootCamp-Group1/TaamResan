package service

import (
	"TaamResan/internal/role"
	"context"
)

type RoleService struct {
	roleOps *role.Ops
}

func NewRoleService(roleOps *role.Ops) *RoleService {
	return &RoleService{roleOps: roleOps}
}

func (s *RoleService) InitializeRoles(ctx context.Context) error {
	roles := []role.Role{
		{ID: role.Customer, Name: role.CUSTOMER},
		{ID: role.Admin, Name: role.ADMIN},
		{ID: role.RestaurantOwner, Name: role.RESTAURANT_OWNER},
		{ID: role.RestaurantOperator, Name: role.RESTAURANT_OPERATOR},
	}

	for _, r := range roles {
		err := s.roleOps.Create(ctx, &r)
		if err != nil {
			return err
		}
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

func (s *RoleService) Get(ctx context.Context, id uint) (role.Role, error) {
	model, err := s.roleOps.Get(ctx, id)
	if err != nil {
		return role.Role{}, err
	}
	return *model, nil
}

func (s *RoleService) GetAll(ctx context.Context) ([]*role.Role, error) {
	models, err := s.roleOps.GetAll(ctx)
	if err != nil {
		return []*role.Role{}, err
	}
	return models, nil
}
