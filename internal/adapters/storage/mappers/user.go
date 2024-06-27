package mappers

import (
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/user"
	"gorm.io/gorm"
)

func UserEntityToDomain(entity *entities.User) *user.User {
	return &user.User{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Mobile:    entity.Mobile,
		BirthDate: entity.BirthDate,
		Password:  entity.Password,
	}
}

func DomainToUserEntity(model *user.User) *entities.User {
	return &entities.User{
		Model:     gorm.Model{ID: model.ID},
		Uuid:      model.Uuid,
		Name:      model.Name,
		Email:     model.Email,
		Mobile:    model.Mobile,
		BirthDate: model.BirthDate,
		Password:  user.HashPassword(model.Password),
	}
}
