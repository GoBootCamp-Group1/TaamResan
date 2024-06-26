package storage

import (
	"TaamResan/internal/order"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) order.Repo { return &orderRepo{db: db} }

func (o orderRepo) Create(ctx context.Context, data *order.InputData) (*order.Order, error) {

	fmt.Println("data inside repo: ", data)

	//TODO implement me
	return nil, nil
}
