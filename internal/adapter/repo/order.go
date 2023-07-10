package repo

import (
	"context"

	"abc.com/demo/internal/domain"
)

type OrderRepo interface {
	Create(ctx context.Context, order domain.Order) error
	Delete(ctx context.Context, id string) error
}
