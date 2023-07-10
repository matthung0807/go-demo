package repo

import (
	"context"

	"abc.com/demo/internal/domain"
)

type SagaRepo interface {
	Insert(ctx context.Context, saga domain.Saga) error
}
