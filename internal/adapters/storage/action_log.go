package storage

import (
	"TaamResan/internal/action_log"
	"TaamResan/internal/adapters/storage/entities"
	"TaamResan/internal/adapters/storage/mappers"
	"context"
	"errors"
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

var (
	ErrFetchingLogs = errors.New("error fetching logs")
)

func (w *actionLogRepo) Create(ctx context.Context, actionLog *action_log.ActionLog) (*action_log.ActionLog, error) {

	actionLogEntity := mappers.DomainToActionLogEntity(actionLog)

	if err := w.db.Create(&actionLogEntity).Error; err != nil {
		return nil, err
	}

	return mappers.ActionLogEntityToDomain(actionLogEntity), nil
}

func (w *actionLogRepo) GetAllByUserId(ctx context.Context, userId uint) ([]*action_log.ActionLog, error) {
	var logs []*entities.ActionLog
	if err := w.db.WithContext(ctx).Model(&entities.ActionLog{}).Where("user_id = ?", userId).Find(&logs).Error; err != nil {
		return nil, ErrFetchingLogs
	}

	var models []*action_log.ActionLog
	if len(logs) > 0 {
		for _, l := range logs {
			models = append(models, mappers.ActionLogEntityToDomain(l))
		}
	}

	return models, nil
}

func (w *actionLogRepo) GetAllByRestaurantId(ctx context.Context, restaurantId uint) ([]*action_log.ActionLog, error) {
	var logs []*entities.ActionLog
	if err := w.db.WithContext(ctx).Model(&entities.ActionLog{}).Where("entity_id = ? and entity_type = ?", restaurantId, action_log.RestaurantEntityType).Find(&logs).Error; err != nil {
		return nil, ErrFetchingLogs
	}

	var models []*action_log.ActionLog
	if len(logs) > 0 {
		for _, l := range logs {
			models = append(models, mappers.ActionLogEntityToDomain(l))
		}
	}

	return models, nil
}
