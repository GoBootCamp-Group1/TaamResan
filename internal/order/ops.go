package order

import (
	"context"
)

type Ops struct {
	repo Repo
}

func NewOps(repo Repo) *Ops { return &Ops{repo: repo} }

func (o Ops) Create(ctx context.Context, data *InputData) (*Order, error) {

	return o.repo.Create(ctx, data)
}
