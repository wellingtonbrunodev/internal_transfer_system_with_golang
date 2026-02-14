package service

import (
	"context"
	"github.com/shopspring/decimal"
)

type TransactionRepositoryInterface interface {
	Transfer(
		ctx context.Context,
		sourceID int64,
		destinationID int64,
		amount decimal.Decimal,
	) error
}
