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
	SearchRestaurants(ctx context.Context, searchData *RestaurantSearch) ([]*Restaurant, error)
}

type Restaurant struct {
	ID             uint            `json:"id"`
	Name           string          `json:"name"`
	OwnedBy        uint            `json:"owned_by"`
	ApprovalStatus uint            `json:"approval_status"`
	Address        address.Address `json:"address"`
	CourierSpeed   float64         `json:"courier_speed"`
}

type RestaurantSearch struct {
	Name       string
	CategoryID *uint
	Lat        *float64
	Lng        *float64
}

// ApprovalStatus
const (
	NotApproved uint = iota + 1
	Approved
)
