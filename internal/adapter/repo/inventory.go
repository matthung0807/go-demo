package repo

import (
	"context"

	"abc.com/demo/internal/domain"
)

type InventoryRepo interface {
	Create(ctx context.Context, inventory domain.Inventory) error
	Update(ctx context.Context, inventory domain.Inventory) error
}
