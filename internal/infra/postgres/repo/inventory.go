package repo

import (
	"context"

	"abc.com/demo/internal/domain"
	"gorm.io/gorm"
)

type InventoryRepo struct {
	db *gorm.DB
}

func NewInventoryRepo(db *gorm.DB) *InventoryRepo {
	return &InventoryRepo{
		db: db,
	}
}

func (repo *InventoryRepo) Create(ctx context.Context, inventory domain.Inventory) error {
	return nil
}

func (repo *InventoryRepo) Update(ctx context.Context, inventory domain.Inventory) error {
	return nil
}
