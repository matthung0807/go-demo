package impl

import (
	"context"

	"abc.com/demo/internal/adapter/repo"
	"abc.com/demo/internal/domain"
)

type InventoryService struct {
	inventoryRepo repo.InventoryRepo
}

func NewInventoryService(inventoryRepo repo.InventoryRepo) *InventoryService {
	return &InventoryService{
		inventoryRepo: inventoryRepo,
	}
}

func (s *InventoryService) Create(ctx context.Context, order domain.Inventory) (domain.Inventory, error) {
	err := s.inventoryRepo.Create(ctx, order)
	if err != nil {
		return domain.Inventory{}, err
	}
	return order, nil
}

func (s *InventoryService) Update(ctx context.Context, order domain.Inventory) (domain.Inventory, error) {
	err := s.inventoryRepo.Update(ctx, order)
	if err != nil {
		return domain.Inventory{}, err
	}
	return order, nil
}
