package service

import (
	"context"

	"abc.com/demo/internal/domain"
)

type InventoryService interface {
	Create(ctx context.Context, order domain.Inventory) (domain.Inventory, error)
	Update(ctx context.Context, order domain.Inventory) (domain.Inventory, error)
}
