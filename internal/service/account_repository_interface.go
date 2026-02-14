package service

import (
	"context"
	"github.com/shopspring/decimal"
)

type AccountRepositoryInterface interface {
	Create(ctx context.Context, id int64, balance decimal.Decimal) error 
	GetByID(ctx context.Context, id int64) (int64, decimal.Decimal, error)
}
