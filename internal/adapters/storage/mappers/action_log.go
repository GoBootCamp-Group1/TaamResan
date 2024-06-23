package mappers

import (
	"TaamResan/internal/action_log"
	"TaamResan/internal/adapters/storage/entities"
)

func ActionLogEntityToDomain(entity *entities.ActionLog) *action_log.ActionLog {
	return &action_log.ActionLog{
		UserID:   entity.UserID,
		Action:   entity.Action,
		IP:       entity.IP,
		Endpoint: entity.Endpoint,
		Payload:  entity.Payload,
		Method:   entity.Method,
	}
}

func DomainToActionLogEntity(model *action_log.ActionLog) *entities.ActionLog {
	return &entities.ActionLog{
		ID:       model.ID,
		UserID:   model.UserID,
		Action:   model.Action,
		IP:       model.IP,
		Endpoint: model.Endpoint,
		Payload:  model.Payload,
		Method:   model.Method,
	}
}
