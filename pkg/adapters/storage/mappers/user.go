package mappers

import (
	"TaamResan/internal/user"
	"TaamResan/pkg/adapters/storage/entities"
)

func UserEntityToDomain(entity *entities.User) *user.User {
	return &user.User{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Email:     entity.Email,
		Password:  entity.Password,
		Role:      user.Role(entity.Role),
	}
}
