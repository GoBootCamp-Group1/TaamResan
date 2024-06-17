package mappers

import (
	"TaamResan/internal/user"
	"TaamResan/pkg/adapters/storage/entities"
)

func UserEntityToDomain(entity *entities.User) *user.User {
	return &user.User{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Mobile:    entity.Mobile,
		BirthDate: entity.BirthDate,
		Password:  entity.Password,
		Role:      user.RoleUser, // TODO: fix this when Role entity is created
	}
}
