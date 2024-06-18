package mappers

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/user"
)

func UserEntityToDomain(entity *entities.User) *user.User {
	return &user.User{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Mobile:    entity.Mobile,
		BirthDate: entity.BirthDate,
		Password:  entity.Password,
		Roles:     []user.Role{user.Customer}, // TODO: fix this when Role entity is created
	}
}