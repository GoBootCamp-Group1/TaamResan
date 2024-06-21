package address

import (
	"context"
	"errors"
)

type Address struct {
	ID    uint
	Title string
	Lat   float64
	Lng   float64
}

type Repo interface {
	Create(ctx context.Context, address *Address) error
	Update(ctx context.Context, address *Address) error
	Delete(ctx context.Context, address *Address) error
	GetByID(ctx context.Context, id uint) (*Address, error)
	GetAll(ctx context.Context) ([]*Address, error)
}

var (
	ErrAddressNotFound = errors.New("address not found")
	ErrInvalidLocation = errors.New("invalid address location")
)
