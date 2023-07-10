package service

import (
	"context"

	"abc.com/demo/internal/domain"
)

type OrderService interface {
	Create(ctx context.Context, order domain.Order) error
	Delete(ctx context.Context, id string) error
}
