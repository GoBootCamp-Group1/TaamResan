package mappers

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/role"
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
		Roles: []role.Role{
			{ID: role.Customer, Name: role.CUSTOMER},
		}, // TODO
	}
}

func DomainToUserEntity(model *user.User) *entities.User {
	return &entities.User{
		ID:        model.ID,
		Uuid:      model.Uuid,
		Name:      model.Name,
		Email:     model.Email,
		Mobile:    model.Mobile,
		BirthDate: model.BirthDate,
		Password:  user.HashPassword(model.Password),
	}
}
