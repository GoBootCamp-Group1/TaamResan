package mappers

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/user_roles"
)

func UserRolesEntityToDomain(entity *entities.UserRoles) *user_roles.UserRoles {
	return &user_roles.UserRoles{
		ID:     entity.ID,
		UserID: entity.UserId,
		RoleID: entity.RoleId,
	}
}

func DomainToUserRolesEntity(model *user_roles.UserRoles) *entities.UserRoles {
	return &entities.UserRoles{
		ID:     model.ID,
		UserId: model.UserID,
		RoleId: model.RoleID,
	}
}
