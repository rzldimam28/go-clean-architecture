package cache

import (
	"belajar-golang-clean-architecture/model/domain"
	"context"
)

type ProductCache interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) ([]domain.Product, error)
}