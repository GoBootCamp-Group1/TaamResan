package action_log

import "context"

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops {
	return &Ops{
		repo: repo,
	}
}

func (o *Ops) Create(ctx context.Context, actionLog *ActionLog) (*ActionLog, error) {
	return o.repo.Create(ctx, actionLog)
}

func (o *Ops) GetAllByUserId(ctx context.Context, userId uint) ([]*ActionLog, error) {
	return o.repo.GetAllByUserId(ctx, userId)
}

func (o *Ops) GetAllByRestaurantId(ctx context.Context, restaurantId uint) ([]*ActionLog, error) {
	return o.repo.GetAllByRestaurantId(ctx, restaurantId)
}
