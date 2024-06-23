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
