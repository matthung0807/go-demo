package repo

import (
	"context"
	"log"

	"abc.com/demo/internal/domain"
	"gorm.io/gorm"
)

type SagaRepo struct {
	db *gorm.DB
}

func NewSagaRepo(db *gorm.DB) *SagaRepo {
	return &SagaRepo{
		db: db,
	}
}

func (repo *SagaRepo) Insert(ctx context.Context, saga domain.Saga) error {
	log.Printf("insert saga=[%v]", saga)
	return nil
}
