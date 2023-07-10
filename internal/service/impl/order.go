package impl

import (
	"context"

	"abc.com/demo/internal/adapter/repo"
	"abc.com/demo/internal/domain"
)

type OrderService struct {
	orderRepo repo.OrderRepo
}

func NewOrderService(
	orderRepo repo.OrderRepo,
) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (s *OrderService) Create(ctx context.Context, order domain.Order) error {
	err := s.orderRepo.Create(ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (s *OrderService) Delete(ctx context.Context, id string) error {
	err := s.orderRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
