package service

import (
	"TaamResan/internal/action_log"
	"context"
	"errors"
	"fmt"
)

type ActionLogService struct {
	actionLogOps *action_log.Ops
}

var (
	ErrCreateActionLog = errors.New("can not store action log")
)

func NewActionLogService(actionLogOps *action_log.Ops) *ActionLogService {
	return &ActionLogService{
		actionLogOps: actionLogOps,
	}
}

func (s *ActionLogService) Create(ctx context.Context, actionLog *action_log.ActionLog) (*action_log.ActionLog, error) {
	log, err := s.actionLogOps.Create(ctx, actionLog)

	if err != nil {
		return nil, fmt.Errorf(ErrCreateActionLog.Error()+": %w", err)
	}

	return log, nil
}
