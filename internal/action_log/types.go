package action_log

import (
	"context"
	"github.com/google/uuid"
)

const LogCtxKey = "Action-Log"

type ActionLog struct {
	ID       uuid.UUID
	UserID   *uint
	Action   string
	IP       string
	Endpoint string
	Payload  map[string]any
	Method   string
}

type Repo interface {
	Create(ctx context.Context, actionLog *ActionLog) (*ActionLog, error)
}
