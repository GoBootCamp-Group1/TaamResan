package mappers

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/role"
	"gorm.io/gorm"
)

func RoleEntityToDomain(entity *entities.Role) *role.Role {
	return &role.Role{
		ID:   entity.ID,
		Name: entity.Name,
	}
}

func DomainToRoleEntity(model *role.Role) *entities.Role {
	return &entities.Role{
		Model: gorm.Model{
			ID: model.ID,
		},
		Name: model.Name,
	}
}
