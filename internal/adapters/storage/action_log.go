package storage

import (
	"TaamResan/internal/action_log"
	"TaamResan/internal/adapters/storage/mappers"
	"context"
	"gorm.io/gorm"
)

type actionLogRepo struct {
	db   *gorm.DB
	repo action_log.Repo
}

func NewActionLogRepo(db *gorm.DB) action_log.Repo {
	return &actionLogRepo{
		db: db,
	}
}

func (w *actionLogRepo) Create(ctx context.Context, actionLog *action_log.ActionLog) (*action_log.ActionLog, error) {

	actionLogEntity := mappers.DomainToActionLogEntity(actionLog)

	if err := w.db.Create(&actionLogEntity).Error; err != nil {
		return nil, err
	}

	return mappers.ActionLogEntityToDomain(actionLogEntity), nil
}
