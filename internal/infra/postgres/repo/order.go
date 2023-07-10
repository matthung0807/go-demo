package repo

import (
	"context"

	"abc.com/demo/internal/domain"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{
		db: db,
	}
}

func (repo *OrderRepo) Create(ctx context.Context, order domain.Order) error {
	return nil
}

func (repo *OrderRepo) Delete(ctx context.Context, id string) error {
	return nil
}

func (repo *OrderRepo) GetById(ctx context.Context, id string) *domain.Order {
	return nil
}
