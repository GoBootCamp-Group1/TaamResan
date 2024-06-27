package action_log

import (
	"context"
	"github.com/google/uuid"
)

const LogCtxKey = "Action-Log"

type ActionLog struct {
	ID         uuid.UUID
	UserID     *uint
	Action     string
	IP         string
	Endpoint   string
	Payload    map[string]any
	Method     string
	EntityType string
	EntityID   uint
}

type Repo interface {
	Create(ctx context.Context, actionLog *ActionLog) (*ActionLog, error)
	Update(ctx context.Context, actionLog *ActionLog) error
	GetAllByUserId(ctx context.Context, userId uint) ([]*ActionLog, error)
	GetAllByRestaurantId(ctx context.Context, restaurantId uint) ([]*ActionLog, error)
}

var (
	RestaurantEntityType = "restaurant"
)
