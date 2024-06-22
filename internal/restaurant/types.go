package restaurant

import (
	"TaamResan/internal/address"
	"context"
)

type Repo interface {
	Create(ctx context.Context, restaurant *Restaurant) (uint, error)
	Update(ctx context.Context, restaurant *Restaurant) error
	Delete(ctx context.Context, id uint) error
	GetById(ctx context.Context, id uint) (*Restaurant, error)
	GetAll(ctx context.Context) ([]*Restaurant, error)
	Approve(ctx context.Context, id uint) error
	DelegateOwnership(ctx context.Context, id uint, newOwnerId uint) error
}

type Restaurant struct {
	ID             uint
	Name           string
	OwnedBy        uint
	ApprovalStatus uint
	Address        address.Address
	CourierSpeed   float64
}

// ApprovalStatus
const (
	NotApproved uint = iota + 1
	Approved
)
